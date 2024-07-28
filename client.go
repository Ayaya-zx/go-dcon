package dcon

import (
	"bytes"
	"fmt"
	"strconv"
)

type Client struct {
	handler *Handler
}

func NewClient(handler *Handler) *Client {
	return &Client{handler: handler}
}

func (c *Client) ReadName(addr int) (string, error) {
	var cmd bytes.Buffer
	cmd.WriteByte('$')
	cmd.WriteString(convertAddress(addr))
	cmd.Write([]byte{'M', '\r', '\n'})

	res, err := c.handler.send(cmd.Bytes())
	if err != nil {
		return "", err
	}

	if len(res) == 0 || res[0] != '!' {
		return "", fmt.Errorf("unable to execute command")
	}

	return string(res[1:]), nil
}

func (c *Client) ReadDiscreteIOStatus(addr int) ([2]byte, error) {
	var cmd bytes.Buffer

	cmd.WriteByte('@')
	cmd.WriteString(convertAddress(addr))
	cmd.Write([]byte{'\r', '\n'})

	res, err := c.handler.send(cmd.Bytes())
	if err != nil {
		return [2]byte{}, err
	}

	if len(res) < 6 {
		return [2]byte{}, fmt.Errorf("response was not read successfully")
	}

	if res[0] != '>' {
		return [2]byte{}, fmt.Errorf("unable to execute command")
	}

	firstData, err := strconv.ParseInt(string(res[1:3]), 16, 64)
	if err != nil {
		return [2]byte{}, err
	}

	secondData, err := strconv.ParseInt(string(res[3:5]), 16, 64)
	if err != nil {
		return [2]byte{}, err
	}

	return [2]byte{byte(firstData), byte(secondData)}, nil
}

func convertAddress(addr int) string {
	if addr < 10 {
		return "0" + strconv.Itoa(addr)
	}
	return strconv.Itoa(addr)
}
