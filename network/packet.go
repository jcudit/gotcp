package network

import (
	"encoding/binary"
	"encoding/hex"
)

const (
	defaultTTL    = 10
	defaultMTU    = 1500
	defaultHdrLen = 40
)

// Packet is a data structure that represents a Network layer transmission unit.
type Packet struct {
	rawBytes []byte
	payOff   int
	ttl      int
	version  int
}

// NewPacket initializes a new `Packet` data structure with sane default values.
// Returns a pointer to the newly created `Packet`.
func NewPacket() *Packet {
	return &Packet{
		// TODO(judit): Generate Layer 3 header from connection information associated
		// to the `Packet`. For now assume IPv6.
		rawBytes: make([]byte, defaultHdrLen, defaultHdrLen),
		payOff:   defaultHdrLen - 1,
		ttl:      -1,
	}
}

// Fill sets the contents of a `Packet` including both header and payload.
//
// A buffer containing the entire packet's raw bytes is passed in and stored.
// This is typically called when consuming Layer 2 payloads.
// TODO(judit): Parse bytes for the location of the payload offset
func (p *Packet) Fill(buf []byte) error {
	p.rawBytes = buf
	return nil
}

// FillPayload populates only the payload of a `Packet`.
//
// A buffer containing the raw bytes of the packet payload is passed in and stored.
// This is typically called when creating Layer 2 payloads.
func (p *Packet) FillPayload(buf []byte) error {
	p.rawBytes = append(p.rawBytes, buf...)
	return nil
}

// Bytes returns the raw bytes of this `Packet` including both header and payload.
//
// This is typically called when creating Layer 2 payloads.
func (p *Packet) Bytes() []byte {
	return p.rawBytes
}

// Version returns the Internet Protocol version of the `Packet`
func (p *Packet) Version() int {
	if p.ttl == -1 {
		p.ttl = int(uint16(p.rawBytes[0]) >> 4)
	}
	return p.ttl
}

// SetVersion assigns the Internet Protocol version of the `Packet`
func (p *Packet) SetVersion(v int) error {
	var b []byte
	binary.BigEndian.PutUint16(b, uint16(v))
	p.rawBytes[0] = b[0] >> 4
	return nil
}

// String returns a string representation of a packet conforming to the
// fmt.Stringer interface.
// TODO: Format bytes using hex representation
func (p *Packet) String() string {
	str := hex.EncodeToString(p.rawBytes)
	return str
}

// Length returns the number of bytes that make up this `Packet`.
func (p *Packet) Length() int {
	return len(p.rawBytes)
}
