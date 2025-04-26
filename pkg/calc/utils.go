package calc

import "math"

// Round float to 5 decimal places
func RoundToDecimal(value float64) float64 {
	return math.Round(value*10000)/10000
}

// Mini function to use in other 'get' functions
func calculateTotal(freqMap map[string]int, compute func(word string, freq int) int) int {
	var total int
	for word, freq := range freqMap {
		total += compute(word, freq)
	}
	return total
}
