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

	// o valor de x vai se acrescentando pois o x é definido na primeira func
	// ou seja, na definição em "a := i()"
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())

	// como nao estou atribuindo a nenhuma variavel o x sempre vai se reiniciar
	fmt.Println(i()())
	fmt.Println(i()())
}
