package classfile

import (
	"encoding/binary"
)

type (
	Reader struct {
		data []byte
		index int
	}
)

func NewReader(data []byte) *Reader {
	return &Reader{data, 0}
}

func (r *Reader) read1() u1 {
	r.index = r.index + 1
	return u1(r.data[r.index-1])
}

func (r *Reader) read2() u2 {
	val := binary.BigEndian.Uint16(r.data[r.index:])
	r.index = r.index + 2
	return u2(val)
}

func (r *Reader) read4() u4 {
	val := binary.BigEndian.Uint32(r.data[r.index:])
	r.index = r.index + 4
	return u4(val)
}

func (r *Reader) read8() u8 {
	val := binary.BigEndian.Uint64(r.data[r.index:])
	r.index = r.index + 8
	return u8(val)
}

func (r *Reader) readBytes(n int) (buf []byte) {
	buf = r.data[r.index:r.index + n]
	r.index = r.index + n
	return
}