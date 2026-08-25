// mapas são conjuntos de pares chave-valor com tipos especificos nas chaves e nos valores
// a chave é unica e o valor pode ser alterado

package main

import "fmt"

func main() {

	idades := make(map[string]int)
	// map é um tipo de dado que representa um mapa, onde a chave é do tipo string e o valor é do tipo int
	// make é uma função que cria um mapa vazio

	idades["alice"] = 25
	idades["bob"] = 30
	idades["carol"] = 28

	fmt.Println(idades)

	alturas := map[string]float64{}
	alturas["alice"] = 1.65
	alturas["bob"] = 1.80
	alturas["carol"] = 1.75

	fmt.Println(alturas)

	// atribuindo o valor da chave "carol" a uma variável
	alturaCarol := alturas["carol"]
	fmt.Println("Altura da Carol:", alturaCarol)

	// se eu tentar acessar uma chave que não existe, o valor retornado será o valor zero do tipo do valor
	// ele não dará erro, mas retornará o valor zero do tipo do valor
	alturaDavid := alturas["david"]
	fmt.Println("Altura do David:", alturaDavid) // valor zero do tipo float64 é 0.0

	// para verificar se uma chave existe no mapa, podemos usar a sintaxe de atribuição com vírgula
	alturaExistente, ok := alturas["david"]
	if ok {
		fmt.Println("Altura do David:", alturaExistente)
	} else {
		fmt.Println("David não está no mapa de alturas")
	}
	// ele não quebra o codigo, ele valida se a chave existe e retorna um booleano indicando se a chave existe ou não

	// outra forma de declarar e checar se uma chave existe no mapa é usando a sintaxe de atribuição com vírgula dentro de um if
	if valor, ok := alturas["bob"]; ok {
		fmt.Println("Altura do Bob:", valor)
	} else {
		fmt.Println("Bob não está no mapa de alturas")
	}
	// "valor" nao pode ser utilizado fora do escopo do if, pois ele é declarado dentro do if/ uma variavel temporária
	// muito comum no go para validar funcao e acesso a algum tipo de recurso
}
