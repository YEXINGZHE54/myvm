package sync

import (
	"github.com/YEXINGZHE54/myvm/pkg/utils"
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/instructions"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
)

const (
	monitorenter_op = 0xc2
)

type (
	monitorenterInst struct {
	}
)

func (i *monitorenterInst) Clone() instructions.Inst {
	return i
}

func (i *monitorenterInst) Fetch(reader *instructions.CodeReader) {

}

func (i *monitorenterInst) Exec(f *stack.Frame) {
	utils.Log("executing instruction monitorenter")
	f.PopOpstackRef()
}

func init()  {
	instructions.Register(monitorenter_op, &monitorenterInst{})
}