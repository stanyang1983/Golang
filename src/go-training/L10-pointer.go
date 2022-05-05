package main

import "fmt"

func main() {

	var x int = 5
	// &x 取得記憶體的位址!!! 酷!
	fmt.Println("原來的資料:", x)
	fmt.Println("資料在記憶體中的地址:", &x)

	//指標資料型態: *資料型態
	var xPtr *int = &x // *int = &x
	fmt.Println(xPtr)

	//反解指標變數
	fmt.Println("反解指標回原來的資料 :", *xPtr) //*xPtr

	var txt string = "hello"
	fmt.Println(txt)
	var txtPtr *string = &txt
	fmt.Println(txtPtr)
	var txt2 = *txtPtr
	fmt.Println(txt2)

}

/* 指標
建立資料變數 > 取得記憶體位址
存放到指標變數 > 反解指標變數
*/
