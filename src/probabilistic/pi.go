package probabilistic

import(
	"math/rand"
//	"math"
	"fmt"
//	"time"
)

func ComputePi(num int) float64{
	hit := 0

	for i := 0 ; i < num; i++{
		//s1 := rand.NewSource(time.Now().Unix())
		//r1 := rand.New(s1)

		x := rand.Float64()

		//s2 := rand.NewSource(time.Now().Unix())
		//r2 := rand.New(s2)
		y := rand.Float64()

		//fmt.Println(x, y)

		if x * x + y * y <= 1.0{
			hit++
		}
	}
	fmt.Println("hit = ", hit)
	return 4 * float64(hit) / float64(num)
}

func Test_ComputePi(){
	tryNums := []int{100, 200, 500, 1000, 2000, 5000, 10000, 20000, 50000, 100000, 200000, 500000, 1000000, 2000000, 5000000, 10000000, 20000000, 50000000, 100000000}

	for _, v :=range(tryNums){
		fmt.Println("n = ", v, ": pi = ", ComputePi(v))
	}
}