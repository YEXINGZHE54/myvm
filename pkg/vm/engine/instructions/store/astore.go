package store

import (
	"github.com/YEXINGZHE54/myvm/pkg/utils"
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/instructions"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
)

const (
	astore_op = 0x3a
	astore0_op = 0x4b
	astore1_op = 0x4c
	astore2_op = 0x4d
	astore3_op = 0x4e
)

type (
	AstoreInst struct {
		idx int
	}
)

var (
	astore0 = &AstoreInst{1}
	astore1 = &AstoreInst{2}
	astore2 = &AstoreInst{3}
	astore3 = &AstoreInst{4}
)

func (i *AstoreInst) Clone() instructions.Inst {
	if i.idx > 0 {
		return i
	}
	return &AstoreInst{}
}

func (i *AstoreInst) Fetch(coder *instructions.CodeReader) {
	if i.idx > 0 {
		return
	}
	i.idx = int(coder.Read1()) + 1
}

func (i *AstoreInst) Exec(f *stack.Frame) {
	utils.Log("executing instruction astore")
	astore(f, i.idx-1)
}

func astore(f *stack.Frame, idx int)  {
	ref := f.PopOpstackRef()
	f.SetLocalRef(idx, ref)
}

func init() {
	instructions.Register(astore_op, &AstoreInst{})
	instructions.Register(astore0_op, astore0)
	instructions.Register(astore1_op, astore1)
	instructions.Register(astore2_op, astore2)
	instructions.Register(astore3_op, astore3)
}
