// Package weather allows to tools to show weather in current condition and current location.
package weather

var (
	// CurrentCondition represents condition in a string.
	CurrentCondition string
	// CurrentLocation represents location in a string.
	CurrentLocation string
)

// Forecast returns current location and current condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
