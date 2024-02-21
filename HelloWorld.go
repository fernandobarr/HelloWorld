package main

import "fmt"

var mensagem string = "Olá/Hello"

func readMensagem() string {
	var nome string = ""
	fmt.Println("Digite seu nome: / Enter Your Name:")
	fmt.Scanln(&nome)
	return nome
}
func getMensagem(m string) string {
	return m + " " + readMensagem()
}

func main() {
	fmt.Println(getMensagem(mensagem))
}
