package network

// Packet is a data structure that represents a Network layer transmission unit
type Packet struct {
	Header  []byte
	Payload []byte
}
