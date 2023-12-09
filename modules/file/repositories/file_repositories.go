package repositories

import (
	"github.com/bukharney/FileServer/modules/entities"
	"github.com/jmoiron/sqlx"
)

type FileRepo struct {
	Db *sqlx.DB
}

func NewFileRepo(db *sqlx.DB) entities.FileRepository {
	return &FileRepo{Db: db}
}

func (f *FileRepo) Upload(req *entities.FileUploadReq) (entities.FileUploadRes, error) {
	res := entities.FileUploadRes{}

	return res, nil
}
