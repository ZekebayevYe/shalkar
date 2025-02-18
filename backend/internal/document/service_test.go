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

	fileHeader := &multipart.FileHeader{
		Filename: "testfile.txt",
		Size:     123, 
	}

	t.Run("Successful upload", func(t *testing.T) {
		mockRepo.EXPECT().Save(gomock.Any()).Return(nil)

		uploadedFile, err := fileService.UploadFile(fileHeader, "testuser", "admin")

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
		_, err := fileService.UploadFile(fileHeader, "testuser", "user")

		if err == nil {
			t.Fatalf("Expected error, got nil")
		}
		expectedErr := "Only amdin can upload files"
		if err.Error() != expectedErr {
			t.Fatalf("Expected error '%s', got '%v'", expectedErr, err)
		}
	})

	t.Run("Error saving to repository", func(t *testing.T) {
		mockRepo.EXPECT().Save(gomock.Any()).Return(errors.New("database error"))

		_, err := fileService.UploadFile(fileHeader, "testuser", "admin")

		if err == nil {
			t.Fatalf("Expected error, got nil")
		}
		expectedErr := "Error in saving data in DB: database error"
		if err.Error() != expectedErr {
			t.Fatalf("Expected error '%s', got '%v'", expectedErr, err)
		}
	})
}