package utils

import "math/rand/v2"

func RandIntInRange(minIncl, maxIncl int) int {
	if minIncl >= maxIncl {
		return minIncl
	}
	return rand.IntN(maxIncl-minIncl+1) + minIncl
}

func RandProbability(prob float64) bool {
	return rand.Float64() < prob
}

func RandChoice[T any](choices []T) (T, bool) {
	if len(choices) == 0 {
		var zero T
		return zero, false
	}
	return choices[rand.IntN(len(choices))], true
}

func Shuffled[T any](slice []T) []T {
	shuffled := make([]T, len(slice))
	copy(shuffled, slice)
	rand.Shuffle(len(shuffled), func(i, j int) {
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})

	return shuffled
}
