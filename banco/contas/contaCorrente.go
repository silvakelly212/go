package contas

import (
	"banco/clientes"
)

// criando uma aplicação para gerenciar conta corrente
// Estrutura de conta corrente
// type e um novo tipo criado dentro de uma estrutura chamada - struct no caso ContaCorrente

type ContaCorrente struct {
	Titular            clientes.Titular
	NumeroAgencia      int
	NumeroConta        int
	SaldoContaCorrente float64
}

// Quando chamamos a função com (c *ContaCorrente), significa que quando a funcao for chamada, o codigo apontara para a conta que chamada
func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	//boleando as condicoes precisam ser verdadeiras
	// o valor do saque precisar ser maior que 0 && o valor do saque precisa ser menor ou igual ao saldo
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.SaldoContaCorrente

	if podeSacar {
		c.SaldoContaCorrente -= valorDoSaque
		return "Saque realiado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

// multiplos retornos

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	podeDepositar := valorDoDeposito > 0

	if podeDepositar {
		c.SaldoContaCorrente += valorDoDeposito
		return "Deposito realizado com Sucesso!", c.SaldoContaCorrente
	} else {
		return "Deposito nao efetuado", c.SaldoContaCorrente
	}

}

// alterando o valor da conta corrente
func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorDaTransferencia <= c.SaldoContaCorrente && valorDaTransferencia > 0 {
		c.SaldoContaCorrente -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return true
	} else {
		return false
	}

}

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

// verificando o contrato
type verificarConta interface {
	Sacar(valor float64) string
}
