package transport

// type TCPFlag byte

const (
	finMask byte = 1 << iota
	synMask
	rstMask
	pshMask
	ackMask
	urgMask
)

// parseFlags builds and returns a new `Flags` data structure corresponding to
// the TCP flag state passed in.
//
// Returns a pointer to the newly created `Flags` struct.
func parseFlags(flags byte) *Flags {
	f := NewFlags(true)

	if flags&urgMask != 0 {
		(*f)[urg] = true
	}
	if flags&ackMask != 0 {
		(*f)[ack] = true
	}
	if flags&pshMask != 0 {
		(*f)[psh] = true
	}
	if flags&rstMask != 0 {
		(*f)[rst] = true
	}
	if flags&synMask != 0 {
		(*f)[syn] = true
	}
	if flags&finMask != 0 {
		(*f)[fin] = true
	}

	return f
}
