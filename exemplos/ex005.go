// Crie um struc Pessoa com campos "nome", "sobrenome" e "idade"
// Crie um method para pessoa que mostre o nome completo e idade
// Crie um valor de tipo Pessoa
// Utilize o metodo criado para demonstar esse valor
package main

import "fmt"

type Pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func (p Pessoa) show_info() {
	fmt.Println("My name's:", p.nome, p.sobrenome)
	fmt.Println("I'm:", p.idade)
}

func main() {
	pessoa1 := Pessoa{
		nome:      "Marcos",
		sobrenome: "Vincius",
		idade:     24,
	}
	pessoa1.show_info()
}
