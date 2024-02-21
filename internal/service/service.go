package service

import (
	"Boosters_test_task/pkg/database/queries"
	"github.com/jackc/pgx/v5"
)

type Service struct {
	q queries.Querier
}

type Serve interface {
	ServPosts
}

func NewService(db *pgx.Conn) Service {
	return Service{
		q: queries.New(db),
	}
}


