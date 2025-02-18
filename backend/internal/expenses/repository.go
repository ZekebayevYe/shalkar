package expenses

import (
	"gorm.io/gorm"
)

type ExpenseRepository interface {
	Save(expense *Expense) error
	GetByUserID(userID int) ([]Expense, error)
	GetExpenseByID(expenseID int, userID int) (Expense, error)
}

type expenseRepo struct {
	db *gorm.DB
}

func NewExpenseRepository(db *gorm.DB) ExpenseRepository {
	return &expenseRepo{db: db}
}

func (r *expenseRepo) Save(expense *Expense) error {
	return r.db.Create(expense).Error
}

func (r *expenseRepo) GetByUserID(userID int) ([]Expense, error) {
	var expenses []Expense
	err := r.db.Where("user_id = ?", userID).Order("created_at DESC").Find(&expenses).Error
	return expenses, err
}

func (r *expenseRepo) GetExpenseByID(expenseID int, userID int) (Expense, error) {
	var expense Expense
	err := r.db.Where("id = ? AND user_id = ?", expenseID, userID).First(&expense).Error
	return expense, err
}
