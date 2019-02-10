package stack

import (
	"myvm/pkg/vm/engine/data"
)

type (
	Stack struct {
		max int
		frames []*Frame
	}
	Frame struct {
		localVars []Slot
		opStack *OPStack
		pc int
	}
	OPStack struct {
		top int
		slots []Slot
	}
	Slot struct {
		val int
		ref *data.Object
	}
)