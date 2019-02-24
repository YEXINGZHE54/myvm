package store

import (
	"github.com/YEXINGZHE54/myvm/pkg/utils"
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/instructions"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
)

const (
	lstore_op = 0x37
	lstore0_op = 0x3f
	lstore1_op = 0x40
	lstore2_op = 0x41
	lstore3_op = 0x42
)

type (
	LStore struct {
		idx int
	}
)

var (
	lstore0 = &LStore{1}
	lstore1 = &LStore{2}
	lstore2 = &LStore{3}
	lstore3 = &LStore{4}
)

func (i *LStore) Clone() instructions.Inst {
	if i.idx > 0 {
		return i
	}
	return &LStore{}
}

func (i *LStore) Fetch(coder *instructions.CodeReader) {
	if i.idx > 0 {
		return
	}
	i.idx = int(coder.Read1()) + 1
}

func (i *LStore) Exec(f *stack.Frame) {
	utils.Log("executing instruction lstore")
	lstore(f, i.idx-1)
}

func lstore(f *stack.Frame, idx int)  {
	f.SetLocalLong(idx, f.PopOpstackLong())
}

func init() {
	instructions.Register(lstore_op, &LStore{})
	instructions.Register(lstore0_op, lstore0)
	instructions.Register(lstore1_op, lstore1)
	instructions.Register(lstore2_op, lstore2)
	instructions.Register(lstore3_op, lstore3)
}
