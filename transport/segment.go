package transport

// Segment is a data structure that represents a Transport layer transmission unit
type Segment struct {
	Header  []byte
	Payload []byte
}
