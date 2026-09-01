// array é um conjunto de elementos com um número de elementos fixo

package main

import "fmt"

func main() {
	// aqui a variavel será um array de tamanho 10 com valor igual a 0

	var x [10]float64
	fmt.Println(x)

	// aqui a variavel será um array de tamanho 20 com valor igual a 0

	var y = [20]float64{0.32, 12}
	y[18] = 33
	y[19] = 44 // mudar os elementos com base no indice
	fmt.Println(y)
	fmt.Println("Primeiro elemento de y = y[0] = ", y[0])
	fmt.Println("Dois primeiros elementos de y = y[0:2] = ", y[0:2])           // indices 0 e 1
	fmt.Println("Dois últimos elementos de y = y[len(y)-2:] = ", y[len(y)-2:]) // indices 18 e 19

	//para somar os elementos de um array
	var total float64
	for i, v := range y {
		fmt.Println("Índice =", i, "| Vetor =", v)
		total += v
	}

	// aqui ele começaria com 0.32 depois 12, e os 8 números seguintes seriam 0

	t := [10]string{"Carol", "Carolina"}
	fmt.Println(t)

	// aqui ele preenche o array com as duas strings, e as 8 restantes ficam com 0
	// a sintaxe é diferente também
	fmt.Println("Valores do array t:", len(t))
}
