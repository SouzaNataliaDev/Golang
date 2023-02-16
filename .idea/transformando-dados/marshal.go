package main
//analisa um json e transforma
//import (
//"encoding/json"
//"fmt"
//)
//
//type user struct {
//	Name string
//	Age   int
//}
//
//func main() {
//	u1 := user{
//		Name: "Nome usuario 1", //campos com numeros minusculos nao sao exportados
//		Age:   32,
//	}
//
//
//	u2 := user{
//		Name: "Nome usuario 2",
//		Age:   27,
//	}
//
//	u3 := user{
//		Name: "Nome usuario 3",
//		Age:   54,
//	}
//
//	users := []user{u1, u2, u3}
//
//	fmt.Println(users)
//
//	sb, err := json.Marshal(users)
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	fmt.Println(string(sb))
//
//}
