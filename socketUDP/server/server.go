package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	ConsultaServerUDP()
	_, _ = fmt.Scanln()
}

func ConsultaServerUDP() {

	// define o endpoint do servidor UDP
	addr, err := net.ResolveUDPAddr("udp", "localhost:1313")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	// prepara o endpoint UDP para receber requests
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	fmt.Println("Servidor UDP aguardando conexão...")

	// fecha conn
	defer func(conn *net.UDPConn) {
		err := conn.Close()
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
	}(conn)

	for {
		req := make([]byte, 16)
		rep := make([]byte, 16)
		var reqMod []byte
		var size int

		// recebe request
		_, addr, err := conn.ReadFromUDP(req)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}

		// processa request
		for i := 0; i < 16; i++ {
			if req[i] == 0 {
				reqMod = make([]byte, 0, i)
				size = i
				break
			}
		}
		for j := 0; j < size; j++ {
			reqMod = append(reqMod, req[j])
		}
		rep = []byte(Consulta(strings.TrimSpace(string(reqMod))))
		// envia reposta
		_, err = conn.WriteTo(rep, addr)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
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
