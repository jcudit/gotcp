package network

import "github.com/jcudit/gotcp/transport"

// Layer is an interface that represents a read/write source of
// Network layer packets
type Layer interface {
	Open(path string) error
	Close() error
	Read() *Packet
	Write(*transport.Segment) error
}
