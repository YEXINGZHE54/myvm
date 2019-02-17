package objects

import (
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/reflect"
	thread2 "github.com/YEXINGZHE54/myvm/pkg/vm/engine/thread"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
)

func init_class(f *stack.Frame, cls *reflect.Class) (inited bool, err error) {
	inited = cls.Started
	if inited {
		return
	}
	thread := f.Stack.Thread().(*thread2.Thread)
	thread.InitClass(cls)
	return
}

func revertPC(f *stack.Frame)  {
	thread := f.Stack.Thread().(*thread2.Thread)
	f.SetPC(thread.GetPC())
}

func invokeMethod(f *stack.Frame, m *reflect.Method)  {
	newFrame := stack.NewFrame(m)
	f.Stack.Push(newFrame)
	for i := m.ArgSlot - 1; i >= 0; i = i - 1 {
		newFrame.SetLocalSlot(i, f.PopOpstackSlot())
	}
}