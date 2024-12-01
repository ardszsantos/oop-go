package main

import (
	"fmt"
	"oop-go/contas"
)

func main() {
	contaDoDenis := contas.ContaPoupanca{}
	contaDoDenis.Depositar(100)
	contaDoDenis.Sacar(2)

	fmt.Println(contaDoDenis.ObterSaldo())
}
