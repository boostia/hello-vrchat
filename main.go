package main

import "github.com/hypebeast/go-osc/osc"

func main() {
	client := osc.NewClient("127.0.0.1", 9000)

	msg := osc.NewMessage("/chatbox/input")
	msg.Append("Hello, world!")

	if err := client.Send(msg); err != nil {
		panic(err)
	}
}
