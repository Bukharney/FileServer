package server

import (
	"net/http"

	_fileController "github.com/bukharney/FileServer/modules/file/controllers"
	_fileRepo "github.com/bukharney/FileServer/modules/file/repositories"
	_fileUsecase "github.com/bukharney/FileServer/modules/file/usecases"

	"github.com/gin-gonic/gin"
)

func (s *Server) MapHandlers() error {

	v1 := s.App.Group("/v1")
	fileGroup := v1.Group("/file")
	fileRepo := _fileRepo.NewFileRepo(s.DB)
	fileUsecase := _fileUsecase.NewFileUsecase(fileRepo)
	_fileController.NewFileControllers(fileGroup, fileUsecase)

	s.App.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not Found"})
	})

	return nil
}

func (s *Server) FileServer() error {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	s.App.GET("/static/*filepath", func(c *gin.Context) {
		mux.ServeHTTP(c.Writer, c.Request)
	})

	s.App.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not Found"})
	})

	return nil
}
