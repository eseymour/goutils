package hexdump

import (
	"bytes"
	"fmt"
	"unicode"
)

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func makePrintable(r rune) rune {
	if !unicode.IsPrint(r) {
		return '.'
	}
	return r
}

// Bytes returns a string with the contents of data written in a hexdump style
// format with 16 bytes per line.
func Bytes(data []byte) string {
	var str string

	for i, n := 0, len(data); i < (15+n)/16; i++ {
		base := i * 16
		chunk := data[base:min(base+16, n)]
		chunkHex := fmt.Sprintf("% x", chunk)
		chunkStr := bytes.Map(makePrintable, chunk)
		str = str + fmt.Sprintf("%08x: %-47s  %s\n", base, chunkHex, chunkStr)
	}

	return str
}
