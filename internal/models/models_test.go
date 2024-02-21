package models

import (
	"Boosters_test_task/pkg/database/queries"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestPost_FromQueryResult(t *testing.T) {
	timeParam := time.Now()
	expect := Post{
		Id:        1,
		Title:     "A",
		Content:  "AAAa",
		CreatedAt: timeParam,
		UpdatedAt: timeParam,
	}

	params := queries.Post{
		ID: 1,
		Title:     pgtype.Text{String: "A", Valid: true},
		Content:   pgtype.Text{String: "AAAa", Valid: true},
		CreatedAt: pgtype.Timestamptz{Time: timeParam, Valid: true},
		UpdatedAt: pgtype.Timestamptz{Time: timeParam, Valid: true},
	}

	var result Post

	result.FromQueryResult(params)
	assert.Equal(t, result, expect)
}

func TestPost_ToQueryParams(t *testing.T) {
	timeParam := time.Now()
	params := Post{
		Id:        1,
		Title:     "A",
		Content:  "AAAa",
		CreatedAt: timeParam,
		UpdatedAt: timeParam,
	}

	expect := queries.CreatePostParams{
		Title:     pgtype.Text{String: "A", Valid: true},
		Content:   pgtype.Text{String: "AAAa", Valid: true},
		CreatedAt: pgtype.Timestamptz{Time: timeParam, Valid: true},
		UpdatedAt: pgtype.Timestamptz{Time: timeParam, Valid: true},
	}

	result := params.ToQueryParams()
	assert.Equal(t, result, expect)
}
