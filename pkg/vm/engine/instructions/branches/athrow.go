package branches

import (
	"github.com/YEXINGZHE54/myvm/pkg/utils"
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/instructions"
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/reflect"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
)

const (
	athrow_op = 0xbf
)

type (
	AthrowInst struct {
	}
)

func (i *AthrowInst) Clone() instructions.Inst {
	return i
}

func (i *AthrowInst) Fetch(_ *instructions.CodeReader) {
	return
}

func (i *AthrowInst) Exec(f *stack.Frame) {
	utils.Log("executing instruction athrow")
	ex := f.PopOpstackRef()
	stack := f.Stack
	for stack.Current() != nil {
		frm := stack.Current()
		pc := frm.GetPC() - 1
		epc := searchExceptionHandlerPC(frm.GetMethod(), ex, uint16(pc))
		if epc >= 0 { // found
			frm.ClearOpstack()
			frm.PushOpstackRef(ex)
			frm.SetPC(epc)
			return
		}
		stack.Pop()
	}
	panic(ex)
}

func searchExceptionHandlerPC(method *reflect.Method, ex *reflect.Object, pc uint16) (epc int) {
	epc = -1
	for _, eh := range method.ExceptionTable {
		if eh.Start <= pc && eh.End > pc {
			if eh.Caught == nil { // caught all
				epc = int(eh.HandlerPC)
				return
			}
			if eh.Caught.Ref == nil {
				err := method.Cls.Loader.ResolveClass(eh.Caught)
				if err != nil {
					panic(err)
				}
			}
			if ex.Class.ExtendsOrSame(eh.Caught.Ref) {
				epc = int(eh.HandlerPC)
				return
			}
		}
	}
	return
}

func init()  {
	instructions.Register(athrow_op, &AthrowInst{})
}