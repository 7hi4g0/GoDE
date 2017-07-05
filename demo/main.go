package main

import (
	"fmt"
	"log"

	"github.com/7hi4g0/GoDE"
)

func main() {
	env := GoDE.Connect()
	defer env.Close()

	key := GoDE.Keyboard(1 << 24)

	for {
		select {
		case key = <-env.Key:
			fmt.Println(key.Key())
			fmt.Println(key.Mask())
		case env.CO <- "Ping":
			log.Println("Sent msg")
		}
	}
}
