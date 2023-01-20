package main

import (
	"fmt"
)

func plus(a int, b int) int {
	return a + b
}

func calculosMatematicos(n1, n2 int) (int, int, int) {
	soma := n1 + n2
	subtracao := n1 - n2
	divisao := n1 / n2
	return soma, subtracao, divisao
}

func main() {

	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	_, resultadoSubtracao, _ := calculosMatematicos(10, 15) //no go pode colocar mais que um retorno, como essa
	fmt.Println(resultadoSubtracao)                         // funcao tem 2 retornos o _ informa para ignorar o primeiro

}
