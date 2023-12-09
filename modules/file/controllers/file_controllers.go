package controllers

import (
	"net/http"

	"github.com/bukharney/FileServer/middlewares"
	"github.com/bukharney/FileServer/modules/entities"
	"github.com/gin-gonic/gin"
)

type FileController struct {
	FileUsecase entities.FileUsecase
}

func NewFileControllers(r gin.IRoutes, fileUsecase entities.FileUsecase) {
	controllers := &FileController{

		FileUsecase: fileUsecase,
	}

	r.POST("/upload", controllers.Upload, middlewares.JwtAuthentication())
}

const MAX_UPLOAD_SIZE = 1024 * 1024 * 100

func (f *FileController) Upload(c *gin.Context) {
	role, err := middlewares.GetUserByToken(c)
	if err != nil {
		return
	}

	if role.Role != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You are not admin"})
		return
	}

	var req entities.FileUploadReq

	if err := c.Request.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	files := c.Request.MultipartForm.File["file"]
	req.File = files

	res, err := f.FileUsecase.Upload(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, res)
}
