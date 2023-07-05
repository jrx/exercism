// Package weather provides a forecast of the current weather conditions in various cities.
package weather

// CurrentCondition represents the condition of the weather forecast.
var CurrentCondition string

// CurrentLocation represents the location of the weather forecast.
var CurrentLocation string

// Forecast returns a string value of the current weather location and condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
