package circular

import "errors"

const testVersion = 4

type Buffer struct {
	bytes                     []byte
	readPos, writePos, length int
}

func NewBuffer(size int) *Buffer {
	return &Buffer{bytes: make([]byte, size)}
}

func (buffer *Buffer) ReadByte() (byte, error) {
	if buffer.empty() {
		return 0, errors.New("Can't read, buffer is empty")
	}
	result := buffer.bytes[buffer.readPos]
	buffer.readPos = buffer.next(buffer.readPos)
	buffer.length--
	return result, nil
}

func (buffer *Buffer) WriteByte(b byte) error {
	if buffer.full() {
		return errors.New("Can't write, buffer is full")
	}
	buffer.write(b)
	return nil
}

func (buffer *Buffer) Overwrite(b byte) {
	if buffer.full() {
		buffer.readPos = buffer.next(buffer.readPos)
	}
	buffer.write(b)
}

func (buffer *Buffer) Reset() {
	buffer.readPos, buffer.writePos, buffer.length = 0, 0, 0
}

func (buffer *Buffer) write(b byte) {
	buffer.bytes[buffer.writePos] = b
	buffer.writePos = buffer.next(buffer.writePos)
	if !buffer.full() {
		buffer.length++
	}
}

func (buffer *Buffer) empty() bool {
	return buffer.length == 0
}

func (buffer *Buffer) full() bool {
	return buffer.length >= len(buffer.bytes)
}

func (buffer *Buffer) next(i int) int {
	return (i + 1) % len(buffer.bytes)
}
