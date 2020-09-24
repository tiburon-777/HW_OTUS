package main

import (
	"bufio"
	"errors"
	"io"
	"net"
	"time"
)

type TelnetClient interface {
	Connect() error
	Send() error
	Receive() error
	Close() error
}

type Client struct {
	address     string
	timeout     time.Duration
	in          io.ReadCloser
	inScanner   *bufio.Scanner
	out         io.Writer
	conn        net.Conn
	connScanner *bufio.Scanner
}

func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) TelnetClient {
	return &Client{address: address, timeout: timeout, in: in, out: out}
}

func (t *Client) Connect() error {
	var err error
	t.conn, err = net.DialTimeout("tcp", t.address, t.timeout)
	t.inScanner = bufio.NewScanner(t.in)
	t.connScanner = bufio.NewScanner(t.conn)
	return err
}

func (t *Client) Send() error {
	if !t.inScanner.Scan() {
		return errors.New("...EOF")
	}
	_, err := t.conn.Write(append(t.inScanner.Bytes(), '\n'))
	return err
}

func (t *Client) Receive() error {
	if !t.connScanner.Scan() {
		return errors.New("...connection closed by peer")
	}
	_, err := t.out.Write(append(t.connScanner.Bytes(), '\n'))
	return err
}

func (t *Client) Close() error {
	return t.conn.Close()
}
