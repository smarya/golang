package main

//code online https://play.golang.org/p/nfM-c0PWDMp
import "fmt"

func main() {
	i := 1
	//for loops
	for i <= 10 {
		fmt.Print(i, " ")
		i++
	}
	fmt.Println()
	for j := 0; j < 5; j++ {
		fmt.Print(j, " ")
	}
	fmt.Print("\n")

	yourAge := 18
	// If loops
	if yourAge >= 16 {
		println("you can drive")
	} else if yourAge >= 18 {
		println("You can vote")
	} else {
		print("you can't drive\n")
	}
	//Switch statement
	switch yourAge {
	case 16:
		fmt.Println("you can drive")
	case 18:
		fmt.Println("You can vote")
	default:
		fmt.Println("Go have fun")
	}

}
