package main

import "fmt"

func main() {

	// var x int
	// fmt.Print("請輸入100:")
	// fmt.Scanln(&x)
	// if x == 100 {
	// 	fmt.Print("輸入正確")
	// } else {
	// 	fmt.Print("輸入錯誤")
	// }
	var money int
	fmt.Print("請問想領多少錢:")
	fmt.Scanln(&money)
	if money < 100 {
		fmt.Print("金額太小")
	} else if money <= 100000 {
		fmt.Print("OK")
	} else {
		fmt.Print("超過額度")
	}
}
