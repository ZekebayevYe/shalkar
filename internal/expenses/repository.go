package expenses

import (
	"database/sql"
	"log"
)

type Repository struct {
	DB *sql.DB
}

// сохраняем итоги в бд
func (r *Repository) SaveTotalCost(userID int, totalCost float64) error {
	query := `INSERT INTO calculations (user_id, total_cost) VALUES ($1, $2)`
	_, err := r.DB.Exec(query, userID, totalCost)
	if err != nil {
		log.Printf("Ошибка сохранения итоговой суммы: %v", err)
		return err
	}
	return nil
}
