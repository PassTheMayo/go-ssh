package messages

import (
	"encoding/binary"
	"io"
	"strings"
)

type KeyExchangeInitMessage struct {
	Cookie                              [16]byte
	KeyExchangeAlgorithms               []string
	ServerHostKeyAlgorithms             []string
	EncryptionAlgorithmsClientToServer  []string
	EncryptionAlgorithmsServerToClient  []string
	MACAlgorithmsClientToServer         []string
	MACAlgorithmsServerToClient         []string
	CompressionAlgorithmsClientToServer []string
	CompressionAlgorithmsServerToClient []string
	LanguagesClientToServer             []string
	LanguagesServerToClient             []string
	FirstKeyExchangePacketFollows       bool
	Reserved                            uint32
}

func (p KeyExchangeInitMessage) ID() MessageType {
	return KeyExchangeInit
}

func (p KeyExchangeInitMessage) Marshal(w io.Writer) error {
	var (
		keyExchangeAlgorithms               string = strings.Join(p.KeyExchangeAlgorithms, ",")
		serverHostKeyAlgorithms             string = strings.Join(p.ServerHostKeyAlgorithms, ",")
		encryptionAlgorithmsClientToServer  string = strings.Join(p.EncryptionAlgorithmsClientToServer, ",")
		encryptionAlgorithmsServerToClient  string = strings.Join(p.EncryptionAlgorithmsServerToClient, ",")
		macAlgorithmsClientToServer         string = strings.Join(p.MACAlgorithmsClientToServer, ",")
		macAlgorithmsServerToClient         string = strings.Join(p.MACAlgorithmsServerToClient, ",")
		compressionAlgorithmsClientToServer string = strings.Join(p.CompressionAlgorithmsClientToServer, ",")
		compressionAlgorithmsServerToClient string = strings.Join(p.CompressionAlgorithmsServerToClient, ",")
		languagesClientToServer             string = strings.Join(p.LanguagesClientToServer, ",")
		languagesServerToClient             string = strings.Join(p.LanguagesServerToClient, ",")
	)

	if _, err := w.Write(p.Cookie[:]); err != nil {
		return err
	}

	if err := binary.Write(w, binary.BigEndian, uint32(len(keyExchangeAlgorithms))); err != nil {
		return err
	}

	if _, err := w.Write([]byte(keyExchangeAlgorithms)); err != nil {
		return err
	}

	if err := binary.Write(w, binary.BigEndian, uint32(len(serverHostKeyAlgorithms))); err != nil {
		return err
	}

	if _, err := w.Write([]byte(serverHostKeyAlgorithms)); err != nil {
		return err
	}

	if err := binary.Write(w, binary.BigEndian, uint32(len(encryptionAlgorithmsClientToServer))); err != nil {
		return err
	}

	if _, err := w.Write([]byte(encryptionAlgorithmsClientToServer)); err != nil {
		return err
	}

	if err := binary.Write(w, binary.BigEndian, uint32(len(encryptionAlgorithmsServerToClient))); err != nil {
		return err
	}

	if _, err := w.Write([]byte(encryptionAlgorithmsServerToClient)); err != nil {
		return err
	}

	if err := binary.Write(w, binary.BigEndian, uint32(len(macAlgorithmsClientToServer))); err != nil {
		return err
	}

	if _, err := w.Write([]byte(macAlgorithmsClientToServer)); err != nil {
		return err
	}

	if err := binary.Write(w, binary.BigEndian, uint32(len(macAlgorithmsServerToClient))); err != nil {
		return err
	}

	if _, err := w.Write([]byte(macAlgorithmsServerToClient)); err != nil {
		return err
	}

	if err := binary.Write(w, binary.BigEndian, uint32(len(compressionAlgorithmsClientToServer))); err != nil {
		return err
	}

	if _, err := w.Write([]byte(compressionAlgorithmsClientToServer)); err != nil {
		return err
	}

	if err := binary.Write(w, binary.BigEndian, uint32(len(compressionAlgorithmsServerToClient))); err != nil {
		return err
	}

	if _, err := w.Write([]byte(compressionAlgorithmsServerToClient)); err != nil {
		return err
	}

	if err := binary.Write(w, binary.BigEndian, uint32(len(languagesClientToServer))); err != nil {
		return err
	}

	if _, err := w.Write([]byte(languagesClientToServer)); err != nil {
		return err
	}

	if err := binary.Write(w, binary.BigEndian, uint32(len(languagesServerToClient))); err != nil {
		return err
	}

	if _, err := w.Write([]byte(languagesServerToClient)); err != nil {
		return err
	}

	if err := binary.Write(w, binary.BigEndian, p.FirstKeyExchangePacketFollows); err != nil {
		return err
	}

	if err := binary.Write(w, binary.BigEndian, p.Reserved); err != nil {
		return err
	}

	return nil
}

