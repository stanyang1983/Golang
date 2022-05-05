package main

import (
	"fmt"
)

//函式
func main() {

	//1.定義(建立)函式---> 2.呼叫函式
	/*func 函式名稱(參數列表){
		函式內部的程式碼
	}
	*/
	//呼叫函式
	show("Hello")
	show("您好")
	multiply(3, 8)
	multiply(5, -8)

	// 計算 1+2+3+....+10 的函示包裝
	sum(20)
	sum(50)
	sum(100)
}

//show函式
func show(msg string) {
	fmt.Println(msg)
}

//乘函式
func multiply(n1 int, n2 int) {
	var result int = n1 * n2
	fmt.Println(result)
}

func sum(max int) {
	var result int = 0
	var n int
	for n = 1; n <= max; n++ {
		result += n //result = result + n

	}
	fmt.Println(result)
}
