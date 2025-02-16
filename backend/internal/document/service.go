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
    DeleteFile(id uint) error
}

type fileService struct {
    repo FileRepository
}

func NewFileService(repo FileRepository) FileService {
    return &fileService{repo: repo}
}

func (s *fileService) UploadFile(fileHeader *multipart.FileHeader, username string, role string) (*File, error) {
    if role != "admin" {
        return nil, errors.New("Only amdin can upload files")
    }

    uploadPath := "uploads/"
    if err := os.MkdirAll(uploadPath, os.ModePerm); err != nil {
        return nil, fmt.Errorf("Error in creating file to upload: %v", err)
    }

    filePath := fmt.Sprintf("%s%s", uploadPath, fileHeader.Filename)

    src, err := fileHeader.Open()
    if err != nil {
        return nil, fmt.Errorf("Error in opening file: %v", err)
    }
    defer src.Close()

    dst, err := os.Create(filePath)
    if err != nil {
        return nil, fmt.Errorf(`Error creating file: %v`, err)
    }
    defer dst.Close()

    if _, err := io.Copy(dst, src); err != nil {
        return nil, fmt.Errorf(`Error in saving file: %v`, err)
    }

    file := &File{
        Name:       fileHeader.Filename,
        Path:       filePath,
        UploadedBy: username,
        CreatedAt:  time.Now(),
    }

    if err := s.repo.Save(file); err != nil {
        return nil, fmt.Errorf(`Error in saving data in DB: %v`, err)
    }

    return file, nil
}

func (s *fileService) GetFile(id uint) (*File, error) {
    return s.repo.GetByID(id)
}

func (s *fileService) DeleteFile(id uint) error {
    file, err := s.repo.GetByID(id) 
    if err != nil {
        return fmt.Errorf("file not found: %v", err)
    }

    if err := os.Remove(file.Path); err != nil {
        return fmt.Errorf("error deleting file from storage: %v", err)
    }

    if err := s.repo.DeleteFile(id); err != nil {
        return fmt.Errorf("error deleting file from database: %v", err)
    }

    return nil
}

func (s *fileService) ListFiles() ([]File, error) {
    return s.repo.GetAll()
}
