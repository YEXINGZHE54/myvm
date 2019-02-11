package load

import (
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/instructions"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
)

const (
	aload1_op = 0x2b
)

type (
	Aload1Inst struct{}
)

func (i *Aload1Inst) Clone() instructions.Inst {
	return i
}

func (i *Aload1Inst) Fetch(coder *instructions.CodeReader) {

}

func (i *Aload1Inst) Exec(f *stack.Frame) {
	astore(f, 1)
}

func astore(f *stack.Frame, idx uint)  {
	f.PushOpstackRef(f.GetLocalRef(idx))
}

func init() {
	instructions.Register(aload1_op, &Aload1Inst{})
}