func (p *KeyExchangeInitMessage) Unmarshal(r io.Reader) error {
	var (
		keyExchangeAlgorithmsLength               uint32
		serverHostKeyAlgorithmsLength             uint32
		encryptionAlgorithmsClientToServerLength  uint32
		encryptionAlgorithmsServerToClientLength  uint32
		macAlgorithmsClientToServerLength         uint32
		macAlgorithmsServerToClientLength         uint32
		compressionAlgorithmsClientToServerLength uint32
		compressionAlgorithmsServerToClientLength uint32
		languagesClientToServerLength             uint32
		languagesServerToClientLength             uint32
	)

	cookie := make([]byte, 16)

	if _, err := r.Read(cookie); err != nil {
		return err
	}

	if err := binary.Read(r, binary.BigEndian, &keyExchangeAlgorithmsLength); err != nil {
		return err
	}

	keyExchangeAlgorithms := make([]byte, keyExchangeAlgorithmsLength)

	if _, err := r.Read(keyExchangeAlgorithms); err != nil {
		return err
	}

	if err := binary.Read(r, binary.BigEndian, &serverHostKeyAlgorithmsLength); err != nil {
		return err
	}

	serverHostKeyAlgorithms := make([]byte, serverHostKeyAlgorithmsLength)

	if _, err := r.Read(serverHostKeyAlgorithms); err != nil {
		return err
	}

	if err := binary.Read(r, binary.BigEndian, &encryptionAlgorithmsClientToServerLength); err != nil {
		return err
	}

	encryptionAlgorithmsClientToServer := make([]byte, encryptionAlgorithmsClientToServerLength)

	if _, err := r.Read(encryptionAlgorithmsClientToServer); err != nil {
		return err
	}

	if err := binary.Read(r, binary.BigEndian, &encryptionAlgorithmsServerToClientLength); err != nil {
		return err
	}

	encryptionAlgorithmsServerToClient := make([]byte, encryptionAlgorithmsServerToClientLength)

	if _, err := r.Read(encryptionAlgorithmsServerToClient); err != nil {
		return err
	}

	if err := binary.Read(r, binary.BigEndian, &macAlgorithmsClientToServerLength); err != nil {
		return err
	}

	macAlgorithmsClientToServer := make([]byte, macAlgorithmsClientToServerLength)

	if _, err := r.Read(macAlgorithmsClientToServer); err != nil {
		return err
	}

	if err := binary.Read(r, binary.BigEndian, &macAlgorithmsServerToClientLength); err != nil {
		return err
	}

	macAlgorithmsServerToClient := make([]byte, macAlgorithmsServerToClientLength)

	if _, err := r.Read(macAlgorithmsServerToClient); err != nil {
		return err
	}

	if err := binary.Read(r, binary.BigEndian, &compressionAlgorithmsClientToServerLength); err != nil {
		return err
	}

	compressionAlgorithmsClientToServer := make([]byte, compressionAlgorithmsClientToServerLength)

	if _, err := r.Read(compressionAlgorithmsClientToServer); err != nil {
		return err
	}

	if err := binary.Read(r, binary.BigEndian, &compressionAlgorithmsServerToClientLength); err != nil {
		return err
	}

	compressionAlgorithmsServerToClient := make([]byte, compressionAlgorithmsServerToClientLength)

	if _, err := r.Read(compressionAlgorithmsServerToClient); err != nil {
		return err
	}

	if err := binary.Read(r, binary.BigEndian, &languagesClientToServerLength); err != nil {
		return err
	}

	languagesClientToServer := make([]byte, languagesClientToServerLength)

	if _, err := r.Read(languagesClientToServer); err != nil {
		return err
	}

	if err := binary.Read(r, binary.BigEndian, &languagesServerToClientLength); err != nil {
		return err
	}

	languagesServerToClient := make([]byte, languagesServerToClientLength)

	if _, err := r.Read(languagesServerToClient); err != nil {
		return err
	}

	if err := binary.Read(r, binary.BigEndian, &p.FirstKeyExchangePacketFollows); err != nil {
		return err
	}

	if err := binary.Read(r, binary.BigEndian, &p.Reserved); err != nil {
		return err
	}

	p.Cookie = [16]byte(cookie)
	p.KeyExchangeAlgorithms = strings.Split(string(keyExchangeAlgorithms), ",")
	p.ServerHostKeyAlgorithms = strings.Split(string(serverHostKeyAlgorithms), ",")
	p.EncryptionAlgorithmsClientToServer = strings.Split(string(encryptionAlgorithmsClientToServer), ",")
	p.EncryptionAlgorithmsServerToClient = strings.Split(string(encryptionAlgorithmsServerToClient), ",")
	p.MACAlgorithmsClientToServer = strings.Split(string(macAlgorithmsClientToServer), ",")
	p.MACAlgorithmsServerToClient = strings.Split(string(macAlgorithmsServerToClient), ",")
	p.CompressionAlgorithmsClientToServer = strings.Split(string(compressionAlgorithmsClientToServer), ",")
	p.CompressionAlgorithmsServerToClient = strings.Split(string(compressionAlgorithmsServerToClient), ",")
	p.LanguagesClientToServer = strings.Split(string(languagesClientToServer), ",")
	p.LanguagesServerToClient = strings.Split(string(languagesServerToClient), ",")

	return nil
}

var _ MessagePacket = &KeyExchangeInitMessage{}
