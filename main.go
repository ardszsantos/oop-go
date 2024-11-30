package main

import (
	"fmt"
	"oop-go/clientes"
	"oop-go/contas"
)

func main() {
	clienteBruno := clientes.Titular{"Bruno", "123.123.123.12", "Desenvolvedor go"}
	contaDoBruno := contas.ContaCorrente{clienteBruno, 123, 12314, 100}

	fmt.Println(contaDoBruno)
}
