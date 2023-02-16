package main

import "fmt"

func main() {
	byteString := []byte{70, 117, 110, 99, 105, 111, 110, 97, 32, 109, 101, 115, 109, 111}

	fmt.Println(string(byteString)) //transformando uma string em bite


	stringByte := []byte("Funciona mesmo") //transformando um byte em string

	fmt.Println(stringByte)

}
