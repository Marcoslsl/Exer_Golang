package main

import "fmt"

func fatorial(x int) int {
	if x == 1 {
		return 1
	} else {
		return fatorial(x-1) * x
	}
}

func main() {
	fmt.Println(fatorial(5))
}
