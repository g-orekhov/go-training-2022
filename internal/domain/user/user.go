package user

type User struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type EventUser struct {
	UserId  int64
	EventId int64
}
