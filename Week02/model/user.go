package model

import "github.com/pkg/errors"

// Model errors that are exposed to API layers
var (
	ErrNoRecord = errors.New("model: no record found")
	ErrNoCreate = errors.New("model: failed to create entity")
)

// User represents an user ...
type User struct {
	ID        int
	FirstName string
	LastName  string
}

// UserDaoService is the interface any storage service needs to provide for UserService
type UserDaoService interface {
	Create(*User) (int, error)
	Get(int) (*User, error)
}

// UserService manages user entities
type UserService struct {
	daoClient UserDaoService
}

// NewUserService returns an instance that handles User query and creation
func NewUserService(daoClient UserDaoService) *UserService {
	return &UserService{
		daoClient: daoClient,
	}
}

func (us *UserService) Create(user *User) (int, error) {
	// More logic
	return us.daoClient.Create(user)
}

func (us *UserService) Get(id int) (*User, error) {
	// More logic
	return us.daoClient.Get(id)
}
