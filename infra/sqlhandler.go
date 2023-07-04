package infra

import "github.com/jmoiron/sqlx"

type SqlHandler struct {
	Conn *sqlx.DB
}

func NewSqlHandler() *SqlHandler {
	// Data Source Name
	dsn, err := dsn()
	if err != nil {
		panic(err.Error)
	}

	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		panic(err.Error)
	}
	if err := db.Ping(); err != nil {
		panic(err.Error)
	}

	return &SqlHandler{Conn: db}
}

func ConnectDB() (*sqlx.DB, error) {
	handler := NewSqlHandler()

	return handler.Conn, nil
}

func dsn() (string, error) {
	return "", nil
}
