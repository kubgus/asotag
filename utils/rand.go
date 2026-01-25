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

func RandWeightedChoice[T comparable](table map[T]int) T {
	totalWeight := 0
	for _, weight := range table {
		totalWeight += weight
	}

	randWeight := RandIntInRange(1, totalWeight)
	currentWeight := 0

	for item, weight := range table {
		currentWeight += weight
		if randWeight <= currentWeight {
			return item
		}
	}

	var zero T
	return zero
}
