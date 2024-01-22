package ssh

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"reflect"

	"github.com/PassTheMayo/go-ssh/messages"
)

type Client struct {
	conn net.Conn
	r    *bufio.Reader
}

func (c Client) readVersionExchangePacket() error {
	data, err := c.r.ReadBytes('\n')

	if err != nil {
		return err
	}

	dr := bufio.NewReader(bytes.NewReader(data))

	// Read suffix
	{
		prefix, err := dr.ReadBytes('-')

		if err != nil {
			return err
		}

		if !reflect.DeepEqual(prefix, []byte("SSH-")) {
			return fmt.Errorf("invalid version exchange value: %s", data)
		}
	}

	// Read client protocol version
	{
		version, err := dr.ReadBytes('-')

		if err != nil {
			return err
		}

		log.Printf("%s\n", version[:len(version)-1])
	}

	// Read client software
	{
		software, err := dr.ReadBytes(' ') // FIXME version exchange may not include a comment, so the proceeding terminator should be a new-line instead of a space

		if err != nil {
			return err
		}

		log.Printf("%s\n", software)
	}

	return nil
}

func (c Client) sendVersionExchangePacket() error {
	buf := &bytes.Buffer{}

	if _, err := buf.Write([]byte(fmt.Sprintf("SSH-2.0-GoSSH_1.0.0"))); err != nil {
		return err
	}

	if _, err := buf.Write([]byte{0x0A}); err != nil {
		return err
	}

	_, err := c.conn.Write(buf.Bytes())

	return err
}

func (c Client) readPacket() error {
	var (
		packetLength  uint32
		paddingLength byte
	)

	for {
		if err := binary.Read(c.r, binary.BigEndian, &packetLength); err != nil {
			return err
		}

		if err := binary.Read(c.r, binary.BigEndian, &paddingLength); err != nil {
			return err
		}

		payload := make([]byte, packetLength-uint32(paddingLength)-1)

		if _, err := c.r.Read(payload); err != nil {
			return err
		}

		randomPadding := make([]byte, paddingLength)

		if _, err := c.r.Read(randomPadding); err != nil {
			return err
		}

		if err := c.handlePacketData(payload); err != nil {
			return err
		}
	}
}

func (c Client) handlePacketData(data []byte) error {
	r := bytes.NewReader(data)

	var messageType byte

	if err := binary.Read(r, binary.BigEndian, &messageType); err != nil {
		return err
	}

	switch messages.MessageType(messageType) {
	case messages.Disconnect:
		return nil
	}

	return nil
}

func (c Client) Close() error {
	return c.conn.Close()
}
