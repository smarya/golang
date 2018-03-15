package main

import (
	"fmt"
)

func main() {
	var numsArr [5]float64
	numsArr[0] = 10
	numsArr[1] = 20.5
	numsArr[2] = 21.3
	numsArr[3] = 25.5
	numsArr[4] = 1.999

	//Print as specific element
	fmt.Println(numsArr[3])
	//print all elements
	fmt.Println(numsArr)

	//way2
	favNums2 := [5]float64{1, 2, 3, 4, 5}
	fmt.Println("both index and values")
	for index, value := range favNums2 {
		fmt.Println(index, value)
	}

	//way2
	favNums3 := [5]float64{1, 2, 3, 4, 5}
	fmt.Println("Only values ")
	for _, value := range favNums3 {
		fmt.Println(value)
	}
	//array indexs start from 0
	//slices
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//take leement sin slice one from 3rd pos to 9 pos and put them in slice2
	slice2 := slice1[3:9]
	fmt.Println("slice2[2]", slice2[2])
}
