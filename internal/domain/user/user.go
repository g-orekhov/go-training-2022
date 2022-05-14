package user

type User struct {
	Id          int64
	Name        string
	EventUserId int64
}

type EventUser struct {
	UserId  int64
	EventId int64
}
