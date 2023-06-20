package main

import (
	"fmt"
	"reflect"
)

func main() {
	var name string = "Crud"
	var age = 42
	version := 1.1
	fmt.Println("Ola sr.", name, "sua idade é", age)
	fmt.Println("Este programa esta na versão", version)
	// fmt.Println("Age:", age)

	fmt.Println("O tipo da variavel nome é:", reflect.TypeOf(name))

	fmt.Println("===============================")
	fmt.Println("| [1] Iniciar Monitoramento   |")
	fmt.Println("| [2] Exibir os Logs          |")
	fmt.Println("| [0] Sair do Programa        |")
	fmt.Println("===============================")
	var option int
	fmt.Scanf("%d", &option)
	// fmt.Scan(&option)
	fmt.Println("O endereço da minha variavel option é:", &option)
}
