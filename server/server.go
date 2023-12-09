package server

import (
	"errors"

	"cloud.google.com/go/storage"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type Server struct {
	App    *gin.Engine
	DB     *sqlx.DB
	Bucket *storage.BucketHandle
}

func NewServer(db *sqlx.DB, bucket *storage.BucketHandle) *Server {
	return &Server{
		App:    gin.Default(),
		DB:     db,
		Bucket: bucket,
	}
}

func (s *Server) Run() error {

	s.App.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	err := s.MapHandlers()
	if err != nil {
		return errors.New("failed to map handlers")
	}

	err = s.FileServer()
	if err != nil {
		return errors.New("failed to run file server")
	}

	err = s.App.Run(":8081")
	if err != nil {
		return errors.New("failed to run gin")
	}

	return nil
}
