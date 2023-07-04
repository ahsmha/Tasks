package infra

import (
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

// The database can be referenced through the global variable db in the repository package
func SetDB(d *sqlx.DB) {
	db = d
}
