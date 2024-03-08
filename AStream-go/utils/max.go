package utils

import "time"

func MaxFloat64(x, y float64) float64 {
	if x > y {
		return x
	}
	return y
}

func MaxDuration(x, y time.Duration) time.Duration {
	if x > y {
		return x
	}
	return y
}
