package main

import (
	"fmt"
)

func main() {
	arr := [3]int{1, 2, 3}
	//var arr2 := [...]{"qwe","sad","asdasd"}
	var arr3 = [...]string{"qwe", "sad", "asdasd"}
	arr5 := [...]string{"foo", "bar"}
	arr3[0] = "qwe"
	arr3[1] = "sad"
	arr3[2] = "asdasdsadsadsadasdasd"
	fmt.Println(arr, arr3, arr5)

	var slice = make([]string, 5, 5)
	fmt.Println(slice, len(slice), cap(slice))
	slice[0] = "qwe"
	slice[1] = "sad"
	slice[2] = "asdasdsadsadsadasdasd"
	slice[3] = "asdasdsadsadsadasdasd"
	slice[4] = "asdasdsadsadsadasdasd"
	fmt.Println(slice, len(slice), cap(slice))
	slice = append(slice, "foo")
	slice = append(slice, "bar")
	fmt.Println(slice, len(slice), cap(slice))
	slice = append(slice, "barbasisca")
	fmt.Println(slice, len(slice), cap(slice))
}
