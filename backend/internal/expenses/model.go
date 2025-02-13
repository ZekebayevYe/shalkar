package expenses

type Expense struct {
	ID          uint    `json:"id" gorm:"primaryKey"`
	UserID      int     `json:"user_id" gorm:"index"`
	ColdWater   float64 `json:"cold_water"`
	HotWater    float64 `json:"hot_water"`
	Heating     float64 `json:"heating"`
	Gas         float64 `json:"gas"`
	Electricity float64 `json:"electricity"`
	TotalCost   float64 `json:"total_cost"`
}
