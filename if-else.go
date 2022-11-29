package main

import "fmt"

func main() {
	var nome string = "Lucas"

	if nome == "Gabriel" || nome == "gabriel" {
		fmt.Println("Olá Gabriel, bem-vindo!")
	} else {
		fmt.Printf("O usuário '%s' não foi reconhecido", nome)
	}
}
