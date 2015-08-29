package humanhex

import (
	"bytes"
	"fmt"
)

// eww, this is bad and i feel bad
func readable(b byte) bool {
	switch {
	case b < 0x20:
		return false
	case b > 0x7f:
		return false
	default:
		return true
	}
}

// less eww
func lookahead(b []byte, p int) int {
	var i = p
	for ; i < len(b); i++ {
		if !readable(b[i]) {
			break
		}
	}
	return i - p
}

// String will return a more human-consumable representation of the bytes in b.
// m specifies the number of contiguous human-friendly bytes required before we
// will display it directly. 2 is a good default. 0 is undefined and panics.
// Note that human-consumable and human-friendly here are limited in scope to
// North American, English-speaking humans. Maybe Go's rune type can do some
// more heavy lifting on our behalf down the road?
func String(b []byte, m int) string {
	if m == 0 {
		panic("invalid value 0 for m")
	}
	out := &bytes.Buffer{}
	for i := 0; i < len(b); {
		if contig := lookahead(b, i); contig >= m {
			out.WriteString(string(b[i : i+contig]))
			i += contig
		} else {
			out.WriteString(fmt.Sprintf("\\x%02x", b[i]))
			i++
		}
	}
	return out.String()
}
