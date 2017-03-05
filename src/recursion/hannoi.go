package recursion

import(
	"fmt"
)

func Hannoi(n int, first, second, third string){
	if n == 1{
		fmt.Println(first, "->", third)
	}else{
		Hannoi(n - 1, first, third, second)
		fmt.Println(first, "->", third)
		Hannoi(n - 1, second, first, third)
	}
}

func TestHannoi(){
	fmt.Println("3个圆盘的情况：")
	Hannoi(3, "A", "B", "C")

	fmt.Println("4个圆盘的情况：")
	Hannoi(4, "A", "B", "C")

	fmt.Println("5个圆盘的情况：")
	Hannoi(5, "A", "B", "C")
}