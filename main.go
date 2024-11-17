package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {

	var contaDoGuilherme ContaCorrente = ContaCorrente{
		titular: "Guilherme",
		saldo:   400.50,
	}

	fmt.Println("Conta do:", contaDoGuilherme.titular, "\n", "Saldo: ", contaDoGuilherme.saldo)
}
