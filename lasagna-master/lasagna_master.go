package lasagna

// PreparationTime returns the time that takes to prepare all layers.
func PreparationTime(layers []string, timePerLayer int) int {
	if timePerLayer == 0 {
		timePerLayer = 2
	}

	layersNum := len(layers)

	return layersNum * timePerLayer
}

// Quantities returns the exact amount of noodle and sauce that needs.
func Quantities(layers []string) (int, float64) {
	noodlePerLayer := 50
	saucePerLayer := 0.2

	var noodleLayersNum int
	var sauceLayersNum float64

	for i := 0; i < len(layers); i++ {
		if layers[i] == "noodles" {
			noodleLayersNum++
		}

		if layers[i] == "sauce" {
			sauceLayersNum++
		}
	}

	noodle := noodleLayersNum * noodlePerLayer
	sauce := sauceLayersNum * saucePerLayer

	return noodle, sauce
}

// AddSecretIngredient sets the last element of  the friend list as the last element of my list.
func AddSecretIngredient(friendList, myList []string) {
	secretIngredient := friendList[len(friendList)-1]

	myList[len(myList)-1] = secretIngredient
}

// ScaleRecipe returns the exact amount for different numbers of portions as a new slice.
func ScaleRecipe(quantities []float64, portionsNum int) []float64 {
	var quantitiesPerOne []float64
	var totalQuantities []float64

	for i := 0; i < len(quantities); i++ {
		quantitiesPerOne = append(quantitiesPerOne, quantities[i]/float64(2))
	}

	for i := 0; i < len(quantitiesPerOne); i++ {
		totalQuantities = append(totalQuantities, quantitiesPerOne[i]*float64(portionsNum))
	}

	return totalQuantities
}
