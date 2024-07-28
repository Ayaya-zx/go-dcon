package dcon

import (
	"fmt"
	"time"

	"go.bug.st/serial"
)

type (
	Parity   int
	StopBits int
)

const (
	NoParity Parity = iota
	OddParity
	EvenParity
	MarkParity
	SpaceParity
)

const (
	OneStopBit StopBits = iota
	OnePointFiveStopBits
	TwoStopBits
)

type NotConnectedError string

func (e NotConnectedError) Error() string {
	return string(e)
}

func notConnectedError() error {
	return NotConnectedError("handler not connected")
}

type Handler struct {
	BaudRate int
	DataBits int
	Parity   Parity
	StopBits StopBits
	port     serial.Port
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Connect(portName string) error {
	mode := &serial.Mode{
		BaudRate: h.BaudRate,
		DataBits: h.DataBits,
		Parity:   serial.Parity(h.Parity),
		StopBits: serial.StopBits(h.StopBits),
	}

	port, err := serial.Open(portName, mode)
	if err != nil {
		return err
	}

	h.port = port
	return nil
}

func (h *Handler) Disconnect() error {
	if h.port == nil {
		return nil
	}
	if err := h.port.Close(); err != nil {
		return err
	}
	h.port = nil
	return nil
}

func (h *Handler) SetTimeout(timeout time.Duration) error {
	if h.port == nil {
		return notConnectedError()
	}

	return h.port.SetReadTimeout(timeout)
}

func (h *Handler) send(cmd []byte) ([]byte, error) {
	if h.port == nil {
		return nil, notConnectedError()
	}

	n, err := h.port.Write(cmd)
	if err != nil {
		return nil, err
	}
	if n != len(cmd) {
		return nil, fmt.Errorf("command was not written completely")
	}

	data := make([]byte, 100)
	n, err = h.port.Read(data)
	if err != nil {
		return nil, err
	}

	return data[:n], nil
}
