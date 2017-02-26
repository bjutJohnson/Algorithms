package recursion

import(
	"fmt"
	"time"
)

func Test_Unit(n int){
	//fmt.Print("before Fib: ")
	//fmt.Println(time.Now())
	t1 := time.Now()

	Fib(n)

	t2 := time.Now()

	//fmt.Println("k = ", k)

	fmt.Println(n, "- Recursion Running Time: ", t2.Sub(t1))
	//fmt.Println(time.Now())

	t3 := time.Now()
	 FibNotRec(n)
	t4 := time.Now()

	fmt.Println(n, "- Not Recursion Running Time: ", t4.Sub(t3))

}

func Test(){
	testNums := []int{10, 20, 30, 40, 50}
	for _, v := range(testNums){
		Test_Unit(v)
	}
}