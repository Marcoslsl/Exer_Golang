// CALLBACK
// passe uma funcao como argumento para outra funcao.

package main

import "fmt"

func func01() {
	fmt.Println("ola")
}

func func2(funcao func()) {
	fmt.Println("prestencao:")
	funcao()

}

func main() {
	func2(func01)
}
