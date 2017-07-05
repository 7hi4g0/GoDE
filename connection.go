package GoDE

import (
	"encoding/gob"
	"net"
)

type Env struct {
	Key   <-chan Keyboard
	key   chan Keyboard
	Mouse <-chan Mouse
	mouse chan Mouse
	CI    <-chan string
	ci    chan string
	CO    chan<- string
	co    chan string
}

func NewEnv() (env *Env) {
	env = &Env{key: make(chan Keyboard), mouse: make(chan Mouse), ci: make(chan string), co: make(chan string)}

	env.Key, env.Mouse, env.CI, env.CO = env.key, env.mouse, env.ci, env.co

	return
}

func (env *Env) Close() {
	close(env.key)
	close(env.mouse)
	close(env.ci)
	close(env.co)
}

func Connect() *Env {
	conn, err := net.Dial("unix", "socket")
	if err != nil {
		panic(err)
	}

	env := NewEnv()

	go reader(conn, env)
	go writer(conn, env)

	return env
}

func reader(conn net.Conn, env *Env) {
	var packet Packet

	connDec := gob.NewDecoder(conn)

	for {
		connDec.Decode(&packet)

		env.key <- packet.Key
	}
}

func writer(conn net.Conn, env *Env) {
	connEnc := gob.NewEncoder(conn)

	for {
		select {
		case cmd := <-env.co:
			packet := Packet{
				CO: cmd,
			}
			connEnc.Encode(packet)
		}
	}
}
