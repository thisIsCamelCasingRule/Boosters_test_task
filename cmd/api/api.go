package api

import (
	"Boosters_test_task/internal/service"
	"github.com/jackc/pgx/v5"
)

type Api struct {
	srv service.Serve
}

func NewApi(db *pgx.Conn) Api {
	return Api{
		srv: service.NewService(db),
	}
}

type Handle interface {
	HandlePosts
}

