package components

func factorial(length int) int {
	if length <= 1 {
		return 1
	}
	return length * factorial(length-1)
}

// Function to calculate the number of combinations
func CalculateCombinations(length, r int) int {
	numerator := factorial(length)
	denominator := factorial(r) * factorial(length-r)
	return numerator / denominator
}