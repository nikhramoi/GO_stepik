package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	randomNum := rand.IntN(500)
	fmt.Println(randomNum)
	max := 50
	min := 10
	randomNum2 := rand.IntN(max-min) + min
	fmt.Println(randomNum2)
}
