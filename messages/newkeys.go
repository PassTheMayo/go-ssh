package messages

import (
	"io"
)

type NewKeysMessage struct{}

func (p NewKeysMessage) ID() MessageType {
	return NewKeys
}

func (p NewKeysMessage) Marshal(w io.Writer) error {
	return nil
}

func (p *NewKeysMessage) Unmarshal(r io.Reader) error {
	return nil
}

var _ MessagePacket = &NewKeysMessage{}
