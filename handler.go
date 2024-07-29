package dcon

import (
	"fmt"
	"time"

	"go.bug.st/serial"
)

type NotConnectedError string

func (e NotConnectedError) Error() string {
	return string(e)
}

func notConnectedError() error {
	return NotConnectedError("handler not connected")
}

// Handler provides low-level communication functionality.
type Handler struct {
	port serial.Port
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Connect(portName string, baudRate int) error {
	mode := &serial.Mode{
		BaudRate: baudRate,
		// Other values are fixed for both
		// the I-7000 and M-7000 series.
		DataBits: 8,
		Parity:   serial.NoParity,
		StopBits: serial.OneStopBit,
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
