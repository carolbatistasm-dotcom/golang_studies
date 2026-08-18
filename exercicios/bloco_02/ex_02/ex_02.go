package main

import "fmt"

// vai dar erro no main, pq o ex_01.go já usa main, então não pode ter outro main no mesmo pacote
// todos os arquivos no 01_intro tem erro no "main" pq o main é o ponto de entrada do programa, então só pode ter um main no mesmo pacote
// a solucao portanto é criar outro pacote, por exemplo "bloco02", e colocar o main dentro dele, assim não vai dar erro no main

func main() {

	fmt.Println("Bom dia! Qual é o seu nome?")

	var nome string
	fmt.Scanf("%s", &nome)

	fmt.Printf("É um prazer te conhecer, %s . \n", nome)
}
