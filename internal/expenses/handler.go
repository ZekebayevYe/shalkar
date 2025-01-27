package expenses

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func CalculateHandler(repo *Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var input ExpenseInput
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		// получаем юзер айди
		userIDStr := r.URL.Query().Get("user_id")
		userID, err := strconv.Atoi(userIDStr)
		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}

		// считаем сумму
		result := CalculateExpenses(input)

		// результат в базу данных кидаем
		if saveErr := repo.SaveTotalCost(userID, result.TotalCost); saveErr != nil {
			http.Error(w, "Failed to save total cost", http.StatusInternalServerError)
			return
		}

		// юзеру резы отправляем
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]float64{"total_cost": result.TotalCost})
	}
}
