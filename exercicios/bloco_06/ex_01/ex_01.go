// faça o programa de uma sorveteria, onde o usuario pode escolher:
// o tipo do sorvete: casquinha (R$1,00), copo (R$2,00) ou cestinha (R$4,00)
// sabor do sorvete: chocolate, baunilha ou morango
// cobertura: granulado (R$0,50), calda de chocolate (R$1,00) ou calda de morango (R$1,00)
// apresente o valor a ser pago
// fazer a solucao com mapas, onde o tipo do sorvete, sabor e cobertura são as chaves e os valores são os preços

package main

import (
	"fmt"
)

func main() {

	tiposSorvete := map[string]float64{
		"casquinha": 1.00,
		"copo":      2.00,
		"cestinha":  4.00,
	}

	saboresSorvete := map[string]float64{
		"chocolate": 0.00,
		"baunilha":  0.00,
		"morango":   0.00,
	}

	coberturasSorvete := map[string]float64{
		"granulado":          0.50,
		"calda de chocolate": 1.00,
		"calda de morango":   1.00,
	}

	var tipo, sabor, cobertura string

	fmt.Println("Escolha o tipo do sorvete (casquinha, copo, cestinha):")
	fmt.Scanln(&tipo)

	fmt.Println("Escolha o sabor do sorvete (chocolate, baunilha, morango):")
	fmt.Scanln(&sabor)

	fmt.Println("Escolha a cobertura do sorvete (granulado, calda de chocolate, calda de morango):")
	fmt.Scanln(&cobertura)

	// eu poderia criar um mapa de mapa, onde a chave é uma string e o valor é outro mapa com string e float64, mas nesse caso não é necessario, pois cada mapa ja tem o seu proprio tipo de chave e valor
	// exemplo:
	//items := map[string]map[string]float64{
	//	"tipos":      tiposSorvete,
	//		"sabores":    saboresSorvete,
	//		"coberturas": coberturasSorvete,
	//	}

	// para validar a entrada do usário eu poderia usar:
	precoTipo, ok := tiposSorvete[tipo]
	if !ok {
		fmt.Println("Tipo de sorvete inválido")
		return
	}

	precoSabor, ok := saboresSorvete[sabor]
	if !ok {
		fmt.Println("Sabor de sorvete inválido")
		return
	}

	precoCobertura, ok := coberturasSorvete[cobertura]
	if !ok {
		fmt.Println("Cobertura de sorvete inválida")
		return
	}

	total := precoTipo + precoSabor + precoCobertura
	fmt.Printf("O valor a ser pago é: R$%.2f\n", total)

}
