package messages

import (
	"encoding/binary"
	"io"
)

type IgnoredDataMessage struct {
	Data string
}

func (p IgnoredDataMessage) ID() MessageType {
	return Ignore
}

func (p IgnoredDataMessage) Marshal(w io.Writer) error {
	if err := binary.Write(w, binary.BigEndian, uint32(len(p.Data))); err != nil {
		return err
	}

	if _, err := w.Write([]byte(p.Data)); err != nil {
		return err
	}

	return nil
}

func (p *IgnoredDataMessage) Unmarshal(r io.Reader) error {
	var dataLength uint32

	if err := binary.Read(r, binary.BigEndian, &dataLength); err != nil {
		return err
	}

	data := make([]byte, dataLength)

	if _, err := r.Read(data); err != nil {
		return nil
	}

	p.Data = string(data)

	return nil
}

var _ MessagePacket = &IgnoredDataMessage{}
