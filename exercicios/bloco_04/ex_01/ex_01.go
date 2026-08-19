//Faça um programa que conte quantas vezes a letra “a” aparece em uma palavra

package main

import "fmt"

func main() {
	var palavra string
	fmt.Print("Digite uma palavra: ")
	fmt.Scanln(&palavra)

	contador := 0
	for _, letra := range palavra {
		if letra == 'a' || letra == 'A' { // Considerando tanto 'a' minúsculo quanto 'A' maiúsculo
			contador++
		}
	}

	fmt.Printf("A letra 'a' aparece %d vezes na palavra '%s'.\n", contador, palavra)
}

// eu poderia usar uma comando como palavra = strings.ToLower(palavra) para
// transformar a palavra em minúscula e assim contar apenas uma vez
