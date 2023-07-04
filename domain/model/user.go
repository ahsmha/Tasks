package model

import (
	"time"
)

type User struct {
	Id       uint      `json:"id" db:"id"`
	Name     string    `json:"name" db:"name"`
	Password []byte    `json:"-"`
	Created  time.Time `json:"created" db:"created"`
	Updated  time.Time `json:"updated" db:"updated"`
}
