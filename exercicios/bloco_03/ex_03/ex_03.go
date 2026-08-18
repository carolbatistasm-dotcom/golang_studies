// faça um programa que verifique se a pessoa pertence a familia "simões"

package main

import (
	"fmt"
	"strings"
)

func main() {
	var sobrenome string
	fmt.Println("Digite seu sobrenome:")
	fmt.Scanf("%s", &sobrenome)

	if strings.Contains(sobrenome, "Simões") || strings.Contains(sobrenome, "simões") {
		fmt.Println("Você deve ser meu/minha parente.")
	} else {
		fmt.Println("Você não pertence à família Simões.")
	}
}

// nao soluciona, porque o Scanf não lê o sobrenome completo, apenas a primeira palavra. Para resolver isso, podemos usar o bufio.NewReader para ler a linha inteira. Aqui está a versão corrigida:

//package main

//import (
//	"bufio"
//	"fmt"
//	"os"
//	"strings"
//)

//func main() {
//	reader := bufio.NewReader(os.Stdin)
//	fmt.Println("Digite seu sobrenome:")
//	sobrenome, _ := reader.ReadString('\n')
//	sobrenome = strings.TrimSpace(sobrenome)
//
//	if strings.Contains(sobrenome, "Simões") || strings.Contains(sobrenome, "simões") {
//		fmt.Println("Você deve ser meu/minha parente.")
//	} else {
//		fmt.Println("Você não pertence à família Simões.")
//	}
//}
