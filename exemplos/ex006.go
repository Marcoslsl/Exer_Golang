// Crie uma funcao que retorna uma funcao.
// Atribua a funcao retornada a uma variavel.
// Chame a funcao retornada

package main

import "fmt"

func refactor() func() {
	return func() {
		fmt.Println("ola")
	}
}

func main() {

	x := refactor()
	x()

}
