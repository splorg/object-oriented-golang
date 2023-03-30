package contas

import (
	"github.com/splorg/object-oriented-golang/clientes"
)

type ContaCorrente struct {
	Titular clientes.Titular
	NumeroAgencia, NumeroConta int
	saldo float64
}

func (c *ContaCorrente) Sacar(valor float64) string {
	podeSacar := valor >= 0 && valor <= c.saldo

	if podeSacar {
		c.saldo -= valor
		return "Saque realizado com sucesso."
	} else {
		return "Transação inválida."
	}
}

func (c *ContaCorrente) Depositar(valor float64) (string, float64) {
	if valor > 0 {
		c.saldo += valor
		return "Depósito realizado com sucesso.", c.saldo
	} else {
		return "Valor do depósito menor que zero.", c.saldo
	}
}

func (c *ContaCorrente) Transferir(valor float64, contaDestino *ContaCorrente) bool {
	if valor > 0 && valor <= c.saldo {
		c.saldo -= valor
		contaDestino.Depositar(valor)
		return true
	} else {
		return false
	}
}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}