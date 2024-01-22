package messages

import "io"

type MessagePacket interface {
	ID() MessageType
	Marshal(w io.Writer) error
	Unmarshal(r io.Reader) error
}
