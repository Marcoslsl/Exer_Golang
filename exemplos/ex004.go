// Criar função que recebe um param variádico e retorne a soma de todos os ints
// Criar uma funcao que recebe um slice of ints

package main

import "fmt"

func func_qualquer(x ...int) int {
	soma := 0
	for _, v := range x {
		soma += v
	}
	return soma
}

func func_qualquer_2(x []int) int {
	soma := 0
	for _, v := range x {
		soma += v
	}
	return soma
}

func main() {
	var_ := []int{1, 2, 3}
	fmt.Println(func_qualquer(var_...))
	fmt.Println(func_qualquer_2(var_))
}
