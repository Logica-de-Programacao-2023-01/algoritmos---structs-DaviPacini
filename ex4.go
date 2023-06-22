//Crie uma struct chamada Playlist com os campos "nome" e "músicas". O campo "músicas" deve ser um slice de structs,
//cada uma representando uma música com os campos "título", "artista" e "duração". Escreva uma função que receba uma
//Playlist como parâmetro e imprima o nome da playlist, o tempo total da playlist e as informações de cada música.

package main

import "fmt"

type Musica struct {
	titulo  string
	artista string
	duracao int
}

type Playlist struct {
	nome    string
	musicas []Musica
}

func main() {
	m := Musica{titulo: "Lover", artista: "Taylor Swift", duracao: 3}
	p := Playlist{nome: "Taylor", musicas: []Musica{m}}
	fmt.Println(p)
}
