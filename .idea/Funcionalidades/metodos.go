package main

import "fmt"

//os metodos estao associado a alguma coisa, por exemplo um struct ou uma interface
//usamos o reciver
//ele adiciona funcionalidades a um tipo especifico
//exercicio
//- Crie um tipo struct "pessoa" que contenha os campos:
//- nome
//- sobrenome
//- idade
//- Crie um método para "pessoa" que demonstre o nome completo e a idade;
//- Crie um valor de tipo "pessoa";
//- Utilize o método criado para demonstrar esse valor.

type Pessoa struct {
	name string
	sobrenome string
	idade int
}

func (p Pessoa) ImprimirPessoa () {
	fmt.Sprintf("O nome completo dessa pessoa é: %s %s, essa pessoa tem %d anos" ,p.name, p.sobrenome, p.idade)

}

func main()  {
	novaPessoa := Pessoa {
		name: "Pessoa 1",
		sobrenome: "Sobrenome pessoa",
		idade: 25,
	}

	fmt.Println(novaPessoa)
}