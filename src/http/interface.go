package http

import "database/sql"

type Http struct {
	DB *sql.DB
}

type iHttp interface {
	Launch()
}

func New(http *Http) iHttp {
	return http
}
