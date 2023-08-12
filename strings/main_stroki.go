package main

import "fmt"

func main() {
	var (
		x1    = "\tHello1\n"
		x2    = 'H'
		x3    = "HHH"
		x4    = "Куку"
		x5    = x4[:6]
		x6    = x4[2:]
		word  = "\tHello"
		hello = word + " world"
	)

	fmt.Println(x1, x2, x3, len(x1), x3[0], len(x4), x4[0], x4[0:8], x5, x6, hello)

}
