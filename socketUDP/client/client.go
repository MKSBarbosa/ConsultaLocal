package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"time"
)

func main() {
	HelloClientUDP(10000)
}

func HelloClientUDP(n int) {
	// req := make([]byte, 1024)
	rep := make([]byte, 1024)

	// retorna o endereço do endpoint UDP
	addr, err := net.ResolveUDPAddr("udp", "localhost:1313")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	// conecta ao servidor -- não cria uma conexão
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	// desconecta do servidor
	defer func(conn *net.UDPConn) {
		err := conn.Close()
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
	}(conn)

	var Bairro = [5]string{"Madalena", "Espinheiro", "Gracas", "Parnamirim", "Pina"}

	for i := 0; i < n; i++ {
		// cria request
		start := time.Now()
		varTeste := Bairro[rand.Intn(4)]
		req := []byte(varTeste)

		// envia request ao servidor
		_, err = conn.Write(req)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}

		// recebe resposta do servidor
		_, _, err := conn.ReadFromUDP(rep)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		DurationR := time.Since(start)
		go HandleFile(int(DurationR))
		fmt.Println(string(req), " -> ", string(rep))
	}
}

func HandleFile(time int) {

	var file *os.File

	_, err := os.Stat("time.txt")
	if err != nil {
		if os.IsNotExist(err) {
			log.Println("File does exist. Creating file ...")
			_, err = os.Create("time.txt")
			if err != nil {
				os.Exit(0)
			}
		}
	}

	// Open the file
	file, err = os.OpenFile("time.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}

	// Close the file after the use
	defer file.Close()

	//Writing yhe time in the file
	if _, err := file.WriteString(strconv.Itoa(time) + "\n"); err != nil {
		log.Fatal(err)
	}
}
