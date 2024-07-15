package main

import (
	"fmt"
)

func main() {

	var name string = "Rodrigo"

	fmt.Println("Olá", name)
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Analisar logs")
	fmt.Println("3 - Sair \n")

	var opcao int
	fmt.Scanf("%d", &opcao)
	print("A opção foi: ", opcao)
}
