package main
//transforma em json
import (
	"encoding/json"
	"fmt"
)

type jsontogo []struct {
	Name   string   `json:"Name"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
}

func main() {
	s := `[{"Name":"James","Last":"Bond","Age":32},{"Name":"Miss","Last":"Moneypenny","Age":27},{"Name":"M","Last":"Hmmmm","Age":54}]`
	fmt.Println(s)

	var resultado jsontogo

	err := json.Unmarshal([]byte(s), &resultado)
	if err != nil {
		fmt.Println("Nao foi possivel executar unmarshal", err)
	}

	fmt.Println(resultado)
	//fmt.Println(resultado[1])
	//fmt.Println(resultado[1].Last)

}
