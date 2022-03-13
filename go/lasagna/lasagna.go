package lasagna

// TODO: define the 'OvenTime' constant

var OvenTime = 40

// RemainingOvenTime returns the remaining minutes based on the `actual` minutes already in the oven.
func RemainingOvenTime(actualMinutesInOven int) int {
	remainingTime := OvenTime - actualMinutesInOven

	return remainingTime

}

// PreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.
func PreparationTime(numberOfLayers int) int {
	preparationTime := 2 * numberOfLayers
	return preparationTime

}

// ElapsedTime calculates the total time needed to create and bake a lasagna.
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	preparation := PreparationTime(numberOfLayers)

	elapsed := preparation + actualMinutesInOven

	return elapsed

}
