package transport

const (
	defaultHdrLen = 20
	flagOffset    = 13
	urg           = "urg"
	ack           = "ack"
	psh           = "psh"
	rst           = "rst"
	syn           = "syn"
	fin           = "fin"
)

// Flags is a data structure containing various TCP flags (i.e. SYN, ACK, FIN).
type Flags map[string]bool

// NewFlags optionally initializes and returns a pointer to a `Flags` data structure.
// When the `initialize` parameter is passed in as `true`, all keys will be set to `false`.
// Otherwise, an empty `Flags` data structure is returned.
func NewFlags(initialize bool) *Flags {
	if initialize {
		f := &Flags{
			urg: false,
			ack: false,
			psh: false,
			rst: false,
			syn: false,
			fin: false,
		}
		return f
	}

	return &Flags{}
}

// Options is a data structure that signals the enabling of a TCP option
type Options struct {
	SelectiveAckPermitted bool
	MaximumSegmentSize    bool
}

// Segment is a data structure that represents a Transport layer transmission unit
type Segment struct {
	rawBytes []byte
	payOff   int
	flags    *Flags
	seqNum   int
	ackNum   int
	length   int
	wndsz    int
	checksum int
	options  *Options
}

// NewSegment initializes a new `Segment` data structure with sane default values.
// Returns a pointer to the newly created `Segment`.
func NewSegment() *Segment {
	return &Segment{
		rawBytes: make([]byte, defaultHdrLen, defaultHdrLen),
		payOff:   defaultHdrLen - 1,
	}
}

// Fill sets the contents of a `Segment` including both header and payload.
//
// A buffer containing the entire segment's raw bytes is passed in and stored.
// This is typically called when consuming Layer 3 payloads.
// TODO(judit): Parse bytes for the location of the payload offset
func (s *Segment) Fill(buf []byte) error {
	s.rawBytes = buf
	return nil
}

// FillPayload populates only the payload of a `Segment`.
//
// A buffer containing the raw bytes of the segment payload is passed in and stored.
// This is typically called when creating Layer 3 payloads.
func (s *Segment) FillPayload(buf []byte) error {
	s.rawBytes = append(s.rawBytes, buf...)
	return nil
}

// SetFlags modifies the state of the TCP flags associated with this segment.
//
// A `Flags` data structure is passed in which describes the desired modifications
// to be made to the current state of flags. Fields within the `Flags` struct
// should be set to `true` or `false` to explicitly define the desired modification.
// Alternatively, if a field is set to `nil`, no change will occur to that flag.
func (s *Segment) SetFlags(flags *Flags) error {
	if s.flags == nil {
		s.flags = NewFlags(false)
	}
	for flag, value := range *flags {
		(*s.flags)[flag] = value
	}

	return nil
}

// Flags returns the current state of the TCP flags associated with this segment.
// If flags were not already set via `SetFlags`, parse the flags out of the this
// segment's bytes.  If no bytes are associated, initialize and return a default
// flag set with all values set to `false`.
func (s *Segment) Flags() *Flags {
	if s.flags == nil {
		if len(s.rawBytes) > flagOffset {
			flags := byte(s.rawBytes[flagOffset])
			s.flags = parseFlags(flags)
		} else {
			s.flags = NewFlags(true)
		}
	}
	return s.flags
}
