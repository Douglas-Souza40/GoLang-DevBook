package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo do arquivo main !!!")

	auxiliar.Escrever()

	//colocar um emailvalido para testar
	err := checkmail.ValidateFormat("teste")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Não teve erro no validador de email")
	}

}
