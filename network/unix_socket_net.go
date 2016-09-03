package network

import "net"

const (
	cUNIX = "unix"
)

// UnixSocketNetworkLayer is a data structure that wraps a Unix socket being
// used to read and write Network layer transmission units
type UnixSocketNetworkLayer struct {
	conn *net.UnixConn
}

// NewUnixSocketNetworkLayer creates a new `UnixSocketNetworkLayer` data
// structure and returns a pointer to it
func NewUnixSocketNetworkLayer() *UnixSocketNetworkLayer {
	nl := &UnixSocketNetworkLayer{conn: nil}
	return nl
}

// Open establishes a connection to an existing Unix socket
func (s *UnixSocketNetworkLayer) Open(path string) error {
	conn, err := net.DialUnix(cUNIX, nil, &net.UnixAddr{Name: path, Net: cUNIX})
	if err != nil {
		return err
	}
	s.conn = conn
	return nil
}

// Close closes a connection to an existing Unix socket
func (s *UnixSocketNetworkLayer) Close() error {
	if err := s.conn.Close(); err != nil {
		return err
	}
	s.conn = nil
	return nil
}

// Read performs a read on the Unix socket returning a Layer 3 packet
func (s *UnixSocketNetworkLayer) Read() (*Packet, error) {
	return nil, nil
}

// Write performs a write on the Unix socket, passing a Layer 3 packet
func (s *UnixSocketNetworkLayer) Write() error {

	// _, err = conn.Write([]byte("hello"))
	// if err != nil {
	// 	t.Fatal(err)
	// }
	return nil
}
