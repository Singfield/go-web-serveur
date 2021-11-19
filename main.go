package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

// tcp serv
func main() {

	// Listen, Accept, read and write to the connection ;)
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}

	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println()
		}
		// goroutine for handle many connection
		go handle(conn)
	}

}

func handle(conn net.Conn) {

	scanner := bufio.NewScanner(conn)
	// pb inifite loop
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		// when stop listenning ?
	}
	defer conn.Close()
}
