package main

import "fmt"
//ponteiro serve pra quando estiver usando dados e nao querer ficar copiando

//- Crie um struct "pessoa"
//- Crie uma função chamada mudeMe que tenha *pessoa como parâmetro. Essa função deve mudar um valor armazenado no endereço *pessoa.


type Person struct {
	name string
}

func mudeMe( person *Person){
	(*person).name = "Mudei de nome"
	fmt.Println(person.name)
}

func mudeMeSemPonteiro(person Person){
	person.name = "Mudei de nome"
	fmt.Println(person.name)
}

func main (){

	pessoa := Person{
		name: "Natalia",
	}
	pessoa2 := Person{
		name: "Natalia",
	}

	mudeMe(&pessoa)
	mudeMeSemPonteiro(pessoa2)

	fmt.Println(pessoa2.name)
	fmt.Println(pessoa.name)

}