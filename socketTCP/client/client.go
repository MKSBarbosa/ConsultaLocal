package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"time"
)

func main() {

	ClientTCP(10000)
}

func ClientTCP(n int) {

	// retorna o endereço do endpoint TCP
	r, err := net.ResolveTCPAddr("tcp", "localhost:1313")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	// connecta ao servidor (sem definir uma porta local)
	conn, err := net.DialTCP("tcp", nil, r)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	// fecha connexão
	defer func(conn *net.TCPConn) {
		err := conn.Close()
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
	}(conn)

	for i := 0; i < n; i++ {

		var Bairro = [5]string{"Madalena", "Espinheiro", "Gracas", "Parnamirim", "Pina"}
		// cria request
		req := Bairro[rand.Intn(4)]

		// envia mensage para o servidor
		start := time.Now()

		_, err := fmt.Fprintf(conn, req+"\n")
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}

		// recebe resposta do servidor
		rep, err := bufio.NewReader(conn).ReadString('\n')

		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		DurationR := time.Since(start)
		go HandleFile(int(DurationR))
		fmt.Print("Bairro Request:", req, " - Restaurante Sugerido:", rep)

		if i+1 == 10000 {
			fmt.Println("Tempo final:", time.Now().Second())
		}

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
