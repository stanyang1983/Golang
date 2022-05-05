package main

import "fmt"

func main() {
	//Pass by Pointer 指標參數
	//指標 儲存資料的記憶體位址

	//函式參數傳遞!
	//呼叫函式時,資料如何透過參數傳遞

	var x int = 10

	add(&x) //&x 傳遞x在記憶體中位址
	fmt.Println("Main Functiomn", x)

	// 和使用者要求輸入, 運用到指標參數  pass by pointer

	var msg string
	fmt.Scanln(&msg) //傳遞字串變數的指標(記憶體位址)  &msg
	fmt.Println(msg)
}

func add(xPtr *int) {
	*xPtr = *xPtr + 10 //*xPtr 轉譯地址回值
	fmt.Println("Add Functiomn", *xPtr)
}
