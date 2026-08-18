package main

import "fmt"

func main() {

	// posso criar string com " "  ou `` e posso escrever com multiplas linhas com crase
	fmt.Println("teste")
	fmt.Printf(`teste
teste
teste`)

	fmt.Println("carolina tem", len("carolina"), "letras")
	fmt.Println("Carolina"[0])         // isso é 1 byte (int8)
	fmt.Println(string("Carolina"[0])) // preciso falar que é uma string
}
