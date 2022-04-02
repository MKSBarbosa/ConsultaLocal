package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func main() {

	ConsultaServerTCP()
	_, _ = fmt.Scanln()

}

func ConsultaServerTCP() {

	var user uint
	user = 0
	// define o endpoint do servidor TCP
	r, err := net.ResolveTCPAddr("tcp", "localhost:1313")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	// cria um listener TCP
	ln, err := net.ListenTCP("tcp", r)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	fmt.Println("Servidor TCP aguardando conexão...")

	for {
		// aguarda/aceita conexão
		conn, err := ln.Accept()
		user += 1
		fmt.Printf("Conecting with client %v, time start %v s.\n", user, time.Now().Second())
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}

		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	// recebe e processa requests

	// Close connection
	defer conn.Close()

	for {
		// recebe request terminado com '\n'
		req, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil && err.Error() == "EOF" {
			conn.Close()
			break
		}
		// processa request
		rep := Consulta(strings.TrimSpace(string(req)))

		// envia resposta
		_, err = conn.Write([]byte(rep + "\n"))
		if err != nil && err.Error() == "EOF" {
			conn.Close()
			break
		}
	}
}

func Consulta(Bairro string) (Restaurante string) {

	var DataBase = make(map[string]string)
	DataBase["Madalena"] = "Haro"
	DataBase["Espinheiro"] = "MammaMia"
	DataBase["Gracas"] = "ChinaInBox"
	DataBase["Parnamirim"] = "Kebab"
	DataBase["Pina"] = "Bistro"

	Restaurante = DataBase[Bairro]
	return
}
