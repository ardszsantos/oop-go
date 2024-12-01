package contas

import (
	"fmt"
	"oop-go/clientes"
)

type ContaPoupanca struct {
	Titular                              clientes.Titular
	NumeroAgencia, NumeroConta, Operacao int
	saldo                                float64
}

func main() {

	fmt.Println("todes")
}
