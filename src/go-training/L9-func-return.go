package main

import "fmt"

//函式 return 結束函式
func main() {
	show("MDFK")
	show("Hello")

	var x int = multiply(3, 4)
	fmt.Println(x) //接收multiply()的回傳值 打印x

	var y int
	var z int
	y, z = multiply2(5, 6) //接收多個回傳值 打印 y,z
	fmt.Println(y, z)

}

//return
func show(msg string) {
	if msg == "Hello" {
		return
	}
	fmt.Println(msg)
}

func multiply(n1 int, n2 int) int { //<--  int回傳值要帶回傳值的參數
	var result int = n1 * n2
	return result
}

func multiply2(n1 int, n2 int) (int, int) { //<--  int回傳值要帶回傳值的參數
	var result1 int = n1 * n2
	var result2 int = n1 + n2
	return result1, result2
}

//函式回傳值
/*func 函式名稱(參數列表) 回傳值資料型態{
	函式內部程式
	return 回傳值,須符合定義中的資料型態
 }
*/

//多個回傳值
/*func 函式名稱(參數列表) (資料型態,資料型態,...){
	函式內部程式
	return 回傳值,回傳值,...    須符合定義中的資料型態
 }
*/
