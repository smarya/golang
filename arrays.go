package main

//code at https://play.golang.org/p/7IbU4dEWxs1
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
	//if we dont specify the first index the 0 will be considered the first index
	fmt.Println(slice1[:5])
	//same case if we dont give anything in back ok : . This is like 2 to last element
	fmt.Println(slice1[2:])
	// create an array of size 10
	slice3 := make([]int, 10)
	fmt.Println(slice3)
	//make has type as first and number as second parameter
	slice4 := make([]int, 5)
	copy(slice4, slice2)
	fmt.Println(slice4)
	slice4 = append(slice4, 0, -1)
	fmt.Println(slice4)
	fmt.Println(slice4[len(slice4)-1])

}
