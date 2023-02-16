package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("isso vai aparecer depois do depois")
	defer fmt.Println("isso vai aparecer depois")
	fmt.Println("isso vai aparecer antes")

	// e se tiver defer duas vezes ?
	// o que entra primeiro é o ultimo a ser executado
}
