package document

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"time"
)

type FileService interface {
    UploadFile(fileHeader *multipart.FileHeader, username string, role string) (*File, error)
    GetFile(id uint) (*File, error)
    ListFiles() ([]File, error)
}

type fileService struct {
    repo FileRepository
}

func NewFileService(repo FileRepository) FileService {
    return &fileService{repo: repo}
}

func (s *fileService) UploadFile(fileHeader *multipart.FileHeader, username string, role string) (*File, error) {
    if role != "admin" {
        return nil, errors.New("доступ запрещен: только админы могут загружать файлы")
    }

    uploadPath := "uploads/"
    if err := os.MkdirAll(uploadPath, os.ModePerm); err != nil {
        return nil, fmt.Errorf("не удалось создать папку для загрузки: %v", err)
    }

    filePath := fmt.Sprintf("%s%s", uploadPath, fileHeader.Filename)

    src, err := fileHeader.Open()
    if err != nil {
        return nil, fmt.Errorf("ошибка открытия файла: %v", err)
    }
    defer src.Close()

    dst, err := os.Create(filePath)
    if err != nil {
        return nil, fmt.Errorf("ошибка создания файла: %v", err)
    }
    defer dst.Close()

    if _, err := io.Copy(dst, src); err != nil {
        return nil, fmt.Errorf("ошибка сохранения файла: %v", err)
    }

    file := &File{
        Name:       fileHeader.Filename,
        Path:       filePath,
        UploadedBy: username,
        CreatedAt:  time.Now(),
    }

    if err := s.repo.Save(file); err != nil {
        return nil, fmt.Errorf("ошибка сохранения информации о файле в БД: %v", err)
    }

    return file, nil
}

func (s *fileService) GetFile(id uint) (*File, error) {
    return s.repo.GetByID(id)
}

func (s *fileService) ListFiles() ([]File, error) {
    return s.repo.GetAll()
}
