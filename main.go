package main

import (
	"fmt"
	"oop-go/contas"
)

func main() {
	contaExemplo := contas.ContaCorrente{}
	contaExemplo.Depositar(2000)

	fmt.Println(contaExemplo.ObterSaldo())
}
