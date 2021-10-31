package main

import (
	"fmt"
)

func main() {

	speed := 5
	batteryDrain := 2
	newCar := NewCar(speed, batteryDrain)
	fmt.Println(newCar)

}

// TODO: define the 'Car' type struct
type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	car := Car{
		speed:        speed,
		batteryDrain: batteryDrain,
		battery:      100,
		distance:     0,
	}

	return car
}

// TODO: define the 'Track' type struct
type Track struct {
	distance int
}

// NewTrack created a new track
func NewTrack(distance int) Track {
	track := Track{distance: distance}
	return track
}

// Drive drives the car one time. If there is not enough battery to drive on more time,
// the car will not move.
func Drive(car Car) Car {
	panic("Please implement the Drive function")
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	panic("Please implement the CanFinish function")
}
