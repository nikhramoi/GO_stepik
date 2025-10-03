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

	// Карты map аналог словарей и питухона
	map1 := map[string]int{
		"Вася":        25,
		"Пупкин Иван": 45,
	}
	map1["Петров Петр"] = 30
	delete(map1, "Вася")
	fmt.Println(map1["Михаил"])
	map1["Михаил"]++
	fmt.Println(map1)

	//Структуры
	type Point struct {
		x int
		y int
	}
	p := Point{
		x: 5,
		y: 19,
	}
	p.x = 25
	fmt.Println(p, p.x, p.y)
}
