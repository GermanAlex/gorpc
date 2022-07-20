package main

import (
	"bufio"
	"log"
	"math/rand"
	"net"
	"os"
	"strings"
	"time"
)

const protocol, address = "tcp4", ":11111"

func main() {
	listener, err := net.Listen(protocol, address)
	if err != nil {
		log.Fatal(err)
	}
	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go proverbsConn(connection)
	}
}

func proverbsConn(connection net.Conn) {
	defer connection.Close()
	reader := bufio.NewReader(connection)
	b, err := reader.ReadBytes('\n')
	if err != nil {
		log.Println(err)
		return
	}

	message := strings.TrimSuffix(string(b), "\n")
	message = strings.TrimSuffix(message, "\r")

	if message == "proverbs" {
		result, err := getProverbs()
		if err != nil {
			log.Println(err)
			return
		}
		connection.Write([]byte(result + "\n"))
	}

}

func getProverbs() (string, error) {
	var proverbsArr []string

	file, err := os.OpenFile("goproverbs.txt", os.O_RDONLY, 0777)

	if err != nil {
		return "", err
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		proverbsArr = append(proverbsArr, scanner.Text())
	}

	file.Close()

	rand.Seed(time.Now().Unix())
	return proverbsArr[rand.Intn(len(proverbsArr))], nil
}
