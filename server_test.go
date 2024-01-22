package ssh_test

import (
	"testing"

	"github.com/PassTheMayo/go-ssh"
)

func TestServer(t *testing.T) {
	server := ssh.NewServer()

	if err := server.ListenAndServe("localhost:2022"); err != nil {
		t.Fatal(err)
	}
}
