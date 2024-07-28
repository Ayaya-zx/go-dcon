package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Ayaya-zx/go-dcon"
)

func main() {
	handler := dcon.NewHandler()
	handler.BaudRate = 9600
	handler.DataBits = 8
	handler.Parity = dcon.NoParity
	handler.StopBits = dcon.OneStopBit

	err := handler.Connect("/dev/ttyUSB0")
	if err != nil {
		log.Fatal(err)
	}
	defer handler.Disconnect()

	err = handler.SetTimeout(300 * time.Millisecond)
	if err != nil {
		log.Fatal(err)
	}

	client := dcon.NewClient(handler)
	name, err := client.ReadName(2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(name)

	state, err := client.ReadDiscreteIOStatus(2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("First data bit: %08b\n", state[0])
	fmt.Printf("Second data bit: %08b\n", state[1])
}
