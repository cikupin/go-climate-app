package libraries

import "math"

// RoundNumber - round tumber to 2 decimals
// Example - 12.3456 to 12.34
func RoundNumber(x float64) float64 {
	return math.Round(x/100) * 100
}
