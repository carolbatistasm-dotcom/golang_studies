package main

import "fmt"

func soma(a int, b int) int {
	res := a + b
	return res

}

func media(a int, b int) (res float64, error error) {
	total := soma(a, b)
	res = float64(total) / 2
	error = nil
	return
}

func main() {

	n1, n2 := 10, 20
	fmt.Println(n1, n2)

	total := soma(n1, n2)
	fmt.Println("Total:", total)

	m, erro := media(n1, n2)
	if erro != nil {
		fmt.Println("Erro:", erro)
		return
	}
	fmt.Println("Média:", m)
}
