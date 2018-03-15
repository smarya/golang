package main

//code at  https://play.golang.org/p/TtljtOXV3HL
import (
	"fmt"
)

func main() {
	presAge := make(map[string]int)
	presAge["Roosevelt"] = 42
	fmt.Println(presAge["Roosevelt"])
	presAge["Jhon F. Kennedy"] = 43
	fmt.Println(len(presAge))
	delete(presAge, "Jhon F. Kennedy")
	fmt.Println(len(presAge))
}
