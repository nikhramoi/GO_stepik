package main

import "fmt"

/* const (
	red = iota
	green
	blue
) */

/* const (
	_ = iota // ignore first value by assigning to blank identifier
	monday
	tuesday
	wednesday
	thursday
	friday
	saturday
	sunday
)

func main() {
	fmt.Println(monday, tuesday, wednesday, thursday, friday, saturday, sunday)
} */

const (
	a = iota
	b
	c = 5
	d
	e = iota
	f = iota - 2
	g = iota - 2
)

func main() {
	fmt.Println(a, b, c, d, e, f, g)
}
