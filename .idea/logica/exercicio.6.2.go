package main

import "fmt"

//- Crie uma função que receba um parâmetro variádico do tipo int e retorne a soma de todos os ints recebidos;
//- Passe um valor do tipo slice of int como argumento para a função;

//- Crie outra função, esta deve receber um valor do tipo slice of int e retornar a soma de todos os elementos da slice;
//- Passe um valor do tipo slice of int como argumento para a função.

func soma ( numeros ... int) int{
valoresSomados := 0
		for _, n := range numeros {

			valoresSomados += n
		}

	return valoresSomados
}

func soma2 (numeros [] int) int{
	valoresSomados := 0
	for _, n := range numeros {

		valoresSomados += n
	}

	return valoresSomados
}
func main()  {
	//slice := []int{10, 11, 12, 13, 14, 15, 16, 17}

	numeros := []int {1,34,55,22,44}

	visualizar := soma(5,5,6,6,6)
	visualizar2 := soma2(numeros)


	fmt.Println(visualizar)
	fmt.Println(visualizar2)
}