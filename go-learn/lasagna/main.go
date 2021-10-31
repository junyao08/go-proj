package main

import "fmt"

var OvenTime = 40

func main() {
	fmt.Println(ElapsedTime(3, 10))
}

func RemainingOvenTime(minute int) int {
	ovenTime := OvenTime

	return ovenTime - minute
}

func PreparationTime(layer int) int {
	totalMinutes := layer * 2

	return totalMinutes
}

func ElapsedTime(layer int, minuteInOven int) int {
	totalMinuteInOven := PreparationTime(layer) + minuteInOven

	return totalMinuteInOven
}
