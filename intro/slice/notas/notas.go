package main

import (
	"fmt"
	"strconv"
)

func main() {

	//var notas []float64
	notas := []float64{}

	for {
		var inputTxt string
		fmt.Printf("Entre com a sua nota:")
		fmt.Scanf("%s", &inputTxt)

		if inputTxt == "" {
			break
		}

		nota, err := strconv.ParseFloat(inputTxt, 64)
		if err != nil {
			fmt.Println("Erro ao converter a nota:", err)
			continue
		}

		notas = append(notas, nota)
	}
	fmt.Println(notas)
}
