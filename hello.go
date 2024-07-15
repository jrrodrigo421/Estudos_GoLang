package main

import (
	"fmt"
	"reflect"
)

func main() {

	var name string = "Rodrigo"
	var typeVar = reflect.TypeOf(name)
	println(">> typeVar: ", typeVar)
	println(">> typeVar: ", reflect.TypeOf(name))
	fmt.Println("Olá", name)
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Analisar logs")
	fmt.Println("3 - Sair \n")

	var opcao int
	fmt.Scanf("%d", &opcao)
	print("A opção foi: ", opcao, "\n")
	print("Endereço da memória ", &opcao)
}
