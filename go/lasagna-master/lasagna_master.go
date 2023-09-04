package lasagna

func PreparationTime(layers []string, time int) int {
	if time == 0 {
		time = 2
	}
	return len(layers) * time
}

func Quantities(layers []string) (noodles int, sauce float64) {
	for i := range layers {
		if layers[i] == "noodles" {
			noodles += 50
		}
		if layers[i] == "sauce" {
			sauce += 0.2
		}
	}
	return
}

func AddSecretIngredient(friendsList []string, myList []string) []string {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
	return myList
}

func ScaleRecipe(quantities []float64, portions int) (scaledQuantities []float64) {
	for i := range quantities {
		scaledQuantities = append(scaledQuantities, quantities[i]*(float64(portions)/2))
	}
	return
}
