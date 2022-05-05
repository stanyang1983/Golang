package main

import "fmt"

func main() {
	//break 強制中斷
	fmt.Println("X迴圈break")
	var x int
	for x = 0; x < 5; x++ {
		if x == 3 {
			break
		}
		fmt.Println(x)
	}

	//continue 強迫回到開頭
	fmt.Println("Y迴圈continue")
	var y int
	for y = 0; y < 9; y++ {
		// if y == 3 {
		// 	continue
		// }
		if y%2 == 0 { // y是偶數,不印出來
			continue
		}
		fmt.Println(y)
	}
	fmt.Println("輸入數字,直到使用者輸入0,取的總和並結束程式")
	//不斷讓使用者輸入數字,直到使用者輸入0,結束程式
	var result int = 0
	for true { //無窮迴圈
		var n int
		fmt.Scanln(&n)
		if n == 0 {
			break
		}
		result += n
	}
	fmt.Println("總和", result)
}
