package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

// клиент реализован для автоматической проверки без telnet

const protocol, address = "tcp4", "localhost:11111"

func main() {

	connection, err := net.Dial(protocol, address)
	if err != nil {
		log.Fatal(err)
	}
	defer connection.Close()

	_, err = connection.Write([]byte("proverbs\n"))
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(connection)

	bArr, err := reader.ReadBytes('\n')
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Proverbs Response: ", string(bArr))
}
