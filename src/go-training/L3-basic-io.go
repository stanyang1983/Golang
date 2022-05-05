package main

import "fmt"

func main() {
	// 從終端介面輸入 fmt.ScanIn(&變數名稱,&變數名稱...)
	// &變數名稱: 取得變數的指標(Pointer)
	// 換行輸出 fmt.Println()
	// 不換行輸出 fmt.Print()
	// EX fmt.Print(3,"Hello",true)

	// var x int
	// fmt.Print("輸入一個數字:")
	// fmt.Scanln(&x)         //將輸入的資料存放到x變數
	// fmt.Println("數值為:", x) //輸出資料

	//輸入兩個數字  取的相乘的結果 && 輸出
	var x int
	var y int
	//方式一

	fmt.Print("輸入第一個數字:")
	fmt.Scanln(&x)

	fmt.Print("輸入第二個數字:")
	fmt.Scanln(&y)

	//方式二
	// fmt.Print("輸入兩個數字,用空格隔開:")
	// fmt.Scanln(&x, &y)

	var result int = x * y
	fmt.Print("相乘的結果為:")
	fmt.Println(result)
}
