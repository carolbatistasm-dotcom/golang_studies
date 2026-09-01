package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d\n", i)
	}
}

// for é a única estrutura de repetição do Go. Ele pode ser usado como um while, for ou for-each.
// a ideia do for é variável; condição; incremento.
// A variável é inicializada, a condição é verificada e o incremento é feito no final de cada iteração.
