package xcrc32

import (
	"fmt"
	"testing"
)

func TestCRC32(t *testing.T) {
	tests := []struct {
		text  string
		value uint32
	}{
		{"Hello, World!", 0x19270120},
		{"The Hungry Caterpillar.", 0xdc5faf0e},
	}
	for i, test := range tests {
		_, full := NewCRC32([]byte(test.text))
		if full != test.value {
			t.Fatalf("test %d failed got:0x%x want:0x%x", i, full, test.value)
		}
		half := len(test.text) / 2
		crc, _ := NewCRC32([]byte(test.text[0:half]))
		final := crc.Append([]byte(test.text[half:]))
		if final != full {
			t.Fatalf("test %d failed composability test, got:0x%x want:0x%x", i, final, full)
		}
	}
}

func ExampleNewCRC32() {
	_, value := NewCRC32([]byte("Hello, World!"))
	fmt.Printf("xcrc32: 0x%x", value)
	// Output: xcrc32: 0x19270120
}
