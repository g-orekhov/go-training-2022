package user

import "fmt"

type Repository interface {
	FindAll() ([]User, error)
	FindOne(id int64) (*User, error)
}

const UsersCount int64 = 10

type repository struct {
	// Some internal data
}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) FindAll() ([]User, error) {
	users := make([]User, UsersCount)
	for i := int64(0); i < UsersCount; i++ {
		users[i] = User{
			Id:   i + 1,
			Name: fmt.Sprintf("User #%d", i+1),
		}
	}
	return users, nil
}

func (r *repository) FindOne(id int64) (*User, error) {
	if id <= UsersCount {
		return &User{
			Id:   id,
			Name: fmt.Sprintf("User #%d", id),
		}, nil
	} else {
		return nil, nil
	}
}
