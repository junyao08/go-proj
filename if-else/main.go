package main

import "fmt"

func main() {

	//s := SuccessRate(0)
	//fmt.Println(s)

	fmt.Println(CalculateProductionRatePerHour(9))

	//fmt.Println(CalculateProductionRatePerMinute(5))

	fmt.Println(CalculateLimitedProductionRatePerHour(5, 1000.0))
	fmt.Println(CalculateLimitedProductionRatePerHour(10, 1000.0))
}

// SuccessRate is used to calculate the ratio of an item being created without
// error for a given speed
func SuccessRate(speed int) float64 {

	if speed >= 0 && speed <= 4 {
		return 1.0
	} else if speed <= 8 {
		return 0.9
	} else if speed <= 10 {
		return 0.77
	} else {
		return 0.0
	}
}

// CalculateProductionRatePerHour for the assembly line, taking into account
// its success rate
func CalculateProductionRatePerHour(speed int) float64 {
	productionRatePerHour := float64(speed) * 221 * SuccessRate(speed)
	return float64(productionRatePerHour)
}

// CalculateProductionRatePerMinute describes how many working items are
// produced by the assembly line every minute
func CalculateProductionRatePerMinute(speed int) int {
	prodRatePerMinute := int(CalculateProductionRatePerHour(speed)) / 60
	return prodRatePerMinute
}

// CalculateLimitedProductionRatePerHour describes how many working items are
// produced per hour with an upper limit on how many can be produced per hour
func CalculateLimitedProductionRatePerHour(speed int, limit float64) float64 {

	if CalculateProductionRatePerHour(speed) > limit {
		return limit
	}

	return CalculateProductionRatePerHour(speed)
}
