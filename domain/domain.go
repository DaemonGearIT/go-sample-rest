package domain

import (
	"time"
)

type (
	Repository interface {
		Exists(id uint64) bool
		Create(interface{}) interface{}
		Update(interface{}) interface{}
		Delete(interface{}) bool
		FindById(id uint64) interface{}
		FindAll() []interface{}
	}

	UserRepository interface {
		Repository
	}

	User struct {
		Id        uint64
		Email     string
		Password  string
		CreatedAt time.Time
		UpdatedAt time.Time
	}
)
