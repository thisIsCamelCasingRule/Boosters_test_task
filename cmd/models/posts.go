package models

import (
	"Boosters_test_task/internal/models"
	"errors"
	"time"
)

type CreatePostRequest struct {
	Title     string    `json:"title,omitempty" form:"title,omitempty"`
	Content   string    `json:"content,omitempty" form:"content,omitempty"`
}

func (cr CreatePostRequest) ToPost() models.Post{
	return models.Post{
		Title:     cr.Title,
		Content:   cr.Content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (cr CreatePostRequest) Validate() error {
	if len(cr.Title) == 0 {
		return errors.New("empty title")
	}

	if len(cr.Content) == 0 {
		return errors.New("empty content")
	}

	return nil
}

type PutPostRequest struct {
	Title     string    `json:"title,omitempty" form:"title,omitempty"`
	Content   string    `json:"content,omitempty" form:"content,omitempty"`
}

func (pr PutPostRequest) ToPost() models.Post{
	return models.Post{
		Title:     pr.Title,
		Content:   pr.Content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (pr PutPostRequest) Validate() error {
	if len(pr.Title) == 0 {
		return errors.New("empty title")
	}

	if len(pr.Content) == 0 {
		return errors.New("empty content")
	}

	return nil
}
