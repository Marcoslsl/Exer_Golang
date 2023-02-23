// CLOSURE

package main

import "fmt"

func i() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func main() {
	a := i()
	b := i()

	for i := 0; i < 3; i++ {
		fmt.Println(a())
		fmt.Println(b())
	}

}
