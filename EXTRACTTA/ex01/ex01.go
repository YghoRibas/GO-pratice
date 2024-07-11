package main

import (
	"fmt"
	"time"
)

func fuelWorthCalculator(gas *float32, eta *float32) {
	clearScreen()
	fmt.Println("gasolina: ", *gas, " etanol: ", *eta)
	if *gas / *eta <= 1.3 {
		fmt.Println("Vale a pena abastercer com gasolina")
	} else {
		fmt.Println("Vale a pena abastercer com etanol")
	}
	time.Sleep(5 * time.Second)
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func printMainMenu(option *int, flag *bool) {
	clearScreen()
	fmt.Println("MENU")
	fmt.Println("1. Definir valores")
	fmt.Println("2. Calcular")
	fmt.Println("3. Sair")
	fmt.Scanln(option)

	if *option == 3 {
		*flag = false
	}
}

func printDefineValuesMenu(option *int, flag *bool, gas *float32, eta *float32) {
	clearScreen()
	fmt.Println("MENU")
	fmt.Println("1. Definir gasolina")
	fmt.Println("2. Definir etanol")
	fmt.Println("3. Voltar")
	fmt.Println("4. Sair")
	fmt.Scanln(option)

	if *option == 4 {
		*flag = false
	}
	switch *option {
	case 1:
		defineGasoline(gas)
	case 2:
		defineEthanol(eta)
	}
}

func defineGasoline(gas *float32) {
	clearScreen()
	fmt.Print("Informe o valor da gasolina: ")
	fmt.Scanln(gas)
}

func defineEthanol(eta *float32) {
	clearScreen()
	fmt.Print("Informe o valor do etanol: ")
	fmt.Scanln(eta)
}

func printMenus() {
	var flag = true
	var gas float32 = 1
	var eta float32 = 1
	for flag {
		var option int = 0
		printMainMenu(&option, &flag)
		switch option {
		case 1:
			printDefineValuesMenu(&option, &flag, &gas, &eta)
		case 2:
			fuelWorthCalculator(&gas, &eta)
		}
	}
}

func main() {
	printMenus()
}
