package main

//bcrypt - metodo que criptografa a senha
import (
	bcrypt "golang.org/x/crypto/bcrypt"
	"fmt"
)

func main()  {
	senha := "natalia"
	senhaErrada := ""

	senhaCriptografada, err := bcrypt.GenerateFromPassword( [] byte (senha), 10)

	if err != nil{
		fmt.Print(err)
	}

	fmt.Println("senha encriptografada:", senhaCriptografada)

result :=	bcrypt.CompareHashAndPassword((senhaCriptografada),[]byte(senha))
result2 :=	bcrypt.CompareHashAndPassword((senhaCriptografada),[]byte(senhaErrada))

fmt.Println(result)
fmt.Println(result2)
}

