package main

import "fmt"

func main() {
	// da para declara a variaveis com o tipo explicito e implicito

	var variavel1 string = "Variavel 1" //explicito

	var (
		variavel3 string = "Variavel 3"
		variavel4 string = "Variavel 4" //declarando mais que uma variavel juntas
	)

	fmt.Println(variavel3, variavel4)

	variavel2 := "Variavel 2" //implicito

	variavel5, variavel6 := "Variavel 5", "Variavel 6" //declarando mais que uma variavel de  modo implicito

	fmt.Println(variavel5, variavel6)

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	const constante1 string = "Constante 1 -> valor imutavel"

	fmt.Println(constante1)

	fmt.Printf("v is of type %T\n", variavel2) //imprimir o tipo da variavel

}
