package store

import (
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/instructions"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
)

const (
	fstore_op = 0x38
	fstore0_op = 0x43
	fstore1_op = 0x44
	fstore2_op = 0x45
	fstore3_op = 0x46
)

type (
	FStore struct {
		idx int
	}
)

var (
	fstore0 = &FStore{1}
	fstore1 = &FStore{2}
	fstore2 = &FStore{3}
	fstore3 = &FStore{4}
)

func (i *FStore) Clone() instructions.Inst {
	if i.idx > 0 {
		return i
	}
	return &FStore{}
}

func (i *FStore) Fetch(coder *instructions.CodeReader) {
	if i.idx > 0 {
		return
	}
	i.idx = int(coder.Read1())
}

func (i *FStore) Exec(f *stack.Frame) {
	fstore(f, i.idx-1)
}

func fstore(f *stack.Frame, idx int)  {
	val := f.PopOpstackFloat()
	f.SetLocalVal(idx, int32(val))
}

func init() {
	instructions.Register(fstore_op, &FStore{})
	instructions.Register(fstore0_op, fstore0)
	instructions.Register(fstore1_op, fstore1)
	instructions.Register(fstore2_op, fstore2)
	instructions.Register(fstore3_op, fstore3)
}
