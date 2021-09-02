// A simple chat server. It also has a (kind of) big problem...

package main

import (
	"fmt"
	"log"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	for {
		message := make([]byte, 128)
		n, err := conn.Read(message)
		if err != nil {
			log.Println(err.Error())
			break
		}
		log.Printf("Read %d bytes\n", n)

		if len(message) > 0 {
			msgStr := string(message)
			fmt.Println(msgStr)
		} else {
			break
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	for {
		if conn, err := listener.Accept(); err == nil {
			handleConnection(conn)
		}
	}
}
