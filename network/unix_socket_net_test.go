package network

import "testing"

func TestOpenCloseUnixSocket(t *testing.T) {
	nl := NewUnixSocketNetworkLayer()
	if err := nl.Open("/dev/net/gotcp-tun"); err != nil {
		t.Fatal(err)
	}
	if err := nl.Close(); err != nil {
		t.Fatal(err)
	}
}

func TestWriteUnixSocket(t *testing.T) {
	t.Fatal(nil)
}

func TestReadUnixSocket(t *testing.T) {
	t.Fatal(nil)
}
