package messages

import (
	"encoding/binary"
	"io"
)

type UnimplementedMessage struct {
	SequenceID uint32
}

func (p UnimplementedMessage) ID() MessageType {
	return Unimplemented
}

func (p UnimplementedMessage) Marshal(w io.Writer) error {
	if err := binary.Write(w, binary.BigEndian, p.SequenceID); err != nil {
		return err
	}

	return nil
}

func (p *UnimplementedMessage) Unmarshal(r io.Reader) error {
	if err := binary.Read(r, binary.BigEndian, &p.SequenceID); err != nil {
		return err
	}

	return nil
}

var _ MessagePacket = &UnimplementedMessage{}
