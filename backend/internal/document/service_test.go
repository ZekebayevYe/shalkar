package document_test

import (
	"errors"
	"mime/multipart"
	"testing"

	"github.com/golang/mock/gomock"
	"UMS/internal/document"
	"UMS/internal/document/mocks"
)

func TestFileService_UploadFile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockFileRepository(ctrl)
	fileService := document.NewFileService(mockRepo)

	// Фиктивный заголовок файла
	fileHeader := &multipart.FileHeader{
		Filename: "testfile.txt",
		Size:     123, // Фиктивный размер файла
	}

	t.Run("Successful upload", func(t *testing.T) {
		// Ожидаем, что метод Save будет вызван с любым аргументом и вернет nil
		mockRepo.EXPECT().Save(gomock.Any()).Return(nil)

		// Вызов метода UploadFile
		uploadedFile, err := fileService.UploadFile(fileHeader, "testuser", "admin")

		// Проверка результата
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		if uploadedFile == nil {
			t.Fatalf("Expected uploaded file, got nil")
		}
		if uploadedFile.Name != "testfile.txt" {
			t.Fatalf("Expected filename 'testfile.txt', got %s", uploadedFile.Name)
		}
	})

	t.Run("Non-admin role", func(t *testing.T) {
		// Вызов метода UploadFile с ролью, отличной от "admin"
		_, err := fileService.UploadFile(fileHeader, "testuser", "user")

		// Проверка ошибки
		if err == nil {
			t.Fatalf("Expected error, got nil")
		}
		expectedErr := "Only amdin can upload files"
		if err.Error() != expectedErr {
			t.Fatalf("Expected error '%s', got '%v'", expectedErr, err)
		}
	})

	t.Run("Error saving to repository", func(t *testing.T) {
		// Ожидаем, что метод Save вернет ошибку
		mockRepo.EXPECT().Save(gomock.Any()).Return(errors.New("database error"))

		// Вызов метода UploadFile
		_, err := fileService.UploadFile(fileHeader, "testuser", "admin")

		// Проверка ошибки
		if err == nil {
			t.Fatalf("Expected error, got nil")
		}
		expectedErr := "Error in saving data in DB: database error"
		if err.Error() != expectedErr {
			t.Fatalf("Expected error '%s', got '%v'", expectedErr, err)
		}
	})
}