package instructions

import (
	"github.com/YEXINGZHE54/myvm/pkg/utils"
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
		utils.Log("opcode not found: %d", opcode)
		return nil
	}
	return ins.Clone()
}

func ReadAll(codes []byte) (val []Inst) {
	coder := NewCodeReader(codes)
	for pc := 0; pc < len(codes); {
		code := coder.Read1()
		v := NewInst(code)
		if v == nil {
			return
		}
		v.Fetch(coder)
		pc = coder.GetPC()
		val = append(val, v)
	}
	return
}

func init() {
	supported = make(map[uint8] Inst)
}