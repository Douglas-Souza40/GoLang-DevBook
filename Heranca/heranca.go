package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
	altura    float32
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {

	p1 := pessoa{"Jão", "Pedro", 20, 178}
	fmt.Println(p1)

	estudante1 := estudante{p1, "TI", "USP"}
	fmt.Println(estudante1.nome)

}
