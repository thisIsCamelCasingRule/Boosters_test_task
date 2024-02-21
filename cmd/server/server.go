package server

import (
	"Boosters_test_task/cmd/api"
	"Boosters_test_task/config"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	"runtime"
	"time"
)

type Server struct {
	srv *http.Server
	api api.Handle
}

func NewServer() Server {
	return Server{}
}

func (s Server) Start(ctx context.Context) error {
	// Get configs
	cfg, err := config.ReadConfig()
	if err != nil {
		log.Fatalf("can not parse config: %s", err)
	}

	// Connect to postgres
	postgresURL := fmt.Sprintf("user=%s host=%s port=%s dbname=%s password=%s sslmode=disable",
		cfg.DB.User,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.DBName,
		cfg.DB.Password)

	postgresCtx, postgresCancel := context.WithTimeout(ctx, time.Duration(time.Second.Nanoseconds()*5))
	defer postgresCancel()

	conn, err := pgx.Connect(postgresCtx, postgresURL)
	if err != nil {
		fmt.Println("postgres error")
		return err
	}
	defer conn.Close(context.Background())

	// Init server struct
	s.api = api.NewApi(conn)

	s.srv = &http.Server{
		Addr:    cfg.Server.Address,
		Handler: s.registerRoutes(),
	}

	// Start http server
	go func() {
		if err := s.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("failed to start server: %s", err)
		}
	}()

	select {
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*15)
		defer cancel()
		return s.srv.Shutdown(timeout)
	}
}

func (s Server) registerRoutes() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		runtime.Gosched()
		c.JSON(http.StatusOK, "pong")
	})

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("http://127.0.0.1:8080/swagger/doc.json")))

	postsRoutes := router.Group("/posts")
	{
		postsRoutes.GET("", s.api.GetPosts)
		postsRoutes.GET("/:id", s.api.GetPostById)

		postsRoutes.PUT("/:id", s.api.PutPostById)

		postsRoutes.POST("", s.api.CreatePost)

		postsRoutes.DELETE("/:id", s.api.DeletePostById)
	}

	return router
}
