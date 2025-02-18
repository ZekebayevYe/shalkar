package feedback

import "time"

type Feedback struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id" gorm:"not null;index"`
	Category  string    `json:"category" gorm:"not null"`
	Rating    int       `json:"rating" gorm:"not null;check:rating >= 0 AND rating <= 5"`
	Comment   string    `json:"comment" gorm:"size:500"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}

var AvailableCategories = []string{
	"Website Functionality",
	"Janitorial Services",
	"Plumbing",
	"Customer Support",
	"Electricity Issues",
	"Building Maintenance",
}
