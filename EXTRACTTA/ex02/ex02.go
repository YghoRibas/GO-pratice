package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Cliente struct {
	Nome string
}

func (c *Cliente) cadastrarClientes() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Informe o nome do cliente: ")
	scanner.Scan()
	c.Nome = scanner.Text()
}

type Conta struct {
	Cliente     Cliente
	Agencia     string
	NumeroConta string
	Saldo       float64
}

func (conta *Conta) gerarConta(cliente Cliente) {
	conta.Cliente = cliente
	conta.Agencia = "0001"
	conta.NumeroConta = "123456"
	conta.Saldo = 0.0
}

func (conta *Conta) realizarSaque(valorSaque float64) {
	if valorSaque <= conta.Saldo {
		conta.Saldo -= valorSaque
	} else {
		fmt.Println("Saldo insuficiente")
	}
}

func (conta *Conta) realizarDeposito(valorDeposito float64) {
	conta.Saldo += valorDeposito
}

func (conta *Conta) transferir(referenciaListCliente []Cliente, referenciaListConta []Conta) {
	// Transfer logic here
	fmt.Println("Feature in Development")
}

func (conta *Conta) getSaldo() float64 {
	return conta.Saldo
}

func main() {
	menu := Menu{
		ReferenciaListCliente: []Cliente{},
		ReferenciaListConta:   []Conta{},
	}

	menu.executarMenu()
}

type Menu struct {
	ReferenciaListCliente []Cliente
	ReferenciaListConta   []Conta
}

func (m *Menu) executarMenu() {
	scanner := bufio.NewScanner(os.Stdin)
	exit := false

	for !exit {
		fmt.Println("------------------------------")
		fmt.Println("|        Banco System        |")
		fmt.Println("|   O seu dinheiro está aqui |")
		fmt.Println("------------------------------")
		fmt.Println("| 1- Login                   |")
		fmt.Println("| 2- Criar Conta             |")
		fmt.Println("| 3- Sair                    |")
		fmt.Println("------------------------------")

		scanner.Scan()
		optionUser, _ := strconv.Atoi(scanner.Text())

		switch optionUser {
		case 1:
			login := Login{}
			contaUser := login.realizarLogin(m.ReferenciaListConta)
			m.acessarConta(contaUser)
		case 2:
			cliente := Cliente{}
			conta := Conta{}

			cliente.cadastrarClientes()
			conta.gerarConta(cliente)

			m.ReferenciaListCliente = append(m.ReferenciaListCliente, cliente)
			m.ReferenciaListConta = append(m.ReferenciaListConta, conta)
		case 3:
			fmt.Println("/---- Saindo ... ----/")
			exit = true
		}
	}
}

func (m *Menu) acessarConta(conta Conta) {
	scanner := bufio.NewScanner(os.Stdin)
	exit := false

	clienteLogado := conta.Cliente

	for !exit {
		fmt.Println("/********************************/")
		fmt.Printf("Olá %s\n", clienteLogado.Nome)
		fmt.Printf("Ag.: %s ---- Conta: %s\n", conta.Agencia, conta.NumeroConta)
		fmt.Println("/********************************/")
		fmt.Println("/***********   Menu   ***********/")
		fmt.Println("/* 1 - Consultar o saldo        */")
		fmt.Println("/* 2 - Realizar Saque           */")
		fmt.Println("/* 3 - Realizar Depósito        */")
		fmt.Println("/* 4 - Realizar Transferencia   */")
		fmt.Println("/* 5 - Alterar Senha            */")
		fmt.Println("/* 6 - Encerrar Sessão          */")
		fmt.Println("/********************************/")

		scanner.Scan()
		optionUser, _ := strconv.Atoi(scanner.Text())

		switch optionUser {
		case 1:
			fmt.Printf("Saldo: %.2f\n", conta.getSaldo())
		case 2:
			fmt.Print("Informe o valor que deseja sacar: ")
			scanner.Scan()
			valorSaque, _ := strconv.ParseFloat(scanner.Text(), 64)
			conta.realizarSaque(valorSaque)
			fmt.Printf("O valor em sua conta é: %.2f\n", conta.getSaldo())
		case 3:
			fmt.Print("Informe o valor que deseja depositar: ")
			scanner.Scan()
			valorDeposito, _ := strconv.ParseFloat(scanner.Text(), 64)
			conta.realizarDeposito(valorDeposito)
			fmt.Printf("O valor em sua conta é: %.2f\n", conta.getSaldo())
		case 4:
			conta.transferir(m.ReferenciaListCliente, m.ReferenciaListConta)
		case 5:
			fmt.Println("Feature in Development")
		case 6:
			exit = true
		}
	}
}

type Login struct{}

func (l *Login) realizarLogin(referenciaListConta []Conta) Conta {
	if len(referenciaListConta) > 0 {
		return referenciaListConta[0]
	} else {
		return Conta{}
	}
}
