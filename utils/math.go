package utils

import "math"

// RoundFloat rounds the float value to 2 decimals
func RoundFloat(v float64) float64 {
	return math.Round(v*100) / 100
}
