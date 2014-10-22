package circular

import "errors"

// TestVersion is required by tests, no idea why.
var TestVersion = 1

// Buffer represents a circular buffer.
type Buffer struct {
	buf        []byte
	head, tail int // head==tail -> empty, head+1==tail ->full
}

// NewBuffer creates a circular buffer.
func NewBuffer(size int) *Buffer {
	return &Buffer{buf: make([]byte, size+1)}
}

var (
	errEmpty = errors.New("empty Buffer")
	errFull  = errors.New("full Buffer")
)

// ReadByte reads a byte from tail of circular buffer.
func (b *Buffer) ReadByte() (byte, error) {
	if b.head == b.tail {
		return 0, errEmpty
	}
	reply := b.buf[b.tail]
	b.tail = (b.tail + 1) % len(b.buf)
	return reply, nil
}

// WriteByte writes a byte to head of circular buffer.
func (b *Buffer) WriteByte(c byte) error {
	newHead := (b.head + 1) % len(b.buf)
	if newHead == b.tail {
		return errFull
	}
	b.buf[b.head] = c
	b.head = newHead
	return nil
}

// Overwrite writes a byte to head of circular buffer, pushing tail if required.
func (b *Buffer) Overwrite(c byte) {
	newHead := (b.head + 1) % len(b.buf)
	if newHead == b.tail {
		b.tail = (newHead + 1) % len(b.buf)
	}
	b.buf[b.head] = c
	b.head = newHead
}

// Reset puts buffer in an empty state.
func (b *Buffer) Reset() {
	b.head = b.tail
}
