# DCON Golang Client (Master)
Package dcon provides Go implementation of DCON-master protocol.

## Description
Package provides a functionality, that allows to communicate
with devices implementing the DCON-slave protocol
(e.g. ICP DAS I-7000 I/O Modules), via serial port.

Current implementation is based on the ICP DAS specification:
http://ftp.icpdas.com/pub/cd/8000cd/napdos/7000/manual/7000dio.pdf

## Installation
```bash
$ go get github.com/Ayaya-zx/go-dcon
```

## Usage
Basic usage:

```go
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
```

Scan for devices:

```go
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
```