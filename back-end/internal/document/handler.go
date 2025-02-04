package document

import (
	"fmt"
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
        c.JSON(http.StatusUnauthorized, gin.H{"error": "user_id не найден"})
        return
    }
    roleVal, exists := c.Get("role")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Роль не найдена"})
        return
    }

    userID, ok := userIDVal.(int)
    if !ok {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка преобразования user_id"})
        return
    }
    role, ok := roleVal.(string)
    if !ok {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка преобразования роли"})
        return
    }

    file, err := c.FormFile("file")
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка при получении файла: " + err.Error()})
        return
    }

    fmt.Println("Перед сохранением файла userID =", userID)

    uploadedFile, err := h.service.UploadFile(file, strconv.Itoa(userID), role)
    if err != nil {
        c.JSON(http.StatusForbidden, gin.H{"error": "Ошибка загрузки файла: " + err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Файл загружен", "file": uploadedFile})
}

func (h *FileHandler) DownloadFile(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID файла"})
        return
    }

    file, err := h.service.GetFile(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Файл не найден"})
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
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении списка файлов"})
        return
    }
    c.JSON(http.StatusOK, files)
}