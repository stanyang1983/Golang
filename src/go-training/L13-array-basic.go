package main

import "fmt"

func main() {

	fmt.Println("array 陣列")

	//整數陣列--------------------------------
	arr1 := [4]int{1, 3, 5, 7}
	//插入資料-------------------------------------
	arr1[1] = 9
	fmt.Println(arr1)

	//取得陣列長度-------------------------------
	fmt.Println(len(arr1))

	//字串陣列-----------------------------
	arr2 := [3]string{"Hello", "Mother", "Fucker"}
	fmt.Println(arr2)

	//練習: 巡迴陣列中的資料---------------------
	grades := [4]int{60, 90, 75, 10}
	sum := 0 //總和
	for index := 0; index < len(grades); index++ {
		sum += grades[index]
	}
	fmt.Println(sum)               //計算平均的總和
	fmt.Println(sum / len(grades)) //計算平均數

	//讓使用者自行輸入分數的做法-------------------
	fmt.Println("請逐一輸入資料")
	grades2 := [4]int{}

	for index := 0; index < len(grades2); index++ {
		fmt.Scanln(&grades2[index])
	}
	sum2 := 0 //總和
	for index := 0; index < len(grades); index++ {
		sum2 += grades2[index]
	}

	fmt.Println("總和", sum2)               //計算平均的總和
	fmt.Println("平均數", sum2/len(grades2)) //計算平均數

}
