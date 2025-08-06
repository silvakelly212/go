package contas

import (
	"banco/clientes"
)

type ContaPoupanca struct {
	Titular                              clientes.Titular
	NumeroAgencia, NumeroConta, Operacao int
	ObterSaldo                           float64
}

func (c *ContaPoupanca) Sacar(valorDoSaque float64) string {
	//boleando as condicoes precisam ser verdadeiras
	// o valor do saque precisar ser maior que 0 && o valor do saque precisa ser menor ou igual ao saldo
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.ObterSaldo

	if podeSacar {
		c.ObterSaldo -= valorDoSaque
		return "Saque realiado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

// multiplos retornos

func (c *ContaPoupanca) Depositar(valorDoDeposito float64) (string, float64) {
	podeDepositar := valorDoDeposito >= 0

	if podeDepositar {
		c.ObterSaldo += valorDoDeposito
		return "Deposito realizado com Sucesso!", c.ObterSaldo
	} else {
		return "Deposito nao efetuado", c.ObterSaldo
	}

}
