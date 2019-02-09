package thread

import (
	"myvm/pkg/vm/memory/stack"
)

func (t *Thread) Run() (err error) {
	f := stack.NewFrame(10, 10)
	f.SetLocalVal(0, 1)
	f.SetLocalVal(2, 100)
	f.SetLocalRef(4, nil)
	println("local 0:")
	println(f.GetLocalVal(0))
	println("local 2:")
	println(f.GetLocalVal(2))
	println("local 4:")
	println(f.GetLocalRef(4))
	f.PushOpstackVal(100)
	f.PushOpstackRef(nil)
	println(f.PopOpstackRef())
	println(f.PopOpstackVal())
	return
}