package humanhex

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

// String will return a more human-consumable representation of the bytes in b.
// m specifies the number of contiguous human-friendly bytes required before we
// will display it directly. 2 is a good default.
func String(b []byte, m uint) string {

	return ""
}
