package ssh

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"net"
)

type Server struct {
	listener net.Listener
}

func NewServer() *Server {
	return &Server{
		listener: nil,
	}
}

func (s *Server) acceptConnections() error {
	conn, err := s.listener.Accept()

	if err != nil {
		if errors.Is(err, net.ErrClosed) {
			return err
		}

		fmt.Println(err)
	}

	// TODO create Client struct and do proper shutdown on errors here vvv

	go func() {
		if err = s.handleConn(conn); err != nil {
			fmt.Println(err)
		}
	}()

	return nil
}

func (s *Server) handleConn(conn net.Conn) error {
	log.Printf("Received a connection from %s\n", conn.RemoteAddr())

	client := &Client{
		conn: conn,
		r:    bufio.NewReader(conn),
	}

	defer client.Close()

	if err := client.readVersionExchangePacket(); err != nil {
		return err
	}

	if err := client.sendVersionExchangePacket(); err != nil {
		return err
	}

	for {
		if err := client.readPacket(); err != nil {
			return err
		}
	}
}

func (s *Server) Listen(addr string) error {
	listener, err := net.Listen("tcp", addr)

	if err != nil {
		return err
	}

	s.listener = listener

	go func() {
		for {
			if err := s.acceptConnections(); err != nil {
				fmt.Println(err)

				break
			}
		}
	}()

	return nil
}

func (s *Server) ListenAndServe(addr string) error {
	listener, err := net.Listen("tcp", addr)

	if err != nil {
		return err
	}

	s.listener = listener

	for {
		if err := s.acceptConnections(); err != nil {
			return err
		}
	}
}

func (s *Server) Close() error {
	return s.listener.Close()
}
