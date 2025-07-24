// #### 2. Trabalhando com variáveis #####
// package main

// import "fmt"

// func main() {
// 	var nome string = "kelly"
// 	var versao float32 = 1.1
// 	var idade string = 49
// 	fmt.Println("Olá, srta.", nome, idade)
// 	fmt.Println("Este programa esta na versão", versao)

// }

// ou eu posso executar
// package main

// import (
// 	"fmt"
// 	"reflect"
// )

// func main() {
// 	var nome = "Kelly"
// 	var idade = 24
// 	var versao = 1.1
// 	fmt.Println("Olá sr.", nome, "sua idade é", idade)
// 	fmt.Println("Este programa está na versão", versao)
// 	// O reflect descobre qual e tipo de variavel
// 	fmt.Println("O tipo da variável versão é", reflect.TypeOf(versao))
// }

// trabalhando com variáveis flutuante
// func main() {
// 	var precoLeite float32 = 2.99
// 	var precoOvo float32 = 3.99
// 	var precoPao float32 = 0.99

//		fmt.Println("O preço do leite é R$", precoLeite)
//		fmt.Println("O preço do ovo é R$", precoOvo)
//		fmt.Println("O preço do pão é R$", precoPao)
//	}

// ##### INFERINDO UMA VARIAVEL SEM PRECISAR INFORMAR VAR NOME = "KELLY" AO INVES DE INFORMAR VAR SO ALTERAR
// = PARA := ENTAO O VAR SERA SUBSTITUIDO
// package main

// import (
// 	"fmt"
// 	"reflect"
// )

// func main() {
// 	nome := "Kelly"
// 	idade := 24
// 	versao := 1.1
// 	fmt.Println("Olá sr.", nome, "sua idade é", idade)
// 	fmt.Println("Este programa está na versão", versao)
// 	// O reflect descobre qual e tipo de variavel
// 	fmt.Println("O tipo da variável versão é", reflect.TypeOf(versao))
// }

// package main

// import "fmt"

// func main() {
// 	nome := "Kelly"
// 	versao := 1.1
// 	fmt.Println("Olá sr.", nome)
// 	fmt.Println("Este programa está na versão", versao)
// 	// EXIBINDO AS OPCOES DE MENU
// 	fmt.Println("1 - Iniciar Monitoramento")
// 	fmt.Println("2 - Exbir logs")
// 	fmt.Println("3 - Sair")
// 	var comando int
// 	//%d representa um numero inteiro
// 	// a variavel que esta guardando o menu seleciondo
// 	fmt.Scanf("%d", &comando)
// 	// para saber o endereço da variavel nós informamos o & na frente das variavel
// 	// descobrindo o endereco da variavel &comando
// 	fmt.Println("O endereço da minha variavel comando é", &comando)
// 	// ele aparece como exadecimal
// 	fmt.Println("O comando escolhido foi", comando)
// }

// ou posso alterar o comando fmt.Scanf para somente Scan

// package main

// import "fmt"

// func main() {
// 	nome := "Kelly"
// 	versao := 1.1
// 	fmt.Println("Olá sr.", nome)
// 	fmt.Println("Este programa está na versão", versao)
// 	// EXIBINDO AS OPCOES DE MENU
// 	fmt.Println("1 - Iniciar Monitoramento")
// 	fmt.Println("2 - Exbir logs")
// 	fmt.Println("3 - Sair")
// 	var comando int
// 	//%d representa um numero inteiro
// 	// a variavel que esta guardando o menu seleciondo
// 	fmt.Scan(&comando)
// 	// para saber o endereço da variavel nós informamos o & na frente das variavel
// 	// descobrindo o endereco da variavel &comando
// 	fmt.Println("O endereço da minha variavel comando é", &comando)
// 	// ele aparece como exadecimal
// 	fmt.Println("O comando escolhido foi", comando)
// }

////   ##### Mão na massa #### 
package main

import "fmt"

func main() {
	nome := "Kelly"
	versao := 1.1
	fmt.Println("Olá sr.", nome)
	fmt.Println("Este programa está na versão", versao)
	// EXIBINDO AS OPCOES DE MENU
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do Programa")

	var comando int
	fmt.Scan(&comando)
	fmt.Println("0 valor da variável comando é", comando)

}

