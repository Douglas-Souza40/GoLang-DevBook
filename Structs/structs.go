package main

import "fmt"

type usuario struct {
	nome     string
	idade    int8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     int
}

func main() {

	var u usuario
	u.nome = "Dolgras"
	u.idade = 42

	fmt.Println(u)

	end := endereco{"rua A ", 44}

	u2 := usuario{"katiuca", 48, end}
	fmt.Println(u2)

	u3 := usuario{nome: "Dreisson"}
	fmt.Println(u3)

}
