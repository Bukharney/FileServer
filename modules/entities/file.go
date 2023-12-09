package entities

import (
	"context"
	"io"
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
	Upload(ctx context.Context, req *FileUploadReq) (FileUploadRes, error)
	StreamFileUpload(w io.Writer, bucket, object string, file multipart.File) error
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
