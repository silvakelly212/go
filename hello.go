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
// package main

// import "fmt"

// func main() {
// 	nome := "Kelly"
// 	versao := 1.1
// 	fmt.Println("Olá sr.", nome)
// 	fmt.Println("Este programa está na versão", versao)
// 	// EXIBINDO AS OPCOES DE MENU
// 	fmt.Println("1 - Iniciar Monitoramento")
// 	fmt.Println("2 - Exibir logs")
// 	fmt.Println("0 - Sair do Programa")

// 	var comando int
// 	fmt.Scan(&comando)
// 	fmt.Println("0 valor da variável comando é", comando)

// }

// 03 Controlando o fluxo do script

// package main

// import "fmt"

// func main() {
// 	nome := "Kelly"
// 	versao := 1.1
// 	fmt.Println("Olá sr.", nome)
// 	fmt.Println("Este programa está na versão", versao)

// 	fmt.Println("1 - Iniciar Monitoramento")
// 	fmt.Println("2 - Exibir logs")
// 	fmt.Println("0 - Sair do Programa")

// 	var comando int
// 	fmt.Scan(&comando)
// 	fmt.Println("O comando escolhido foi", comando)

// 	switch comando {
// 	case 1:
// 		fmt.Println("Monitorando...")
// 	case 2:
// 		fmt.Println("Exibindo Logs...")
// 	case 0:
// 		fmt.Println("Saindo do programa...")
// 	default:
// 		fmt.Println("Não conheço este comando")
// 	}
// }

//introdução as funçoes

package main

//possui as função de impressão
import (
	"fmt"
	// função de sair os
	"os"
)

func main() {
	//chamando a função exibe introdução
	exibeIntroduvao()
//loop infinito
//     for {
// 		exibeMenu()
// 		comando := leComando()

// 	switch comando {
// 	case 1:
// 		iniciarMonitoramento()
// 	case 2:
// 		fmt.Println("Exibindo Logs...")
// 	case 0:
// 		fmt.Println("Saindo do programa...")
// 		os.Exit(0)
// 	default:
// 		fmt.Println("Não conheço este comando")
// 		os.Exit(-1)
// 	}
// }
	//chamando a função de exibir menu
	exibeMenu()
	// chamando a função ler comando em int
	comando := leComando()

	switch comando {
	case 1:
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Exibindo Logs...")
	case 0:
		fmt.Println("Saindo do programa...")
		os.Exit(0)
	default:
		fmt.Println("Não conheço este comando")
		os.Exit(-1)
	}
}

// Criando funcoes no Go

func exibeIntroduvao() {
	nome := "Kelly"
	versao := 1.1
	fmt.Println("Olá sr.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)
	return comandoLido
}

func exibeMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do Programa")

}
//Fazendo requisicão web acessando o site alura
//Função que retorna multiplos valores

// func iniciarMonitoramento() {
// 	fmt.Println("Monitoramento....")
// 	site := "https://www.alura.com.br"
// 	resp, _ := http.Get(site)
// 	fmt.Println(resp)
// }

// Para saber o retorno do status no GO
func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	// site com URL inexistente
	site := "https://httpbin.org/status/404" // ou 200
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
	}
}

//Trabalhando com lista slice que funciona em cima de um array Go

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	var sites [4]string
	sites[0] = "https://random-status-code.herokuapp.com/"
	sites[1] = "https://www.alura.com.br"
	sites[3] = "https://www.caelum.com.br"

	fmt.Sprintln(sites)

	site := "https://random-status-code.heoruapp.com/"
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", sites, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", sites, "está com problemas. Status Code:", resp.StatusCode)
	}
}

}

func exibeNomes() {
	//slice simula uma arry
	nomes := []string{"Douglas", "Daniel", "Bernardo"}

	fmt.Println("O meu slice tem", len(nomes), "itens")
	fmt.Println("O meu slice tem capacidade para", cap(nomes), "itens")
	//acrescenta o numero de posição
	nomes = append(nomes, "Aparecida")
	fmt.Println("O meu slice tem", len(nomes), "itens")
	fmt.Println("O meu slice tem capacidade para", cap(nomes), "itens")
}

//
