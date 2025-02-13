package expenses

import (
	"log"

	"gorm.io/gorm"
)

type ExpenseRepository interface {
	Save(expense *Expense) error
	GetByUserID(userID int) ([]Expense, error)
}

type expenseRepo struct {
	db *gorm.DB
}

func NewExpenseRepository(db *gorm.DB) ExpenseRepository {
	return &expenseRepo{db: db}
}

func (r *expenseRepo) Save(expense *Expense) error {
	log.Println("📌 Сохранение расходов:", expense)
	err := r.db.Create(expense).Error
	if err != nil {
	  log.Println("❌ Ошибка сохранения в БД:", err)
	}
	return err
  }

func (r *expenseRepo) GetByUserID(userID int) ([]Expense, error) {
	var expenses []Expense
	err := r.db.Where("user_id = ?", userID).Find(&expenses).Error
	return expenses, err
}
