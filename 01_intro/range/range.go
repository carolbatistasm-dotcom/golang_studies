package main

import "fmt"

func main() {
	nome := "Carolina"
	for i, v := range nome {
		fmt.Printf("Índice: %d, Valor: %c\n", i, v)
	}
	// i me dá o índice do elemento
	// v me dá o valor do elemento
	// O range percorre a string e me dá o índice e o valor de cada elemento da string
	// Se eu não quiser o valor, posso usar apenas o índice.
	// %c é usado para imprimir o valor do elemento como caractere

	for i := range nome {
		fmt.Println(i, nome[i], string(nome[i]))
	}
	// já aqui i me dá o indice do elemento
	// nome[i] me dá o valor em bites do elemento
	// string(nome[i]) me dá o valor do elemento em string
}
