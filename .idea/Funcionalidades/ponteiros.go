package main

import "fmt"

func changeValue(a *int) {

	ponteiro := a
	*ponteiro = 50
	fmt.Println(ponteiro)

}

//func changeValue2(a int) int {
//	ponteiro = &a
//	*ponteiro = 50
//	return a
//
//}

func main() {
	var variavel1 int = 10
	var variavel2 int = variavel1

	fmt.Println(variavel1, variavel2)

	variavel2++
	fmt.Println(variavel1, variavel2)

	// PONTEIRO É UMA REFERÊNCIA DE MEMÓRIA
	var variavel3 int
	var ponteiro *int

	variavel3 = 100
	ponteiro = &variavel3

	fmt.Println(variavel3, ponteiro)

	variavel3 = 150
	fmt.Println(variavel3, *ponteiro)
	var a = 100
	changeValue(&a)

	fmt.Println(a)

	//Para demonstrar o endereço de memoria de um ponteiro adiiona o & na frente
}
