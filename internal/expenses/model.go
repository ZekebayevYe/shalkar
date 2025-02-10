package expenses

type ExpenseInput struct {
	ColdWater    float64 `json:"cold_water"`
	HotWater     float64 `json:"hot_water"`
	Heating      float64 `json:"heating"`
	Gas          float64 `json:"gas"`
	Electricity  float64 `json:"electricity"`
	People       int     `json:"people"`
	UseColdWater bool    `json:"use_cold_water"`
	UseHotWater  bool    `json:"use_hot_water"`
	UseHeating   bool    `json:"use_heating"`
	UseGas       bool    `json:"use_gas"`
	UseElectricity bool  `json:"use_electricity"`
}

type ExpenseResult struct {
	ColdWaterCost   float64 `json:"cold_water_cost"`
	HotWaterCost    float64 `json:"hot_water_cost"`
	HeatingCost     float64 `json:"heating_cost"`
	GasCost         float64 `json:"gas_cost"`
	ElectricityCost float64 `json:"electricity_cost"`
	GarbageCost     float64 `json:"garbage_cost"`
	TotalCost       float64 `json:"total_cost"`
}
