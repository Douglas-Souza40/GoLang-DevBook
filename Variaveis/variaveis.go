package main

import "fmt"

func main() {
	var variavel string = "Variavel1"
	variavel2 := "Variavel2"

	var (
		variavel3 string = "variavel3"
		variavel4 string = "variavel4"
	)

	fmt.Println(variavel, variavel2, variavel3, variavel4)

	variavel5, variavel6 := "variavel5", "variavel6"

	fmt.Println(variavel5, variavel6)

	const constante1 string = "Constante 1"

	fmt.Println(constante1)

	//não é possivel atribuir novo valor a Constante
	//	constante1 = "Constante"

	variavel5, variavel6 = variavel6, variavel5

	fmt.Println(variavel5, variavel6)
}
