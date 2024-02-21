package api

import (
	models2 "Boosters_test_task/cmd/models"
	_"Boosters_test_task/internal/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type HandlePosts interface {
	GetPosts(c *gin.Context)
	CreatePost(c *gin.Context)
	GetPostById(c *gin.Context)
	PutPostById(c *gin.Context)
	DeletePostById(c *gin.Context)
}

// GetPosts godoc
// @Summary get all posts
// @Tags posts
// @Produce json
// @Success 200 {object} float64
// @Router /posts [get]
func (a Api) GetPosts(c *gin.Context) {
	result, err := a.srv.GetPosts()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, "Can not get posts")
		return
	}

	c.JSON(http.StatusOK, result)
}

// CreatePost godoc
// @Summary create post
// @Tags posts
// @Produce json
// @Param request body models.CreatePostRequest true "query params"
// @Success 200 {string} string
// @Router /posts [post]
func (a Api) CreatePost(c *gin.Context) {
	var post models2.CreatePostRequest

	err := c.Bind(&post)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "Invalid request")
		return
	}

	if err = post.Validate(); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, fmt.Sprintf("Invalid request: %s", err))
		return
	}

	postModel := post.ToPost()

	err = a.srv.CreatePost(postModel)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, "Can not create post")
		return
	}

	c.JSON(http.StatusOK, "OK")
}

// GetPostById godoc
// @Summary get post by id
// @Tags posts
// @Produce json
// @Param id path int true "post id"
// @Success 200 {object} models.Post
// @Router /posts/{id} [get]
func (a Api) GetPostById(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "Invalid query param id")
		return
	}

	result, err := a.srv.GetPostById(id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, "Can not get post by id")
		return
	}

	c.JSON(http.StatusOK, result)
}

// PutPostById godoc
// @Summary put post by id
// @Tags posts
// @Produce json
// @Param request body models.PutPostRequest true "query params"
// @Param id path int true "post id"
// @Success 200 {object} models.Post
// @Router /posts/{id} [put]
func (a Api) PutPostById(c *gin.Context) {
	var post models2.PutPostRequest

	idString := c.Param("id")
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "Invalid query param id")
		return
	}

	err = c.ShouldBind(&post)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "Invalid request")
		return
	}

	if err = post.Validate(); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, fmt.Sprintf("Invalid request: %s", err))
		return
	}

	postModel := post.ToPost()
	postModel.Id = id

	result, err := a.srv.PutPost(postModel)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, "Can not put post")
		return
	}

	c.JSON(http.StatusOK, result)
}

// DeletePostById godoc
// @Summary delete post by id
// @Tags posts
// @Produce json
// @Param id path int true "post id"
// @Success 200 {object} models.Post
// @Router /posts/{id} [delete]
func (a Api) DeletePostById(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "Invalid query param id")
		return
	}

	err = a.srv.DeletePostById(id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, "Can not delete post by id")
		return
	}

	c.JSON(http.StatusOK, "deleted")
}
