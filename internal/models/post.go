package models

import (
	"Boosters_test_task/pkg/database/queries"
	"github.com/jackc/pgx/v5/pgtype"
	"time"
)

type Post struct {
	Id        int64     `json:"id,omitempty" `
	Title     string    `json:"title,omitempty" `
	Content   string    `json:"content,omitempty" `
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (p *Post) FromQueryResult(post queries.Post) {
	p.Id = post.ID
	p.Title = post.Title.String
	p.Content = post.Content.String
	p.UpdatedAt = post.UpdatedAt.Time
	p.CreatedAt = post.CreatedAt.Time
}

func (p *Post) ToQueryParams() queries.CreatePostParams {
	return queries.CreatePostParams{
		Title:     pgtype.Text{p.Title, true},
		Content:   pgtype.Text{p.Content, true},
		CreatedAt: pgtype.Timestamptz{Time: p.CreatedAt, Valid: true},
		UpdatedAt: pgtype.Timestamptz{Time: p.UpdatedAt, Valid: true},
	}
}