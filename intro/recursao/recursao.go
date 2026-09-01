package main

import "fmt"

// a recursão é uma técnica de programação onde uma função chama a si mesma para resolver um problema
// A função fib(n) calcula o n-ésimo número da sequência de Fibonacci, que é definida como:

func fib(n uint) int { // uint é um tipo de dado inteiro sem sinal, ou seja, não pode ser negativo. Ele é usado aqui para garantir que o valor de n seja sempre positivo
	if n <= 1 {
		return int(n) // se n for 0 ou 1, a função retorna n, pois os dois primeiros números da sequência de Fibonacci são 0 e 1
	} else {
		return fib(n-1) + fib(n-2) // caso contrário, a função retorna a soma dos dois números anteriores da sequência de Fibonacci
	}
}
func main() {

	fmt.Println("fib(0) =", fib(0))
	fmt.Println("fib(0) =", fib(1))
	fmt.Println("fib(0) =", fib(2))
	fmt.Println("fib(0) =", fib(3))
	fmt.Println("fib(0) =", fib(4))
	fmt.Println("fib(0) =", fib(5))

}
