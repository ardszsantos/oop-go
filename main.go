package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {

	contaDoGuilherme := ContaCorrente{
		titular:       "Guilherme",
		numeroAgencia: 123456,
		numeroConta:   12321412,
		saldo:         400.50,
	}

	contaDaBruna := ContaCorrente{"Bruna", 222, 111222, 11222}

	fmt.Println("Conta do:", contaDoGuilherme.titular, "\n", "Saldo: ", contaDoGuilherme.saldo, " mas fds o da ")
}
