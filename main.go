package main

import (
	"fmt"
	"io"
	"log"
	"net"
)
// ici on construit un serveur web tcp la couche en dessous de Http
// sur lequel on construira de quoi retourner des réponses formatté en hypertext protocol ( http)
func main() {
	// ici on accept toutes sort de connection en tcp sur le port 8080
	// on verra plus tard pour la reponse etc...
	// Listen(" quel type de connection, le port")
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
		io.WriteString(conn, "\nHello from TCP server\n")
		fmt.Fprintln(conn, "How is your day ?")
		fmt.Fprintln(conn, "%v", "well, I hope!")
		conn.Close()
	}

}
