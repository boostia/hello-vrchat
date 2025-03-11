package main

import (
	"fmt"

	"github.com/hypebeast/go-osc/osc"
)

func main() {
	addr := "127.0.0.1:9001"
	d := osc.NewStandardDispatcher()

	if err := d.AddMsgHandler("/avatar/parameters/VelocityX", func(msg *osc.Message) {
		fmt.Printf("VelocityX : ")
		osc.PrintMessage(msg)
	}); err != nil {
		panic(err)
	}

	if err := d.AddMsgHandler("/avatar/parameters/VelocityY", func(msg *osc.Message) {
		fmt.Printf("VelocityY : ")
		osc.PrintMessage(msg)
	}); err != nil {
		panic(err)
	}

	if err := d.AddMsgHandler("/avatar/parameters/VelocityZ", func(msg *osc.Message) {
		fmt.Printf("VelocityZ : ")
		osc.PrintMessage(msg)
	}); err != nil {
		panic(err)
	}

	server := &osc.Server{
		Addr:       addr,
		Dispatcher: d,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
