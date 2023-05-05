// Crie uma struct chamada Pessoa com os campos "nome", "idade" e "endereço". O campo "endereço" deve ser u
// ma outra
// struct com os campos "rua", "número", "cidade" e "estado". Escreva uma função que receba uma Pessoa
// como parâmetro e imprima seu endereço completo.
package main

import "fmt"

type Endereco struct {
	rua    int
	cidade string
	estado string
}

type Pessoa struct {
	nome     string
	endereco Endereco
}

func main() {
	var pess Pessoa
	var end Endereco
	fmt.Println("Informe seu nome:")
	fmt.Scan(&pess.nome)
	fmt.Println(pess.nome, ",informe seu endereço completo")
	fmt.Print("Rua: ")
	fmt.Scan(&end.rua)
	fmt.Print("Cidade: ")
	fmt.Scan(&end.cidade)
	fmt.Print("Estado: ")
	fmt.Scan(&end.estado)
	p := Pessoa{nome: pess.nome,
		endereco: Endereco{rua: end.rua,
			cidade: end.cidade,
			estado: end.estado}}
	fmt.Printf("Nome: %s, Rua: %d, Cidade: %s, Estado: %s", p.nome, p.endereco.rua, p.endereco.cidade, p.endereco.estado)
}
