package main

import "fmt"

func main() {

	//算數:+,-,*,/
	//指定運算: = ,+= , -= ,*= ,/=
	//單元運算: ++ , --
	//比較運算: > , < , >= ,<= ,==
	//邏輯: !, && ,||

	//指定運算
	var x int
	x = 5
	x += 2 // x = x+2

	fmt.Println(x)

	var y bool

	y = 4 != 3
	fmt.Println(y)
}
