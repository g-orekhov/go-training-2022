package event

import (
	"database/sql/driver"
	"fmt"
	"math"
	"strconv"
)

type Coords struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}

func (c *Coords) Scan(src interface{}) error {
	/*
		В БД формат point(Lon,Lat)::geometry
		по умолчанию БД возврашает точку в формате - EWKB или что-то вроде (гугл не помогает)
		Структура кторую удалось распарсить:
		первый байт - порядок байт Big/Little-endian (у нас он 1 - little-endian)
		4 байта - тип обьекта (но у нас почему то 3 байта тип, а последний байт Х/З что)
		4 байта - SRID
		8 + 8 байт - значения точки
		Значения читаю как UINT и затем конвертирую в FLOAT64. ParseFloat выдаёт бред какой-то
	*/
	// TODO: Интерфейс src - string, но нужно поискать более правильный способ работы с байтами
	str, ok := src.(string)
	if !ok {
		// If nill in the database
		return nil
	}
	if len(str) != 50 {
		//TODO: может ли точка иметь другой размер? Узнать - обработать.
		c.Lat = 0
		c.Long = 0
	}
	reverse := func(str string, start int, bytes int) string {
		var ret string
		for i := start + (bytes * 2) - 2; i >= start; i -= 2 {
			ret += str[i : i+2]
		}
		return ret
	}
	//TODO: Разобраться возможны ли тут ошибки? Если да - обработать.
	xUint, _ := strconv.ParseUint(reverse(str, 18, 8), 16, 64)
	yUint, _ := strconv.ParseUint(reverse(str, 34, 8), 16, 64)
	c.Long = math.Float64frombits(xUint)
	c.Lat = math.Float64frombits(yUint)
	return nil
}

func (c Coords) Value() (driver.Value, error) {
	/*
		Returns point string to store geopoint in the Postgres
		"SRID=4326;POINT(LONG, LAT)"
	*/
	return fmt.Sprintf("SRID=4326;POINT(%f %f)", c.Long, c.Lat), nil
}

func (c *Coords) String() string {
	/*
		Returns point string to store geopoint in the Postgres
		"SRID=4326;POINT(LONG, LAT)"
	*/
	return fmt.Sprintf("POINT(%f %f)", c.Long, c.Lat)
}
