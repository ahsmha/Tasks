package model

import (
	"time"
)

type User struct {
	Id      int       `json:"id"`
	Name    string    `json:"name"`
	Role    string    `json:"role"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}
