package instructions

type (
	CodeReader struct {
		codes []byte
		pc int
	}
)

func NewCodeReader(codes []byte) *CodeReader {
	return &CodeReader{codes, 0}
}

func (r *CodeReader) ResetPC(pc int) {
	r.pc = pc
}

func (r *CodeReader) GetPC() int {
	return r.pc
}

func (r *CodeReader) Read1() uint8 {
	idx := r.pc
	r.pc = r.pc + 1
	return r.codes[idx]
}

func (r *CodeReader) Read2() uint16 {
	val := (uint16(r.codes[r.pc]) << 8) | uint16(r.codes[r.pc+1])
	r.pc = r.pc + 2
	return val
}

func (r *CodeReader) Read4() uint32 {
	val := (uint32(r.codes[r.pc]) << 24) | 
		(uint32(r.codes[r.pc+1]) << 16) |
		(uint32(r.codes[r.pc+2]) << 8) |
		uint32(r.codes[r.pc+3])
	r.pc = r.pc + 4
	return val
}

func (r *CodeReader) SkipPaddings() {
	for r.pc & 3 != 0 {
		r.Read1()
	}
}

func (r *CodeReader) Peek(l int) []byte {
	return r.codes[r.pc:r.pc+l]
}