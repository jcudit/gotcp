package network

import "os"

// TunInterfaceNetworkLayer is a data structure that wraps a tun interface being
// used to read and write Network layer transmission units
type TunInterfaceNetworkLayer struct {
	fd *os.File
}

// NewTunInterfaceNetworkLayer creates a new `TunInterfaceNetworkLayer` data
// structure and returns a pointer to it
func NewTunInterfaceNetworkLayer() *TunInterfaceNetworkLayer {
	nl := &TunInterfaceNetworkLayer{fd: nil}
	return nl
}

// Open establishes a connection to a `tun` interface.
//
// If a device already exists that matches the `ifName` parameter, it will be
// connected.  If there is no matching device at the path that corresponds
// to `/dev/ifName`, then a new one will be created and connected.
func (t *TunInterfaceNetworkLayer) Open(ifName string) error {

	// Call out to OS-specific device opening routine
	file, err := openDevice(ifName)
	if err != nil {
		return err
	}

	// Call out to OS-specific device creation routine
	if err := createInterface(file, ifName); err != nil {
		file.Close()
		return err
	}

	t.fd = file
	return nil
}

// Close disconnects from the tun interface.
//
// If the interface isn't configured to be persistent, it is
// immediately destroyed by the kernel./
func (t *TunInterfaceNetworkLayer) Close() error {
	if err := t.fd.Close(); err != nil {
		return err
	}
	t.fd = nil
	return nil
}

// Read consumes a layer 3 packet from the kernel and emits the retrieved bytes
// as a `Packet` data structure.
func (t *TunInterfaceNetworkLayer) Read() (*Packet, error) {
	buf := make([]byte, 10000)

	boundary, err := t.fd.Read(buf)
	if err != nil {
		return nil, err
	}
	pkt := NewPacket()
	pkt.Fill(buf[:boundary])

	return pkt, nil
}

// Write sends an entire layer 3 packet through the tunnel interface.
//
// A buffer containing the raw bytes to be passed to the lower layer are
// passed in and wrapped in a `Packet` data structure on its way out of the
// tunnel interface.
func (t *TunInterfaceNetworkLayer) Write(buf []byte) error {
	pkt := NewPacket()
	pkt.FillPayload(buf)

	// Sanity check Packet payload and write out
	_, err := t.fd.Write(pkt.Bytes())
	if err != nil {
		return err
	}
	return nil
}
