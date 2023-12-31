package server

import (
	"errors"

	"cloud.google.com/go/storage"
	"github.com/bukharney/FileServer/configs"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type Server struct {
	App     *gin.Engine
	DB      *sqlx.DB
	Cfg     *configs.Configs
	Storage *storage.Client
}

func NewServer(db *sqlx.DB, cfg *configs.Configs, storage *storage.Client) *Server {
	return &Server{
		App:     gin.Default(),
		DB:      db,
		Cfg:     cfg,
		Storage: storage,
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
