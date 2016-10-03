package transport

import (
	"fmt"
	"strconv"
	"testing"
)

func TestSetFlags(t *testing.T) {
	segment := NewSegment()
	flags := NewFlags(false)

	(*flags)[ack] = true
	(*flags)[syn] = true
	segment.SetFlags(flags)

	actual := segment.Flags()
	if (*actual)[ack] != true || (*actual)[syn] != true {
		t.Fatalf("Flags set using `SetFlags()` failed to remain after `Flags()`")
	}
	for _, flag := range []string{urg, psh, rst, fin} {
		if (*actual)[flag] != false {
			t.Fatalf("Flags unset using `SetFlags()` were set after `Flags()`")
		}
	}
}

func TestParseFlags(t *testing.T) {
	christmas := make([]byte, defaultHdrLen, defaultHdrLen)
	christmas[flagOffset] = 0xff
	segment := Segment{rawBytes: christmas}
	actual := segment.Flags()
	for _, flag := range []string{urg, ack, psh, rst, syn, fin} {
		if (*actual)[flag] != true {
			bits := fmt.Sprintf(strconv.FormatInt(int64(christmas[flagOffset]), 2))
			t.Fatalf("Failed to parse flag %s from bits %s", flag, bits)
		}
	}
}
