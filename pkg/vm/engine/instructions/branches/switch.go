package branches

import (
	"github.com/YEXINGZHE54/myvm/pkg/utils"
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/instructions"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
)

const (
	lookswitch_op = 0xab
)

type (
	Match struct {
		val int32
		offset int
	}
	SwitchInst struct {
		off int
		matches []Match
	}
)

func (i *SwitchInst) Clone() instructions.Inst {
	return &SwitchInst{}
}

func (i *SwitchInst) Fetch(coder *instructions.CodeReader) {
	coder.SkipPaddings()
	i.off = int(coder.Read4())
	matches := int(coder.Read4())
	for idx := 0; idx < matches; idx = idx + 1 {
		i.matches = append(i.matches, Match{
			int32(coder.Read4()),
			int(coder.Read4()),
		})
	}
}

func (i *SwitchInst) Exec(f *stack.Frame) {
	utils.Log("executing instruction lookupswitch")
	k := f.PopOpstackVal()
	offset := i.off //default offset
	for _, m := range i.matches {
		if m.val == k {
			offset = m.offset
			break
		}
	}
	gotoOffset(f, offset)
}

func init() {
	instructions.Register(lookswitch_op, &SwitchInst{})
}