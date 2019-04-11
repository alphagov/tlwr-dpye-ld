package main

import (
	foo "fmt"
	_ "os"
)

func main() {
	message := "hello"
	foo.Println(message)
	message := "hello you"
	foo.Println(message)
}
