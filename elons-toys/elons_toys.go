package elon

import "fmt"

// Drive() drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func (car *Car) Drive() {
	if car.battery-car.batteryDrain >= 0 {
		car.battery = car.battery - car.batteryDrain
		car.distance += car.speed
	}
}

// DisplayDistance() returns the distance as displayed on the LED display as a string.
func (car *Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", car.distance)
}

// DisplayBattery() returns the battery percentage as displayed on the LED display as a string.
func (car *Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", car.battery)
}

// CanFinish() checks if a car is able to finish a race.
func (car *Car) CanFinish(trackDistance int) bool {
	drivenTimes := car.battery / car.batteryDrain
	maxDist := car.speed * drivenTimes

	return maxDist >= trackDistance

}
