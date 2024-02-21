package service

import (
	"Boosters_test_task/internal/models"
	"Boosters_test_task/pkg/database/queries"
	"context"
	"github.com/jackc/pgx/v5/pgtype"
	"time"
)

type ServPosts interface {
	GetPosts() ([]models.Post, error)
	GetPostById(id int64) (models.Post, error)
	CreatePost(post models.Post) error
	PutPost(post models.Post) (models.Post, error)
	DeletePostById(id int64) error
}

func (s Service) GetPosts() ([]models.Post, error) {
	posts, err := s.q.GetPosts(context.Background())
	if err != nil {
		return nil, err
	}

	result := make([]models.Post, len(posts))
	for i, post := range posts {
		tmpPost := models.Post{}
		tmpPost.FromQueryResult(post)
		result[i] = tmpPost
	}

	return result, nil
}

func (s Service) GetPostById(id int64) (models.Post, error) {
	var result models.Post
	postRow, err := s.q.GetPostById(context.Background(), id)
	if err != nil {
		return result, err
	}

	result.FromQueryResult(postRow)

	return result, nil
}

func (s Service) CreatePost(post models.Post) error {
	params := post.ToQueryParams()

	err := s.q.CreatePost(context.Background(), params)
	if err != nil {
		return err
	}

	return nil
}

func (s Service) PutPost(post models.Post) (models.Post, error) {
	postRow, err := s.q.GetPostById(context.Background(), post.Id)
	if err != nil {
		return models.Post{}, err
	}

	if postRow.ID != 0 {
		var updatedPost models.Post

		updateParams := queries.UpdatePostByIdParams{
			Title:   pgtype.Text{post.Title, true},
			Content: pgtype.Text{post.Content, true},
			UpdatedAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},
			ID:      post.Id,
		}

		updatedPostRow, err := s.q.UpdatePostById(context.Background(), updateParams)
		if err != nil {
			return updatedPost, err
		}

		updatedPost.FromQueryResult(updatedPostRow)

		return updatedPost, nil
	}

	createPostParams := post.ToQueryParams()
	err = s.q.CreatePost(context.Background(), createPostParams)
	if err != nil {
		return models.Post{}, err
	}

	return post, nil
}

func (s Service) DeletePostById(id int64) error {
	err := s.q.DeletePostById(context.Background(), id)
	if err != nil {
		return err
	}

	return nil
}