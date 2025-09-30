package main

import "fmt"

// Объявление нескольких констант
const (
	A = true
	B = "Строка"
	C = 3.141
)

func main() {
	var age int = 40
	age2 := 50
	age2 += 1
	fmt.Println(age, age2)
	fmt.Printf("Type: %T\n", age)
	fmt.Printf("Type: %T\n", age2)
	const PI = 3.14
	fmt.Println(PI)
	fmt.Printf("%T", PI)
}
