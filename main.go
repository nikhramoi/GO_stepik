package main

import (
	_ "fmt"
	//ранее выполнена go mod init main (main - название программы)
	"main/nettest" //тут создан пакет в папке nettest она соответствует директиве в файлах package nettest, поскольку в
	// go.mod указан module main то обращаемся к пакету "main/nettest"
	//. "main/nettest" - можно так но в сложных проектах так лучше не делать, окращает обращение к импортируемым методам и функциям
)

func main() {
	nettest.Say_hello()
}
