package store

import (
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/instructions"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
)

const (
	istore_op = 0x36
	istore0_op = 0x3b
	istore1_op = 0x3c
	istore2_op = 0x3d
	istore3_op = 0x3e
)

type (
	IStore struct {
		idx int
	}
)

var (
	istore0 = &IStore{1}
	istore1 = &IStore{2}
	istore2 = &IStore{3}
	istore3 = &IStore{4}
)

func (i *IStore) Clone() instructions.Inst {
	if i.idx > 0 {
		return i
	}
	return &IStore{}
}

func (i *IStore) Fetch(coder *instructions.CodeReader) {
	if i.idx > 0 {
		return
	}
	i.idx = int(coder.Read1()) + 1
}

func (i *IStore) Exec(f *stack.Frame) {
	istore(f, i.idx-1)
}

func istore(f *stack.Frame, idx int)  {
	f.SetLocalVal(idx, f.PopOpstackVal())
}

func init() {
	instructions.Register(istore_op, &IStore{})
	instructions.Register(istore0_op, istore0)
	instructions.Register(istore1_op, istore1)
	instructions.Register(istore2_op, istore2)
	instructions.Register(istore3_op, istore3)
}
