package main

import "fmt"

func soma(x ...int) int {
	soma := 0
	for _, value := range x {
		soma += value
	}
	return soma
}

func somaImpar(f func(x ...int) int, y ...int) int {
	var slice []int
	for _, value := range y {
		if value%2 != 0 {
			slice = append(slice, value)
		}
	}
	total := f(slice...)
	return total
}

func main() {

	t := somaImpar(soma, []int{1, 3, 4, 2, 5}...)
	fmt.Println(t)
}
