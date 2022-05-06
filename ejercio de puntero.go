package main

import "fmt"

func main() {
	var num *int
	if num == nil {
		fmt.Println("num", num, "es nil")
	} else {
		fmt.Println(num, "no es nil")
	}
}
