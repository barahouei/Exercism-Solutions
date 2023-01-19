package meteorology

import "fmt"

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

func (tu TemperatureUnit) String() string {
	tUnits := []string{"°C", "°F"}

	return tUnits[tu]
}

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

func (t Temperature) String() string {
	return fmt.Sprintf("%d %s", t.degree, t.unit)
}

type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

func (su SpeedUnit) String() string {
	sUnits := []string{"km/h", "mph"}

	return sUnits[su]
}

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

func (s Speed) String() string {
	return fmt.Sprintf("%d %s", s.magnitude, s.unit)
}

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

func (m MeteorologyData) String() string {
	// Example of the returned string = "San Francisco: 57 °F, Wind NW at 19 mph, 60% Humidity"
	s := fmt.Sprintf("%s: %d %s, Wind %s at %d %s, %d%% Humidity",
		m.location,
		m.temperature.degree,
		m.temperature.unit,
		m.windDirection,
		m.windSpeed.magnitude,
		m.windSpeed.unit,
		m.humidity)

	return s
}
