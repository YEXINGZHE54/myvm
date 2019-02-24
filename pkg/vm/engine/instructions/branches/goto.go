package branches

import (
	"github.com/YEXINGZHE54/myvm/pkg/utils"
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/instructions"
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/thread"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
)

const (
	goto_op = 0xa7
)

type (
	GotoInst struct {
		idx int16
	}
)

func (i *GotoInst) Clone() instructions.Inst {
	return &GotoInst{}
}

func (i *GotoInst) Fetch(coder *instructions.CodeReader) {
	i.idx = int16(coder.Read2())
}

func (i *GotoInst) Exec(f *stack.Frame) {
	gotoOffset(f, int(i.idx))
}

func gotoOffset(f *stack.Frame, offset int)  {
	utils.Log("executing instruction goto")
	t := f.Stack.Thread().(*thread.Thread)
	nextpc := t.GetPC() + int(offset)
	f.SetPC(nextpc)
}

func init()  {
	instructions.Register(goto_op, &GotoInst{})
}