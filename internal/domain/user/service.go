package user

type Service interface {
	FindAll() ([]User, error)
	FindOne(int64) (*User, error)
	Create(*User) error
	Delete(int64) error
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{
		repo: r,
	}
}

func (s *service) FindAll() ([]User, error) {
	return s.repo.FindAll()
}

func (s *service) FindOne(id int64) (*User, error) {
	return s.repo.FindOne(id)
}

func (s *service) Create(user *User) error {
	return s.repo.Create(user)
}
func (s *service) Delete(id int64) error {
	return s.repo.Delete(id)
}
