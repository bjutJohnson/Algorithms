package probabilistic

import (
	"fmt"
)

func Test() {
	treasure := []int{0, 0, 0, 1, 0, 0, 1, 0, 0, 1}
	lvs := make([]int, 0)
	mcs := make([]int, 0)

	for i := 0; i < 10; i++ {
		lvs = append(lvs, LasVegas(treasure, 1))
		mcs = append(mcs, MC(treasure, 5, 1))
	}
	//lv := probabilistic.LasVegas(treasure, 1)
	//mc := probabilistic.MC(treasure, 5, 1)

	fmt.Println("lvs: ", lvs, ", mcs: ", mcs)
}
