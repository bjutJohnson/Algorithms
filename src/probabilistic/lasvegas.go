package probabilistic

import (
	"fmt"
	"math/rand"
	"time"
)

// 假设有10个山头，其中5个有宝藏，每个山头有一个精灵守护，要进山寻宝，需给精灵3个金币，找到第一个宝藏后，所有的精灵不允许再继续寻宝
// slice中1表示有宝藏，0表示没有保障
func LasVegas(input []int, key int) int {
	retCount := 0
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	idxs := make([]int, 0)
	for i, _ := range input {
		idxs = append(idxs, i)
	}

	for count := 0; count < len(input); count++ {
		fmt.Println("len(idxs):", len(idxs))
		i := r.Intn(len(idxs))
		retCount++
		if input[idxs[i]] == key {
			fmt.Println("===========")
			return 10 - 3*retCount
		}
		idxs = lefts(idxs, idxs[i])
	}
	return 0 - 3*retCount
}

func lefts(idxs []int, idx int) (ret []int) {
	ret = make([]int, 0)

	for i, v := range idxs {
		if v == idx {
			if i == 0 {
				ret = append(ret, idxs[1:]...)
			} else if i == len(idxs)-1 {
				ret = append(ret, idxs[0:i]...)
			} else {
				pre := idxs[0:i]
				post := idxs[i+1:]
				ret = append(ret, pre...)
				ret = append(ret, post...)
			}
		}
	}
	return ret
}
