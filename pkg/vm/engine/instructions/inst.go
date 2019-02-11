package instructions

import (
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
)

type (
	Inst interface {
		Clone() Inst
		Fetch(coder *CodeReader)
		Exec(f *stack.Frame)
	}
)

var (
	supported map[uint8] Inst
)

func Register(opcode uint8, inst Inst) {
	supported[opcode] = inst
}

func NewInst(opcode uint8) Inst {
	ins := supported[opcode]
	if ins == nil {
		println("opcode not found: ")
		println(opcode)
		return nil
	}
	return ins.Clone()
}

func init() {
	supported = make(map[uint8] Inst)
}