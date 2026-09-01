package main

import "fmt"

//uint8,uint16,uint32,uint64 - representacao de numeros positivos e seus bits
//int8,int16,int32,int64
//float32 e float64 - decimais/float

func main() {
	fmt.Println("1+1=", 1+1)
	fmt.Println("1+2.5=", 1+2.5)
	fmt.Println("10/3=", 10/3.0) // pelo menos 1 precisa ser float para aparecer as casas decimais
	fmt.Println("5 % 3=", 5%3)
	fmt.Printf("tipo 100*0.75 = %T \n", 100*0.75) // print f imprime formatos
}
