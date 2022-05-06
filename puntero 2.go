package main

import "fmt"

func main() {
	juan := 2
	maria := &juan
	casa := &maria

	if *maria == juan {
		fmt.Println(&juan)

	}
	if **casa == *maria {
		fmt.Println(&maria)
	}
}
