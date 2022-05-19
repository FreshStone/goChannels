package main

import "errors"
import "log"
import "net"

func main() {
	listener, err := net.Listen("tcp", ":12345")
	if err != nil {
		log.Fatalln(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
		}
		// Handle each client connection in a new goroutine.
		go ClientHandler(conn)
	}
}

func ClientHandler(c net.Conn) {
	defer func() {
		if v := recover(); v != nil {
			log.Println("client handler panic:", v)
		}
		c.Close()
	}()
	panic(errors.New("just a demo.")) // a demo-purpose panic
}
