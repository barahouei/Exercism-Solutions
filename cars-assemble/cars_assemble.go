package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return (successRate / 100) * float64(productionRate)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	perHour := int((successRate / 100) * float64(productionRate))
	perMinute := perHour / 60

	return perMinute
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	groupsOfCars := carsCount / 10
	remainder := carsCount % 10

	costOfTen := groupsOfCars * 95000
	costOfIndividual := remainder * 10000

	finalCost := costOfTen + costOfIndividual
	finalCostToUint := uint(finalCost)

	return finalCostToUint
}
