package code_gen

import (
	"bytes"
	"encoding/binary"
)

func encodeSleb128(b *bytes.Buffer, v int32) {
	for {
		c := uint8(v & 0x7f)
		s := uint8(v & 0x40)
		v >>= 7
		if (v != -1 || s == 0) && (v != 0 || s != 0) {
			c |= 0x80
		}
		b.WriteByte(c)
		if c&0x80 == 0 {
			break
		}
	}
}

func float64ToByte(f float64) []byte {
	var buf bytes.Buffer
	err := binary.Write(&buf, binary.LittleEndian, f)
	if err != nil {
		panic("binary.Write failed: " + err.Error())
	}
	return buf.Bytes()
}
