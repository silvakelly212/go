package main

//para o pacote ser reconhecido é necessario o comando go mod init (voce esta inicializando um modulo Go na pasta onde o arquivo esta)
import (
	"banco/contas"
	"fmt"
)

func main() {
	//Posso usar quando vou usar somente uma informacao
	// contaDaKelly := ContaCorrente{nomeTitular: "kelly", numeroAgencia: 152, numeroConta: 7111085, saldoContaCorrente: 125.05}
	// contaDaKelly2 := ContaCorrente{nomeTitular: "kelly", numeroAgencia: 152, numeroConta: 7111085, saldoContaCorrente: 125.05}
	// Posso usar quando vou ter que manipular mais informacoes
	// contaDaBruna := ContaCorrente{"Bruna", 222, 111222, 200}
	// fmt.Println(contaDaKelly)
	// fmt.Println(contaDaBruna)

	// // 02. Referencia ponteiros e metodos
	// // Quando e incluido *ContaCorrente esta apontando para o tipo informado na struct

	// var contaDoClaudio *ContaCorrente
	// contaDoClaudio = new(ContaCorrente)
	// contaDoClaudio.nomeTitular = "Claudio Arquiteto"
	// fmt.Println(contaDoClaudio)
	// fmt.Println(contaDaKelly, contaDaKelly2)

	//Comparando conteudos ser verdadeiro e false entre os conteudos
	//fmt.Println(contaDaKelly == contaDaKelly2)
	//& comparando enderecos (&contaDaKelly)
	//* comparando conteudo (*contaDaKelly)

	//Debito em conta corrente
	// // contaDoEdu := ContaCorrente{}
	// // contaDoEdu.nomeTitular = "Edu Guedes"
	// //contaDoEdu.saldoContaCorrente = 500

	// //fmt.Println(contaDoEdu.saldoContaCorrente)
	////fmt.Println(contaDoEdu.Sacar(50))

	//Deposito
	////fmt.Println(contaDoEdu.Depositar(100))

	//Transferencia

	// contaDoEdu := contas.ContaCorrente{Titular: "Edu Guedes", SaldoContaCorrente: 600}
	// contaDaAna := contas.ContaCorrente{Titular: "Ana Rickman", SaldoContaCorrente: 900}
	//tranferindo informando o & na  conta da na (&contaDaAna)
	// status := contaDoEdu.Transferir(50, &contaDaAna)

	// fmt.Println(status)
	// fmt.Println(contaDoEdu)
	// fmt.Println(contaDaAna)

	///// Composicao e encapsulamento 04
	// clienteRoberto := clientes.Titular{"Roberto Carlos", "123.111.123.12", "Desenvolvedor"}
	// contaDoRoberto := contas.ContaCorrente{clienteRoberto, 123, 123456, 100.00}
	// fmt.Println(contaDoRoberto)

	// ### Alterando a visibilidade
	// contaExemplo := contas.ContaCorrente{}
	// contaExemplo.Depositar(100)

	// fmt.Println(contaExemplo.SaldoContaCorrente)

	//####### conta poupanca

	// contaDoAngelo := contas.ContaPoupanca{}
	// contaDoAngelo.Depositar(100)
	// contaDoAngelo.Sacar(100.00)
	// fmt.Println(contaDoAngelo.ObterSaldo)

	//Interface Boleto Bancario - tipo contrato

	contaDoDenis := contas.ContaPoupanca{}
	contaDoDenis.Depositar(100)
	//&passando o endereco da conta
	contas.PagarBoleto(&contaDoDenis, 900)
	fmt.Println(contaDoDenis.ObterSaldo)

}
