//Crie uma struct chamada Circulo com o campo "raio". Escreva uma função que receba um Circulo como parâmetro e calcule a área do círculo (área = pi * raio * raio). Dica: use a constante math.Pi para representar o número pi..

package main

import (
	"fmt"
	"math"
)

type circulo struct {
	raio float64
}

func area(c circulo) float64 {
	return c.raio * c.raio * math.Pi
}

func main() {
	var cir circulo
	fmt.Println("Informe o raio da circunferencia: ")
	fmt.Scan(&cir.raio)
	a := area(cir)
	fmt.Printf("Area é %.f ", a)
}
