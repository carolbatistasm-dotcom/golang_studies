package main

import "fmt"

func main() {
	for true {
		fmt.Println("Loop infinito")
	}
}

// se eu não colocar "true" no for, ele vai ser um loop infinito também.
// posso criar a variavel fora da estrutura do for, mas sempre tenho que colocar a condição no for, senão ele vai ser um loop infinito.
