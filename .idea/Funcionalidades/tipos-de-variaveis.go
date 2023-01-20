package main

import (
	"errors"
	"fmt"
)

func main() {
	var erro error = errors.New("Novo erro")
	fmt.Println(erro)

}
