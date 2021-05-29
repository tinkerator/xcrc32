// Package xcrc32 calculates composable CRC32 values in the style of
// what the GNU toolchains produce.
//
// This CRC algorithm appears to be common in embedded applications
// that use gdb for remote debugging for example.
package xcrc32

import (
	"sync"
)

var crcSetup sync.Once
var crcTable [256]uint32

// CRC32 holds a CRC32 value using the algorithm used by QuickFeather,
// which matches gdb. Apparently.
type CRC32 struct {
	value uint32
}

// NewCRC32 computes a CRC from some data. The algorithm appears to be
// custom (not obviously what golang natively supports).
func NewCRC32(data []byte) (*CRC32, uint32) {
	crcSetup.Do(func() {
		for i := uint32(0); i < 256; i++ {
			ch := i << 24
			for j := 8; j > 0; j-- {
				if ch&0x80000000 != 0 {
					ch = (ch << 1) ^ 0x04c11db7
				} else {
					ch = ch << 1
				}
			}
			crcTable[i] = ch
		}
	})
	crc := &CRC32{value: 0xffffffff}
	return crc, crc.Append(data)
}

// Append extends the data summarized by crc. Because this CRC32
// algorithm doesn't perform a final twist this appending can be
// arbitrarily repeated.
func (crc *CRC32) Append(data []byte) uint32 {
	for i := 0; i < len(data); i++ {
		crc.value = (crc.value << 8) ^ crcTable[(crc.value>>24)^uint32(data[i])]
	}
	return crc.value
}
