package usecases

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/bukharney/FileServer/modules/entities"
	"github.com/bukharney/FileServer/utils"
)

const MAX_UPLOAD_SIZE = 1024 * 1024 * 20

type FileUsecase struct {
	FileRepo entities.FileRepository
}

func NewFileUsecase(fileRepo entities.FileRepository) entities.FileUsecase {
	return &FileUsecase{FileRepo: fileRepo}
}

func (f *FileUsecase) Upload(req *entities.FileUploadReq) (entities.FileUploadRes, error) {
	res := entities.FileUploadRes{}

	files := req.File
	FileName := []string{}
	for _, fileHeader := range files {

		if fileHeader.Size > MAX_UPLOAD_SIZE {
			return res, fmt.Errorf("file too large")
		}

		file, err := fileHeader.Open()

		if err != nil {
			return res, fmt.Errorf("error, failed to open file")
		}

		defer file.Close()

		buff := make([]byte, 512)
		_, err = file.Read(buff)
		if err != nil {
			return res, fmt.Errorf("error, failed to read file")
		}

		filetype := http.DetectContentType(buff)
		if filetype != "image/jpeg" && filetype != "image/png" {
			return res, fmt.Errorf("the provided file format is not allowed. Please upload a JPEG or PNG image")
		}

		_, err = file.Seek(0, io.SeekStart)
		if err != nil {
			return res, fmt.Errorf("error, failed to seek file")
		}

		fileName := utils.RandomString(16) + ".jpg"

		err = os.MkdirAll("./static/products/", os.ModePerm)
		if err != nil {
			return res, fmt.Errorf("error, failed to create directory")
		}

		dst, err := os.Create(fmt.Sprintf("./static/products/%s", fileName))
		if err != nil {
			return res, fmt.Errorf("error, failed to create file")
		}

		defer dst.Close()

		if _, err := io.Copy(dst, file); err != nil {
			return res, fmt.Errorf("error, failed to copy file")
		}

		fileName = "/static/products/" + fileName

		FileName = append(FileName, fileName)
	}

	res = entities.FileUploadRes{
		FilePaths: FileName,
	}

	return res, nil

}
