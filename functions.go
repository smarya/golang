package main

//code at https://play.golang.org/p/08156gWMm_r
import (
	"fmt"
)

func main() {
	floatArr := []float64{1, 2, 3, 4, 5, 6, 7}
	sum := sum(floatArr)
	fmt.Println(sum)

	num1, num2 := next2Values(5)

	fmt.Println(num1, num2)

	num3 := substractThem(1, 2, 3, 4, 5, 6, 7)

	fmt.Println(num3)

}

func sum(numbers []float64) float64 {
	sum := 0.0

	for _, val := range numbers {
		sum += val
	}

	return sum
}

func next2Values(number int) (int, int) {

	return number + 1, number + 2
}

//variadic parameters
func substractThem(args ...int) int {
	finalVal := 0
	for _, val := range args {
		finalVal -= val
	}
	return finalVal
}
