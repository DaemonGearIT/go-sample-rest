package postgres

import (
	"github.com/daemongearit/go-sample-rest/domain"
	"infraestructure"
)

type UserRepo struct {
}

/**
Exists(interface) bool
		Create(interface) interface
		Update(interface) interface
		Delete(interface) bool
		FindById(id uint64) interface
		FindAll() []interface
*/

func (repo *UserRepo) Exists(id uint64) bool {
	db := infraestructure.PostgresSession()
	var user domain.User
	db.(&user, id)
	return &user != nil
}
