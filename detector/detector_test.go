package detector

import (
	"fmt"
	"testing"
)

var BYTES = []byte{0xfd, 0x2f, 0x05, 0x6a, 0x9f, 0x50, 0x40, 0x0a, 0x09, 0x4a, 0x65, 0x72, 0x69, 0x20, 0x44, 0x69, 0x61, 0x7a, 0x15, 0x00, 0x00, 0x92, 0xc2, 0x1d}

func TestReadVarintReverse(t *testing.T) {
	idx := NewStringIndex(9)
	idx.String = "Jeri Diaz"
	i := IsPbString(BYTES, *idx)
	fmt.Println(i)
	//fmt.Println(string(BYTES[9:]))
}