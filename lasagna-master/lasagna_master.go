package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, timeTakenForEachLayer int) int {
	if timeTakenForEachLayer <= 0 {
		timeTakenForEachLayer = 2
	}
	return len(layers) * timeTakenForEachLayer
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	var noodle int = 0
	var sauce float64 = 0
	for _, ingredient := range layers {
		if ingredient == "noodles" {
			noodle += 50
		}

		if ingredient == "sauce" {
			sauce += 0.2
		}
	}

	return noodle, sauce

}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendList []string, myList []string) {
	myList[len(myList)-1] = friendList[len(friendList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
	var newQuantities []float64 = make([]float64, len(quantities))
	for i, q := range quantities {
		newQuantities[i] = q / float64(2) * float64(portions)
	}

	return newQuantities
}
