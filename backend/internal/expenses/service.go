package expenses

const (
	ColdWaterRate   = 100.0
	HotWaterRate    = 250.0
	HeatingRate     = 200.0
	GasRate         = 70.0
	ElectricityRate = 20.0
)

type ExpenseService interface {
	CalculateAndSave(userID int, input Expense) (Expense, error)
	GetUserExpenses(userID int) ([]Expense, error)
	GetExpenseDetails(expenseID int, userID int) (Expense, error)
}

type expenseService struct {
	repo ExpenseRepository
}

func NewExpenseService(repo ExpenseRepository) ExpenseService {
	return &expenseService{repo: repo}
}

func (s *expenseService) CalculateAndSave(userID int, input Expense) (Expense, error) {
	input.TotalCost = (input.ColdWater * ColdWaterRate) +
		(input.HotWater * HotWaterRate) +
		(input.Heating * HeatingRate) +
		(input.Gas * GasRate) +
		(input.Electricity * ElectricityRate)

	input.UserID = userID

	err := s.repo.Save(&input)
	return input, err
}

func (s *expenseService) GetUserExpenses(userID int) ([]Expense, error) {
	return s.repo.GetByUserID(userID)
}

func (s *expenseService) GetExpenseDetails(expenseID int, userID int) (Expense, error) {
	return s.repo.GetExpenseByID(expenseID, userID)
}
