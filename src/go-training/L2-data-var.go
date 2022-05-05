package main

import "fmt"

//載入內建的fmt封包 用來做基本輸出輸入

func main() { //程式的進入點
	//fmt.Println(資料,資料,...)

	/*
		fmt.Println(3)      //整數 int
		fmt.Println(3.1415) //浮點數  float64
		fmt.Println("測試")   //字串 string
		fmt.Println(true)   //布林值 bool
		fmt.Println('a')    //字符 rune
	*/

	//變數的使用
	var x int = 4 //宣告變數

	fmt.Println(x)
	x = 10
	fmt.Println(x) //可以改變變數

	var f float64 = 3.1415
	fmt.Println(f)

	var s string = "我是字串"
	fmt.Println(s)

	var b bool = true
	fmt.Println(b)

	var r rune = 'b'
	fmt.Println(r)

}
