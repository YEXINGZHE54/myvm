package sync

import (
	"github.com/YEXINGZHE54/myvm/pkg/utils"
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/instructions"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
)

const (
	monitorexit_op = 0xc3
)

type (
	monitorexitInst struct {
	}
)

func (i *monitorexitInst) Clone() instructions.Inst {
	return i
}

func (i *monitorexitInst) Fetch(reader *instructions.CodeReader) {

}

func (i *monitorexitInst) Exec(f *stack.Frame) {
	utils.Log("executing instruction monitorexit")
	f.PopOpstackRef()
}

func init()  {
	instructions.Register(monitorexit_op, &monitorexitInst{})
}