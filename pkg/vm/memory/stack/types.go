package stack

import (
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/reflect"
)

type (
	Stack struct {
		max int
		frames []*Frame
		thread interface{}
	}
	Frame struct {
		stack *Stack
		method *reflect.Method
		localVars reflect.Slots
		opStack *OPStack
		pc int
	}
	OPStack struct {
		top uint
		slots reflect.Slots
	}
)