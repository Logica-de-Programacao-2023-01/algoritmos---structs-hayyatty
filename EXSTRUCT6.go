package main

import (
	"fmt"
	"time"
)

type funcionario struct {
	nome    string
	salario float64
	idade   int
}

func aumentar(selecionado funcionario) float64 {
	aumento := 0.0
	fmt.Println("Qual a porcentagem que deseja aumentar do salário de", selecionado.nome)
	fmt.Println("Não precisa usar o símbolo %")
	fmt.Scan(&aumento)
	aumentoreal := aumento / 100
	novosala := (selecionado.salario * aumentoreal) + selecionado.salario
	return novosala
}
func diminuir(selecionado funcionario) float64 {
	reducao := 0.0
	fmt.Println("Qual a porcentagem que deseja reduzir do salário de", selecionado.nome)
	fmt.Println("Não precisa usar o símbolo %")
	fmt.Scan(&reducao)
	reducaoreal := reducao / 100
	novosala := selecionado.salario - (selecionado.salario * reducaoreal)
	return novosala
}
func tempodeservico(selecionado funcionario) {
	tempo := 0
	fmt.Println("O funcionário esta com :", selecionado.idade, "anos")
	tempo = selecionado.idade - 18
	fmt.Println("Então o seu tempo de serviço é de aproximadamente", tempo, "anos")
}
func main() {
	i := 10
	sair := false
	selecionado := funcionario{
		nome:    "Carlos",
		salario: 1000,
		idade:   40,
	}
	for !sair {
		fmt.Println("")
		fmt.Println("Você acessou o perfil de funcionário do:", selecionado.nome)
		fmt.Println("O que você deseja fazer, ou verificar ?")
		fmt.Println("[1] Aumentar o salário ")
		fmt.Println("[2] Diminuir o salário ")
		fmt.Println("[3] Verificar o tempo de serviço")
		fmt.Println("[0] Finalizar e mostrar os dados atualizados")
		fmt.Scan(&i)
		switch i {
		case 1:
			fmt.Println("")
			fmt.Println("Salário atual:", selecionado.salario)
			selecionado.salario = aumentar(selecionado)

		case 2:
			fmt.Println("")
			fmt.Println("Salário atual:", selecionado.salario)
			selecionado.salario = diminuir(selecionado)

		case 3:
			fmt.Println("")
			tempodeservico(selecionado)
		case 0:
			fmt.Println("")
			fmt.Println("Os Dados finais são :")
			fmt.Println(selecionado.nome)
			fmt.Println(selecionado.salario)
			fmt.Println(selecionado.idade)
			sair = true
		}
		fmt.Println("Atualizando dados")
		for j := 0; j < 3; j++ {
			time.Sleep(1 * time.Second)
			fmt.Print(".")
		}
		fmt.Println("")
	}

}
