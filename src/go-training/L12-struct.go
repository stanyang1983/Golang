package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	//struct 結構
	//定義結構   >  實體化
	//定義結構: 結構名稱 包含的資料欄位

	//2.實體化結構

	fmt.Println("struct 結構")

	John := Person{"John", 18}
	fmt.Println(John.name, John.age)

	p2 := Person{name: "Mary", age: 30}
	p2.name = "Sue"
	fmt.Println(p2.name, p2.age)

}
