package main

import "fmt"

//1. Fa�a um sistema que leia a idade de uma pessoa em dias, meses, e anos e mostre-a expressa apenas em dias.

func main() {
	var dia, mes, ano, totalDia int = 0, 0, 0, 0

	fmt.Println("Escreva o total de dias de anos")
	fmt.Scan(&ano)
	fmt.Println("Escreva o total de dias do mes")
	fmt.Scan(&mes)
	fmt.Println("Escreva o total de dias")
	fmt.Scan(&dia)

	ano = ano * 365 //convertendo os anos em dias
	mes = mes * 30  //convertendo meses em dias
	totalDia = ano + mes + dia

	fmt.Sprintf("Fazem %d dias que voc� nasceu!", totalDia)

}
