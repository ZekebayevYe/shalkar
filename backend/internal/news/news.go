package news

type News struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
	Likes       int    `json:"likes"`
	Dislikes    int    `json:"dislikes"`
}
