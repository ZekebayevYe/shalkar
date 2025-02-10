package document

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FileHandler struct {
    service FileService
}

func NewFileHandler(service FileService) *FileHandler {
    return &FileHandler{service: service}
}

func (h *FileHandler) UploadFile(c *gin.Context) {
    userIDVal, exists := c.Get("user_id")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "user_id not found"})
        return
    }
    roleVal, exists := c.Get("role")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Роль not found"})
        return
    }

    userID, ok := userIDVal.(int)
    if !ok {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error in user_id"})
        return
    }
    role, ok := roleVal.(string)
    if !ok {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "role"})
        return
    }

    file, err := c.FormFile("file")
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Taking file: " + err.Error()})
        return
    }

    uploadedFile, err := h.service.UploadFile(file, strconv.Itoa(userID), role)
    if err != nil {
        c.JSON(http.StatusForbidden, gin.H{"error": "Uploading file: " + err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "file successful", "file": uploadedFile})
}

func (h *FileHandler) DeleteFile(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "incorrect id"})
        return
    }

    err = h.service.DeleteFile(uint(id))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting file: " + err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "File deleted successfully"})
}


func (h *FileHandler) DownloadFile(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "incorrect id"})
        return
    }

    file, err := h.service.GetFile(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "fail not found"})
        return
    }

    c.Header("Content-Disposition", "attachment; filename="+file.Name)
    c.Header("Content-Type", "application/octet-stream")
    c.Header("Content-Type", "image/jpeg") 
    c.File(file.Path)    
}

func (h *FileHandler) ListFiles(c *gin.Context) {
    files, err := h.service.ListFiles()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error in getting files"})
        return
    }
    c.JSON(http.StatusOK, files)
}