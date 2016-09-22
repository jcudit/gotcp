package network

import (
	"os/exec"
	"testing"
)

func TestOpenCloseTunInterface(t *testing.T) {
	nl := NewTunInterfaceNetworkLayer()

	if err := nl.Open("gotcp0"); err != nil {
		t.Fatal(err)
	}

	if err := nl.Close(); err != nil {
		t.Fatal(err)
	}
}

func TestReadTunInterface(t *testing.T) {
	nl := NewTunInterfaceNetworkLayer()
	if err := nl.Open("gotcp0"); err != nil {
		t.Fatal(err)
	}
	defer nl.Close()

	// Send a TCP/IPv6 SYN to the tunnel interface
	syn6 := func() error {
		cmd := exec.Command("nc", "-z", "-6", "fc00::2", "80")
		if err := cmd.Run(); err != nil {
			return err
		}
		return nil
	}
	go syn6()

	// Test that a packet can be read correctly from the tunnel interface
	// by checking the IP header's version field
	pkt, err := nl.Read()
	if err != nil {
		t.Fatal(err)
	}
	actual := pkt.Version()
	expected := 6
	if actual != expected {
		t.Fatalf("Packet version of %d != expected version of %d", actual, expected)
	}

}

func TestWriteTunInterface(t *testing.T) {
	nl := NewTunInterfaceNetworkLayer()
	if err := nl.Open("gotcp0"); err != nil {
		t.Fatal(err)
	}
	defer nl.Close()

	// Test that the packet was written correctly from the tunnel interface
	// by checking the IP header's version field
	pktChan := readPacket(nl)
	readPkt := <-pktChan

	// Send a layer 3 packet to the tunnel interface
	writtenPkt := NewPacket()
	nl.Write(writtenPkt.Bytes())

	actual := readPkt.Version()
	expected := 6
	if actual != expected {
		t.Fatalf("Packet version of %d != expected version of %d", actual, expected)
	}
}

func readPacket(nl *TunInterfaceNetworkLayer) <-chan *Packet {
	c := make(chan *Packet)
	read := func(chan<- *Packet) {
		pkt, _ := nl.Read()
		c <- pkt
	}
	go read(c)
	return c
}
