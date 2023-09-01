package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculos(n1, n2 int8) (int8, int8, int8, int8) {
	soma := n1 + n2
	subtracao := n2 - n1
	multiplicacao := n1 * n2
	divisao := n2 / n1
	return soma, subtracao, multiplicacao, divisao

}

func main() {

	result := somar(10, 13)
	fmt.Println(result)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	resultado := f("Função F")
	fmt.Println("este é o resultado ", resultado)

	resultadoSoma, resresultadoSub, resresultadoMult, resresultadoDivi := calculos(10, 30)
	fmt.Println(resultadoSoma, resresultadoSub, resresultadoMult, resresultadoDivi)

}
