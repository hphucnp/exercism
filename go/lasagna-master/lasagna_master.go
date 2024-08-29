package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avgTimePerLayer int) int {
	if avgTimePerLayer == 0 {
		return len(layers) * 2
	}
	return len(layers) * avgTimePerLayer
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	noodles_quantity := 0
	sauce_quantity := 0.0
	for _, layer := range layers {
		if layer == "noodles" {
			noodles_quantity += 50
		} else if layer == "sauce" {
			sauce_quantity += 0.2
		}
	}
	return noodles_quantity, sauce_quantity
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) {
	//
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, numberOfPortions int) []float64 {
	scaledQuantities := make([]float64, len(quantities))
	for i := 0;i < len(quantities); i++ {
		scaledQuantities[i] = quantities[i] * float64(numberOfPortions) / 2
	}
	return scaledQuantities
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
// 
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more 
// functionality.
