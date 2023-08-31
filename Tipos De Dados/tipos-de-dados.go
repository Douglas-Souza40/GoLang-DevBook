package main

import "fmt"

func main() {

	// int8, int16, int32, int64 -> varia a quantidade de bytes
	var num1 int64 = 10000000000
	fmt.Println(num1)

	//INT32 == RUNE
	///uint deve receber sem sinal
	var num2 uint32 = 1000
	fmt.Println(num2)

	//float32, float64 -> numeros flutuantes
	var numReal1 float32 = 123.45
	var numReal2 float64 = 12345.67
	fmt.Println(numReal1, numReal2)

	// String
	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	//Char pega o valor da tabela ASC se vc colocar 1 caracter
	char := 'B'
	fmt.Println(char)

	// booleano  = true/false
	var booleano bool
	fmt.Println(booleano)
}
