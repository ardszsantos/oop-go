package main

import (
	"fmt"
	"oop-go/clientes"
	"oop-go/contas"
)

func main() {
	contaDoBruno := contas.ContaCorrente{Titular: clientes.Titular{
		Nome:      "Bruno",
		CPF:       "431.525.324.53",
		Profissao: "Joalheiro",
	}, NumeroAgencia: 123, NumeroConta: 123456, Saldo: 100}

	fmt.Println(contaDoBruno)
}
