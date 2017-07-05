package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net"

	"github.com/7hi4g0/GoDE"
)

type env struct {
	key   chan GoDE.Keyboard
	mouse chan GoDE.Mouse
	ci    chan string
	co    chan string
}

func main() {
	socket, err := net.Listen("unix", "socket")
	if err != nil {
		panic(err)
	}
	defer socket.Close()

	for {
		conn, err := socket.Accept()
		if err != nil {
			log.Println(err)
		}

		go serve(conn)
	}
}

func serve(conn net.Conn) {
	env := env{
		key:   make(chan GoDE.Keyboard, 5),
		mouse: make(chan GoDE.Mouse, 5),
		ci:    make(chan string, 5),
		co:    make(chan string, 5),
	}

	go reader(conn, env)
	// go writer(conn, env)
	for {
		select {
		case cmd := <-env.co:
			fmt.Println(cmd)
		}
	}
}

func reader(conn net.Conn, env env) {
	connDec := gob.NewDecoder(conn)

	for {
		var packet GoDE.Packet
		connDec.Decode(&packet)
		log.Println("Received msg")

		env.co <- packet.CO
	}
}

// func writer(conn net.Conn, env env) {
// 	connEnc := gob.NewEncoder(conn)

// 	for {
// 		var packet GoDE.Packet
// 		connEnc.Encode(&packet)

// 		env.key <- packet.Key
// 	}
// }
