package main

import (
	"fmt"
	"github.com/splorg/object-oriented-golang/contas"
	// "github.com/splorg/object-oriented-golang/clientes"
)

type verificarConta interface {
	Sacar(valor float64) string
}

func PagarBoleto(conta verificarConta, valor float64) {
	conta.Sacar(valor)
}

func main() {
	contaDoGabriel := contas.ContaCorrente{}

	fmt.Println(contaDoGabriel)
	
	contaDoGabriel.Depositar(200)

	fmt.Println(contaDoGabriel.ObterSaldo())

	PagarBoleto(&contaDoGabriel, 50)

	fmt.Println(contaDoGabriel.ObterSaldo())
}