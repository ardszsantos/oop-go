package main

import (
	"fmt"
	"oop-go/contas"
)

func main() {
	contaDaSilvia := contas.ContaCorrente{Titular: "Silvia", Saldo: 300}
	contaDoGustavo := contas.ContaCorrente{Titular: "Gustavo", Saldo: 100}
	fmt.Println(contaDaSilvia, contaDoGustavo)

	status := contaDaSilvia.Transferir(200, &contaDoGustavo)

	fmt.Println("A transferencia foi concluida?", status)

	fmt.Println(contaDaSilvia, contaDoGustavo)

}
