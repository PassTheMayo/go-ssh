package messages

import (
	"encoding/binary"
	"io"
)

type ServiceAcceptMessage struct {
	ServiceName string
}

func (p ServiceAcceptMessage) ID() MessageType {
	return ServiceAccept
}

func (p ServiceAcceptMessage) Marshal(w io.Writer) error {
	if err := binary.Write(w, binary.BigEndian, uint32(len(p.ServiceName))); err != nil {
		return err
	}

	if _, err := w.Write([]byte(p.ServiceName)); err != nil {
		return err
	}

	return nil
}

func (p *ServiceAcceptMessage) Unmarshal(r io.Reader) error {
	var serviceNameLength uint32

	if err := binary.Read(r, binary.BigEndian, &serviceNameLength); err != nil {
		return err
	}

	serviceName := make([]byte, serviceNameLength)

	if _, err := r.Read(serviceName); err != nil {
		return nil
	}

	p.ServiceName = string(serviceName)

	return nil
}

var _ MessagePacket = &ServiceAcceptMessage{}
