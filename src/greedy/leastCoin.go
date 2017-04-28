package greedy

import (
	"fmt"
)

const (
	ONE  = 1
	TWO  = 2
	FIVE = 5
	TEN  = 10
)

// 获取最少硬币数
// 假设硬币种类有：1分、2分、5分、1角，对于任意金额m，求最少硬币的数量
func LeastCoinNum(money int) int {
	tenNum := money / TEN
	money = money - tenNum*TEN

	fiveNum := money / FIVE
	money = money - fiveNum*FIVE

	twoNum := money / TWO
	oneNum := money - twoNum*TWO

	return tenNum + fiveNum + twoNum + oneNum
}

func TestLeastCoinNum() {
	fmt.Print("5角6分的最少硬币数：")
	fmt.Println(LeastCoinNum(56))

	fmt.Print("1元3角9分的最少硬币数：")
	fmt.Println(LeastCoinNum(139))
}
