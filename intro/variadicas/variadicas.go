package main

import "fmt"

// funcao variática é uma função que pode receber um número indefinido de parâmetros do mesmo tipo

func soma(values ...int) int { //... significa que a função pode receber um número indefinido de parâmetros do tipo int

	total := 0
	for _, v := range values {
		total += v
	}
	return total
}
func main() {

	var a, b, c int

	a = 10
	b = 20
	c = 30

	total := soma(a, b, c, 100, 200, 300) //passando os valores para a função soma
	fmt.Println(total)

	valor := []int{1, 2, 3, 4, 5}
	total = soma(valor...) //passando um slice para a função soma
	fmt.Println(total)
}
