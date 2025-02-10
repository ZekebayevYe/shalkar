package expenses

const (
	ColdWaterRate      = 100.0  // тг/м**3
	HotWaterRate       = 250.0  // тг/м**3
	HeatingRate        = 200.0  // тг/м**2
	GasRate            = 70.0   // тг/м**3
	ElectricityRate    = 20.0   // тг/кВт*ч
)

type ExpenseService interface {
	CalculateAndSave(userID int, input Expense) (Expense, error)
	GetUserExpenses(userID int) ([]Expense, error)
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
