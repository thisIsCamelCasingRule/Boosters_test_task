package service

import (
	mockqueries "Boosters_test_task/internal/mock/queries"
	"Boosters_test_task/internal/models"
	"Boosters_test_task/pkg/database/queries"
	"context"
	"errors"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestService_GetPosts_Success(t *testing.T) {
	q := mockqueries.NewQuerier(t)
	ctx := context.Background()
	timeParam := time.Now()
	getPostsResponce := []queries.Post{
		{
			ID:        1,
			Title:     pgtype.Text{String: "A", Valid: true},
			Content:   pgtype.Text{String: "AAAa", Valid: true},
			CreatedAt: pgtype.Timestamptz{Time: timeParam, Valid: true},
			UpdatedAt: pgtype.Timestamptz{Time: timeParam, Valid: true},
		},
		{
			ID:        2,
			Title:     pgtype.Text{String: "B", Valid: true},
			Content:   pgtype.Text{String: "BBBb", Valid: true},
			CreatedAt: pgtype.Timestamptz{Time: timeParam, Valid: true},
			UpdatedAt: pgtype.Timestamptz{Time: timeParam, Valid: true},
		},
	}

	expect := []models.Post{
		{
			Id:        1,
			Title:     "A",
			Content:  "AAAa",
			CreatedAt: timeParam,
			UpdatedAt: timeParam,
		},
		{
			Id:        2,
			Title:     "B",
			Content:   "BBBb",
			CreatedAt: timeParam,
			UpdatedAt: timeParam,
		},
	}

	q.On("GetPosts", ctx).Return(getPostsResponce, nil).Once()

	srv := Service{q: q}

	result, err := srv.GetPosts()
	assert.Nil(t, err)
	assert.Equal(t, expect, result)
}

func TestService_GetPosts_Fail_DB_Error(t *testing.T) {
	q := mockqueries.NewQuerier(t)
	ctx := context.Background()

	errorString := "db error"
	queryError := errors.New(errorString)
	var expectedResult []models.Post

	q.On("GetPosts", ctx).Return(nil, queryError).Once()

	srv := Service{q: q}

	result, err := srv.GetPosts()
	assert.EqualError(t, err, errorString)
	assert.Equal(t, expectedResult, result)
}

func TestService_GetPostById_Success(t *testing.T) {
	q := mockqueries.NewQuerier(t)
	ctx := context.Background()

	id := int64(1)
	timeParam := time.Now()
	getPostByIdResponce := queries.Post{
		ID:        id,
		Title:     pgtype.Text{String: "A", Valid: true},
		Content:   pgtype.Text{String: "AAAa", Valid: true},
		CreatedAt: pgtype.Timestamptz{Time: timeParam, Valid: true},
		UpdatedAt: pgtype.Timestamptz{Time: timeParam, Valid: true},
	}

	expect := models.Post{
		Id:        id,
		Title:     "A",
		Content:  "AAAa",
		CreatedAt: timeParam,
		UpdatedAt: timeParam,
	}

	q.On("GetPostById", ctx, id).Return(getPostByIdResponce, nil).Once()

	srv := Service{q: q}

	result, err := srv.GetPostById(id)
	assert.Nil(t, err)
	assert.Equal(t, expect, result)
}

func TestService_GetPostById_Fail_DB_Error(t *testing.T) {
	q := mockqueries.NewQuerier(t)
	ctx := context.Background()

	id := int64(1)

	var expectedResult models.Post
	errorString := "db error"
	queryError := errors.New(errorString)

	var queryResult queries.Post
	q.On("GetPostById", ctx, id).Return(queryResult, queryError).Once()

	srv := Service{q: q}

	result, err := srv.GetPostById(id)
	assert.EqualError(t, err, errorString)
	assert.Equal(t,expectedResult, result)
}

func TestService_CreatePost_Success(t *testing.T) {
	q := mockqueries.NewQuerier(t)
	ctx := context.Background()

	timeParam := time.Now()
	methodPostParam := models.Post{
		Id:        1,
		Title:     "A",
		Content:  "AAAa",
		CreatedAt: timeParam,
		UpdatedAt: timeParam,
	}

	queryParam := queries.CreatePostParams{
		Title:     pgtype.Text{String: "A", Valid: true},
		Content:   pgtype.Text{String: "AAAa", Valid: true},
		CreatedAt: pgtype.Timestamptz{Time: timeParam, Valid: true},
		UpdatedAt: pgtype.Timestamptz{Time: timeParam, Valid: true},
	}

	q.On("CreatePost", ctx, queryParam).Return(nil).Once()

	srv := Service{q: q}

	err := srv.CreatePost(methodPostParam)
	assert.Nil(t, err)
}

func TestService_CreatePost_Fail_DB_Error(t *testing.T) {
	q := mockqueries.NewQuerier(t)
	ctx := context.Background()

	timeParam := time.Now()
	methodPostParam := models.Post{
		Id:        1,
		Title:     "A",
		Content:  "AAAa",
		CreatedAt: timeParam,
		UpdatedAt: timeParam,
	}

	queryParam := queries.CreatePostParams{
		Title:     pgtype.Text{String: "A", Valid: true},
		Content:   pgtype.Text{String: "AAAa", Valid: true},
		CreatedAt: pgtype.Timestamptz{Time: timeParam, Valid: true},
		UpdatedAt: pgtype.Timestamptz{Time: timeParam, Valid: true},
	}

	errorString := "db error"
	queryError := errors.New(errorString)

	q.On("CreatePost", ctx, queryParam).Return(queryError).Once()

	srv := Service{q: q}

	err := srv.CreatePost(methodPostParam)
	assert.EqualError(t, err, errorString)
}

func TestService_DeletePostById_Success(t *testing.T) {
	q := mockqueries.NewQuerier(t)
	ctx := context.Background()

	id := int64(1)

	q.On("DeletePostById", ctx, id).Return(nil).Once()

	srv := Service{q: q}

	err := srv.DeletePostById(id)
	assert.Nil(t, err)
}

func TestService_DeletePostById_Fail_DB_Error(t *testing.T) {
	q := mockqueries.NewQuerier(t)
	ctx := context.Background()

	id := int64(1)
	errorString := "db error"
	queryError := errors.New(errorString)

	q.On("DeletePostById", ctx, id).Return(queryError).Once()

	srv := Service{q: q}

	err := srv.DeletePostById(id)
	assert.EqualError(t, err, errorString)
}

func TestService_PutPost_Not_Exists_Success(t *testing.T) {
	q := mockqueries.NewQuerier(t)
	ctx := context.Background()

	timeParam := time.Now()
	methodPostParam := models.Post{
		Id:        1,
		Title:     "A",
		Content:  "AAAa",
		CreatedAt: timeParam,
		UpdatedAt: timeParam,
	}

	queryGetParam := queries.Post{}

	queryCreateParam := queries.CreatePostParams{
		Title:     pgtype.Text{String: "A", Valid: true},
		Content:   pgtype.Text{String: "AAAa", Valid: true},
		CreatedAt: pgtype.Timestamptz{Time: timeParam, Valid: true},
		UpdatedAt: pgtype.Timestamptz{Time: timeParam, Valid: true},
	}

	q.On("GetPostById", ctx, methodPostParam.Id).Return(queryGetParam, nil).Once()
	q.On("CreatePost", ctx, queryCreateParam).Return(nil).Once()

	srv := Service{q: q}

	result, err := srv.PutPost(methodPostParam)
	assert.Nil(t, err)
	assert.Equal(t, methodPostParam, result)
}

func TestService_PutPost_Exists_Success(t *testing.T) {
	q := mockqueries.NewQuerier(t)
	ctx := context.Background()

	timeParam := time.Now()
	updTime := timeParam.Add(time.Hour * 20)
	methodPostParam := models.Post{
		Id:        1,
		Title:     "B",
		Content:  "BBBb",
		CreatedAt: timeParam,
		UpdatedAt: timeParam,
	}

	queryGetResponse := queries.Post{
		ID: 1,
		Title:     pgtype.Text{String: "A", Valid: true},
		Content:   pgtype.Text{String: "AAAa", Valid: true},
		CreatedAt: pgtype.Timestamptz{Time: timeParam, Valid: true},
		UpdatedAt: pgtype.Timestamptz{Time: timeParam, Valid: true},
	}

	queryUpdateResponse := queries.Post{
		ID: 1,
		Title:     pgtype.Text{String: "B", Valid: true},
		Content:   pgtype.Text{String: "BBBb", Valid: true},
		CreatedAt: pgtype.Timestamptz{Time: timeParam, Valid: true},
		UpdatedAt: pgtype.Timestamptz{Time: updTime, Valid: true},
	}

	expectResult := models.Post{
		Id:        1,
		Title:     "B",
		Content:   "BBBb",
		CreatedAt: timeParam,
		UpdatedAt: updTime,
	}

	q.On("GetPostById", ctx, methodPostParam.Id).Return(queryGetResponse, nil).Once()
	q.On("UpdatePostById", ctx, mock.AnythingOfType("queries.UpdatePostByIdParams")).Return(queryUpdateResponse, nil).Once()

	srv := Service{q: q}

	result, err := srv.PutPost(methodPostParam)
	assert.Nil(t, err)
	assert.Equal(t, expectResult, result)
}

func TestService_PutPost_Exists_Fail_DB_Error_Update(t *testing.T) {
	q := mockqueries.NewQuerier(t)
	ctx := context.Background()

	timeParam := time.Now()
	methodPostParam := models.Post{
		Id:        1,
		Title:     "B",
		Content:  "BBBb",
		CreatedAt: timeParam,
		UpdatedAt: timeParam,
	}

	queryGetResponse := queries.Post{
		ID: 1,
		Title:     pgtype.Text{String: "A", Valid: true},
		Content:   pgtype.Text{String: "AAAa", Valid: true},
		CreatedAt: pgtype.Timestamptz{Time: timeParam, Valid: true},
		UpdatedAt: pgtype.Timestamptz{Time: timeParam, Valid: true},
	}

	var expectedResult models.Post
	errorString := "db update error"

	queryUpdateError := errors.New(errorString)
	var queryUpdatePostRow queries.Post

	q.On("GetPostById", ctx, methodPostParam.Id).Return(queryGetResponse, nil).Once()
	q.On("UpdatePostById", ctx, mock.AnythingOfType("queries.UpdatePostByIdParams")).Return(queryUpdatePostRow, queryUpdateError).Once()

	srv := Service{q: q}

	result, err := srv.PutPost(methodPostParam)

	assert.EqualError(t, err,  errorString)
	assert.Equal(t, expectedResult, result)
}

func TestService_PutPost_Not_Exists_Fail_DB_Create_Error(t *testing.T) {
	q := mockqueries.NewQuerier(t)
	ctx := context.Background()

	timeParam := time.Now()
	methodPostParam := models.Post{
		Id:        1,
		Title:     "A",
		Content:  "AAAa",
		CreatedAt: timeParam,
		UpdatedAt: timeParam,
	}

	queryGetParam := queries.Post{}

	queryCreateParam := queries.CreatePostParams{
		Title:     pgtype.Text{String: "A", Valid: true},
		Content:   pgtype.Text{String: "AAAa", Valid: true},
		CreatedAt: pgtype.Timestamptz{Time: timeParam, Valid: true},
		UpdatedAt: pgtype.Timestamptz{Time: timeParam, Valid: true},
	}

	var expectedResult models.Post
	errorString := "db create error"
	queryCreateError := errors.New(errorString)

	q.On("GetPostById", ctx, methodPostParam.Id).Return(queryGetParam, nil).Once()
	q.On("CreatePost", ctx, queryCreateParam).Return(queryCreateError).Once()

	srv := Service{q: q}

	result, err := srv.PutPost(methodPostParam)
	assert.EqualError(t, err, errorString)
	assert.Equal(t, expectedResult, result)
}

func TestService_PutPost_Fail_DB_Get_Error(t *testing.T) {
	q := mockqueries.NewQuerier(t)
	ctx := context.Background()

	timeParam := time.Now()
	methodPostParam := models.Post{
		Id:        1,
		Title:     "A",
		Content:  "AAAa",
		CreatedAt: timeParam,
		UpdatedAt: timeParam,
	}

	var expectedResult models.Post
	errorString := "db get error"

	var queryGetResponse queries.Post
	queryGetError := errors.New(errorString)

	q.On("GetPostById", ctx, methodPostParam.Id).Return(queryGetResponse, queryGetError).Once()

	srv := Service{q: q}

	result, err := srv.PutPost(methodPostParam)
	assert.EqualError(t, err, errorString)
	assert.Equal(t, expectedResult, result)
}

