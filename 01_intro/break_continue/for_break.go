package main

import "fmt"

func main() {
	name := "Carolina Simões"

	for _, v := range name { // _ é usado para ignorar o índice, pois não precisamos dele aqui

		letra := string(v)
		if letra == "o" {
			break // se o valor for 'a', sai do loop
		}
		if letra == " " {
			continue // se o valor for espaço, pula para a próxima iteração
		}

		fmt.Println(letra)
	}
}
