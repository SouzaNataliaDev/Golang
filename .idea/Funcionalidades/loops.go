package main

import "fmt"

func main() {
	//i := 0
	//
	//for i < 10 {
	//	i++
	//	fmt.Println("Incrementando i")
	//	time.Sleep(time.Second)
	//}

	//fmt.Println(i)

	// for j := 0; j < 10; j += 5 {
	// 	fmt.Println("Incrementando j", j)
	// 	time.Sleep(time.Second)
	// }

	// nomes := [3]string{"João", "Davi", "Lucas"}

	// for _, nome := range nomes {
	// 	fmt.Println(nome)
	// }

	for _, letra := range "PALAVRA" {
		if string(letra) == "A" {
			fmt.Println("letra a")
		}
		fmt.Println(string(letra))
	}

	// usuario := map[string]string{
	// 	"nome":      "Leonardo",
	// 	"sobrenome": "Silva",
	// }

	// for chave, valor := range usuario {
	// 	fmt.Println(chave, valor)
	// }

	//for {
	//	fmt.Println("Executando infinitamente")
	//	time.Sleep(time.Second)
	//}
}
