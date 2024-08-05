package utils

import (
	"math/rand"
	"time"
)

func Roll(num, dice int) []int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	results := make([]int, num)
	for i := 0; i < num; i++ {
		results[i] = int(r.Int63n(int64(dice))) + 1
	}

	return results
}
