package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", "localhost:42069")
	if err != nil {
		log.Fatal("error", "error", err)
	}
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Fatal("error", "error", err)
	}
	defer conn.Close()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print('>')
		read, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("error", "error", err)
		}
		data := []byte(read)
		_, err = conn.Write(data)
		if err != nil {
			log.Fatal("error", "error", err)
		}
	}

}
