package messages

import (
	"encoding/binary"
	"io"
)

type DebugMessage struct {
	AlwaysDisplay bool
	Message       string
	LanguageTag   string
}

func (p DebugMessage) ID() MessageType {
	return Debug
}

func (p DebugMessage) Marshal(w io.Writer) error {
	if err := binary.Write(w, binary.BigEndian, p.AlwaysDisplay); err != nil {
		return err
	}

	if err := binary.Write(w, binary.BigEndian, uint32(len(p.Message))); err != nil {
		return err
	}

	if _, err := w.Write([]byte(p.Message)); err != nil {
		return err
	}

	if err := binary.Write(w, binary.BigEndian, uint32(len(p.LanguageTag))); err != nil {
		return err
	}

	if _, err := w.Write([]byte(p.LanguageTag)); err != nil {
		return err
	}

	return nil
}

func (p *DebugMessage) Unmarshal(r io.Reader) error {
	var (
		messageLength     uint32
		languageTagLength uint32
	)

	if err := binary.Read(r, binary.BigEndian, &p.AlwaysDisplay); err != nil {
		return err
	}

	if err := binary.Read(r, binary.BigEndian, &messageLength); err != nil {
		return err
	}

	message := make([]byte, messageLength)

	if _, err := r.Read(message); err != nil {
		return nil
	}

	languageTag := make([]byte, languageTagLength)

	if _, err := r.Read(languageTag); err != nil {
		return nil
	}

	p.Message = string(message)
	p.LanguageTag = string(languageTag)

	return nil
}

var _ MessagePacket = &DebugMessage{}
