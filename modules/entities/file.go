package entities

import (
	"mime/multipart"

	"github.com/golang-jwt/jwt/v4"
)

type UsersClaims struct {
	Id       int    `json:"user_id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

type FileUsecase interface {
	Upload(req *FileUploadReq) (FileUploadRes, error)
}

type FileRepository interface {
	Upload(req *FileUploadReq) (FileUploadRes, error)
}

type FileUploadReq struct {
	File []*multipart.FileHeader `form:"file"`
}

type FileUploadRes struct {
	FilePaths []string `json:"file_paths" db:"file_paths"`
}
