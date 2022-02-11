// Package weather contains utils functions that returning the forecast of a city.
package weather

// CurrentCondition is the weather of current location.
var CurrentCondition string

//CurrentLocation is the current city.
var CurrentLocation string

// Forecast is the util function that returning the forecast of a city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
