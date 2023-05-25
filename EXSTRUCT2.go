package main

import (
	"fmt"
)

type pessoa struct {
	nome     string
	idade    int
	endereco endereco
}
type endereco struct {
	rua    string
	número int
	cidade string
}

func main() {
	s := pessoa{
		nome:  "Lucas",
		idade: 18,
		endereco: endereco{
			rua: "QI", número: 901, cidade: "Guara",
		},
	}
	var (
		age  int
		name string
	)
	fmt.Println("Qual o nome da pessoa que deseja procurar ?")
	fmt.Scan(&name)
	fmt.Println("Qual a idade da pessoa que deseja procurar ?")
	fmt.Scan(&age)
	if s.nome == name && s.idade == age {
		fmt.Println("A pessoa foi encontrada")
		fmt.Println("O endereço registrado é :")
		fmt.Println("Rua :", s.endereco.rua)
		fmt.Println("Número :", s.endereco.número)
		fmt.Println("Cidade :", s.endereco.cidade)

	} else {
		fmt.Println("Pessoa não encontrada no banco de dados")
		fmt.Println(name, age)
	}
}
