// File: pkg/utils/percentage_difference.go

package utils

// CalculatePercentageChange calculates the percentage change between two values.
func CalculatePercentageChange(current, previous int) float64 {
	if previous == 0 {
		if current == 0 {
			return 0
		}
		return 100 // If previous is 0 and current is not 0, return 100% change
	}
	return (float64(current-previous) / float64(previous)) * 100
}
