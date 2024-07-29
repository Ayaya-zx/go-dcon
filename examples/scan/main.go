package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Ayaya-zx/go-dcon"
)

func main() {
	handler := dcon.NewHandler()

	err := handler.Connect("/dev/ttyUSB0", 9600)
	if err != nil {
		log.Fatal(err)
	}
	defer handler.Disconnect()

	err = handler.SetTimeout(300 * time.Millisecond)
	if err != nil {
		log.Fatal(err)
	}

	devs := dcon.Scan(handler)

	fmt.Printf("Found devices: %v\n", devs)
}
