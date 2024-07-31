package utils

import (
	"math/rand"
	"time"
)

func Roll(num, dice int64) []int64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	results := make([]int64, num)
	for i := 0; i < int(num); i++ {
		results[i] = r.Int63n(dice) + 1
	}

	return results
}
