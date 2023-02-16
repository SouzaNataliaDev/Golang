package main

//Encoder quando queremos “escrever”
//Decoder quando queremos “ler”
import (
	"encoding/json"
	"fmt"
	"os"
)

type user struct {
	Name string
	Last string
	Age  int
}

func main() {
	u1 := user{
		Name: "James",
		Last: "Bond",
		Age:  32,
	}

	u2 := user{
		Name: "Miss",
		Last: "Moneypenny",
		Age:  27,
	}

	u3 := user{
		Name: "M",
		Last: "Hmmmm",
		Age:  54,
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)


	err := json.NewEncoder(os.Stdout).Encode(users) //usando

	if err != nil {
		fmt.Println("Deu ruim aqui, ó só:", err)
	}

}
