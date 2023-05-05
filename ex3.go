//Crie uma struct chamada Triângulo com os campos "base" e "altura". Escreva uma função que receba um Triângulo como
//parâmetro e calcule a área do triângulo (área = base * altura / 2).

package main

import "fmt"

type Triangulo struct {
	base   int
	altura int
}

func area(t Triangulo) int {
	return t.base * t.altura / 2
}
func main() {
	var tri Triangulo
	fmt.Println("Informe a base e a altura do triângulo:")
	fmt.Scan(&tri.base)
	fmt.Scan(&tri.altura)
	a := area(tri)
	fmt.Printf("A área do Trinângulo é: %d", a)
}
