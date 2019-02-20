package load

import (
	"github.com/YEXINGZHE54/myvm/pkg/utils"
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/instructions"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
)

const (
	aload_op = 0x19
	aload0_op = 0x2a
	aload1_op = 0x2b
	aload2_op = 0x2c
	aload3_op = 0x2d
)

type (
	ALoad struct {
		idx int
	}
)

var (
	aload0 = &ALoad{1}
	aload1 = &ALoad{2}
	aload2 = &ALoad{3}
	aload3 = &ALoad{4}
)

func (i *ALoad) Clone() instructions.Inst {
	if i.idx > 0 {
		return i
	}
	return &ALoad{}
}

func (i *ALoad) Fetch(coder *instructions.CodeReader) {
	if i.idx > 0 {
		return
	}
	i.idx = int(coder.Read1()) + 1
}

func (i *ALoad) Exec(f *stack.Frame) {
	utils.Log("executing instruction aload")
	aload(f, i.idx-1)
}

func aload(f *stack.Frame, idx int)  {
	f.PushOpstackRef(f.GetLocalRef(idx))
}

func init() {
	instructions.Register(aload_op, &ALoad{})
	instructions.Register(aload0_op, aload0)
	instructions.Register(aload1_op, aload1)
	instructions.Register(aload2_op, aload2)
	instructions.Register(aload3_op, aload3)
}
