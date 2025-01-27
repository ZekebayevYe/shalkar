package expenses

const (
	ColdWaterRate      = 100.0  // тг/м**3
	HotWaterRate       = 250.0  // тг/м**3
	HeatingRate        = 200.0  // тг/м**2
	GasRate            = 70.0   // тг/м**3
	ElectricityRate    = 20.0   // тг/кВт*ч
	GarbageRatePerPerson = 500.0 // тг/чел
)

func CalculateExpenses(input ExpenseInput) ExpenseResult {
	var result ExpenseResult

	if input.UseColdWater {
		result.ColdWaterCost = input.ColdWater * ColdWaterRate
	}
	if input.UseHotWater {
		result.HotWaterCost = input.HotWater * HotWaterRate
	}
	if input.UseHeating {
		result.HeatingCost = input.Heating * HeatingRate
	}
	if input.UseGas {
		result.GasCost = input.Gas * GasRate
	}
	if input.UseElectricity {
		result.ElectricityCost = input.Electricity * ElectricityRate
	}

	result.GarbageCost = float64(input.People) * GarbageRatePerPerson

	result.TotalCost = result.ColdWaterCost + result.HotWaterCost +
		result.HeatingCost + result.GasCost + result.ElectricityCost +
		result.GarbageCost

	return result
}
