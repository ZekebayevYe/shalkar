package expenses

import "time"

type Expense struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	UserID      int       `json:"user_id" gorm:"not null;index"`
	ColdWater   float64   `json:"cold_water" gorm:"column:cold_water"`
	HotWater    float64   `json:"hot_water" gorm:"column:hot_water"`
	Heating     float64   `json:"heating" gorm:"column:heating"`
	Gas         float64   `json:"gas" gorm:"column:gas"`
	Electricity float64   `json:"electricity" gorm:"column:electricity"`
	TotalCost   float64   `json:"total_cost" gorm:"not null"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
}
