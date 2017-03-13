package probabilistic

import (
	"math/rand"
	"time"
)

// 假设有10个山头，其中3个有宝藏，每个山头有一个精灵守护，要进山寻宝，需给精灵1个金币，找到第一个宝藏后，所有的精灵不允许再继续寻宝
// slice中1表示有宝藏，0表示没有保障
func LasVegas(input []int, key int) int {
	retCount := 1
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for count := 0; count < len(input); count++ {
		i := r.Intn(len(input))
		if input[i] == key {
			return 10 - retCount
		} else {
			retCount++
		}
	}
	return 0 - retCount
}
