package news

type Events struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Date        string `json:"date"`
	Likes       int    `json:"likes"`
	Dislikes    int    `json:"dislikes"`
}
