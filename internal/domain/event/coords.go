package event

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"log"
	"math"
	"strconv"
)

type Coords struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}

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
func (c *Coords) Scan(src interface{}) (err error) {
	str, ok := src.(string)
	// if #src is not a string interface it probably "nil", it can be if no value in the data base
	if !ok {
		c.Lat, c.Long = 0, 0
		return nil
	}
	// if the length of the string not equal 50, it means that there is not a geo-point in the input
	if len(str) != 50 {
		return errors.New("geo-point must have length 25 bites, probably another data was given")
	}
	// in case of panic #defer will do loging and normal recover, so we are ignoring all the next errors
	defer coordsParsingRecover(&err)
	// parse coordinates
	firstByteValue, _ := strconv.ParseUint(str[:2], 16, 64)
	isLittleEndian := firstByteValue == 1
	xUint, _ := parseUint64FromHexString(str, 18, isLittleEndian)
	yUint, _ := parseUint64FromHexString(str, 34, isLittleEndian)
	c.Long = math.Float64frombits(xUint)
	c.Lat = math.Float64frombits(yUint)
	return nil
}

// Makes #Scan method to return common error in case of panic and logs the panic message.
func coordsParsingRecover(err *error) {
	if errRecover := recover(); errRecover != nil {
		log.Println("Panic in Coors.Scan():", errRecover)
		(*err) = errors.New("failed to parse #Coords in #Coords.Scan()")
	}
}

func parseUint64FromHexString(stringToParse string, startFrom int, isLittleEndian bool) (uint64, error) {
	if isLittleEndian {
		stringToParse = reverseHexString(stringToParse, startFrom, 8)
	}
	return strconv.ParseUint(stringToParse, 16, 64)
}

func reverseHexString(str string, startFrom, length int) (stringToReturn string) {
	for i := startFrom + (length * 2) - 2; i >= startFrom; i -= 2 {
		stringToReturn += str[i : i+2]
	}
	return
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
