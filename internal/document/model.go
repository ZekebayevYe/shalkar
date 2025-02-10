package document

import "time"

type File struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
    Name      string    `json:"name"`
    Path      string    `json:"path"`
    UploadedBy string   `json:"uploaded_by"`
    CreatedAt time.Time `json:"created_at"`
}
