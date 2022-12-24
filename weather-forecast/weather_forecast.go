// Package weather shows weather condition in a specific location.
package weather

// CurrentCondition variable shows how the weather is right now.
var CurrentCondition string

// CurrentLocation variable shows where you are right now.
var CurrentLocation string

// Forecast function shows the weather of where you are.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
