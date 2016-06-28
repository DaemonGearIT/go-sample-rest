package usecases

import (
	"errors"
	"github.con/DaemonGearIT/go-sample-rest/domain"
	"time"
)

type (
	User struct {
		Id       uint64
		Email    string
		Password string
	}
	UserInteractor struct {
		UserRepository domain.UserRepository
	}
)

//Save an existing User or create a new
func (interactor *UserInteractor) SaveUser(user User) error {
	if &user == nil {
		return errors.New("User must not be nil")
	}

	dUser := mapToDomain(user)
	if interactor.UserRepository.Exists(user.Id) {
		interactor.UserRepository.Save(dUser)
	} else {
		interactor.UserRepository.Create(dUser)
	}
}

func mapFromDomain(dUser domain.User) (user User) {

	user.Id = dUser.Id
	user.Email = dUser.Email
	user.Password = "********"

	return
}

//Map usecase.User to domain.User
func mapToDomain(user User) (dUser domain.User) {
	if user.Id > 0 {
		dUser.Id = user.Id
	}

	dUser.Email = user.Email
	dUser.Password = user.Password
	dUser.CreatedAt = time.Now().UTC()
	return
}
