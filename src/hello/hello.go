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

}
