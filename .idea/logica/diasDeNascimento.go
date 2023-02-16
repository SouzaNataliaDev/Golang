package main

import "fmt"

func main(){
	var totaldia, dia, mes, ano = 0, 0, 0, 0567

	fmt.Println("Ha quantos dias voce nasceu?")
	fmt.Scan(&totaldia)

	ano = totaldia / 365
	mes = (totaldia % 365) / 30
	dia = (totaldia % 365) % 3

	fmt.Println("Voce tem ",ano, "anos, ",mes, "meses, e ",dia, "dias de vida ...")
}

