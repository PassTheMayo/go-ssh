package messages

import (
	"encoding/binary"
	"io"
)

type DisconnectReasonCode uint32

const (
	DisconnectReasonHostNotAllowedToConnect     DisconnectReasonCode = 1
	DisconnectReasonProtocolError               DisconnectReasonCode = 2
	DisconnectReasonKeyExchangeFailed           DisconnectReasonCode = 3
	DisconnectReasonReserved                    DisconnectReasonCode = 4
	DisconnectReasonMACError                    DisconnectReasonCode = 5
	DisconnectReasonCompressionError            DisconnectReasonCode = 6
	DisconnectReasonServiceNotAvailable         DisconnectReasonCode = 7
	DisconnectReasonProtocolVersionNotSupported DisconnectReasonCode = 8
	DisconnectReasonHostKeyNotVerifiable        DisconnectReasonCode = 9
	DisconnectReasonConnectionLost              DisconnectReasonCode = 10
	DisconnectReasonByApplication               DisconnectReasonCode = 11
	DisconnectReasonTooManyConnections          DisconnectReasonCode = 12
	DisconnectReasonAuthCancelledByUser         DisconnectReasonCode = 13
	DisconnectReasonNoMoreAuthMethodsAvailable  DisconnectReasonCode = 14
	DisconnectReasonIllegalUserName             DisconnectReasonCode = 15
)

type DisconnectMessage struct {
	ReasonCode  DisconnectReasonCode
	Description string
	LanguageTag string
}

func (p DisconnectMessage) ID() MessageType {
	return Disconnect
}

func (p DisconnectMessage) Marshal(w io.Writer) error {
	if err := binary.Write(w, binary.BigEndian, uint32(p.ReasonCode)); err != nil {
		return err
	}

	if err := binary.Write(w, binary.BigEndian, uint32(len(p.Description))); err != nil {
		return err
	}

	if _, err := w.Write([]byte(p.Description)); err != nil {
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

func (p *DisconnectMessage) Unmarshal(r io.Reader) error {
	var (
		reasonCode        uint32
		descriptionLength uint32
		languageTagLength uint32
	)

	if err := binary.Read(r, binary.BigEndian, &reasonCode); err != nil {
		return err
	}

	if err := binary.Read(r, binary.BigEndian, &descriptionLength); err != nil {
		return err
	}

	description := make([]byte, descriptionLength)

	if _, err := r.Read(description); err != nil {
		return nil
	}

	languageTag := make([]byte, languageTagLength)

	if _, err := r.Read(languageTag); err != nil {
		return nil
	}

	p.ReasonCode = DisconnectReasonCode(reasonCode)
	p.Description = string(description)
	p.LanguageTag = string(languageTag)

	return nil
}

var _ MessagePacket = &DisconnectMessage{}
