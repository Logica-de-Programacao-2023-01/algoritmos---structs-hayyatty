package main

import (
	"fmt"
	"time"
)

type Playlist struct {
	nome    string
	musicas []musica
}
type musica struct {
	titulo  string
	artista string
	duracao time.Duration
}

func imprimir(playlist Playlist) {
	fmt.Println("Playlist :", playlist.nome)
	total := 0 * time.Minute
	for _, musica := range playlist.musicas {
		fmt.Println("Título :", musica.titulo)
		fmt.Println("Artista :", musica.artista)
		fmt.Println("Duração :", musica.duracao)
		total += musica.duracao
	}
	fmt.Println("O tempo total da playlist é de :", total, "minutos")
}

func main() {
	playlist := Playlist{
		nome: "Minhas Músicas",
		musicas: []musica{
			{titulo: "Musica 1", artista: "Artista 1", duracao: 3 * time.Minute},
			{titulo: "Musica 2", artista: "Artista 2", duracao: 4 * time.Minute},
			{titulo: "Musica 3", artista: "Artista 3", duracao: 5 * time.Minute},
		},
	}
	imprimir(playlist)
}
