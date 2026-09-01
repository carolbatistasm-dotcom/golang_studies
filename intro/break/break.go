package main

import "fmt"

// um exemplo com break, que é usado para sair de um loop
func main() {
	nome := "Carolina"
	for i, v := range nome {
		if v == 'a' {
			break // se o valor for 'a', sai do loop
		}
		fmt.Printf("Índice: %d, Valor: %c\n", i, v)
	}
}
