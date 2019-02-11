package thread

import (
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/instructions"
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/reflect"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
)

func (t *Thread) Run() (err error) {
	// prepare class method
	c, err := t.vm.LoadClass(t.class)
	if err != nil {
		return
	}
	main, err := c.GetMain()
	if err != nil {
		return
	}
	// run main method
	t.prepareFrame(main, nil)
	err = t.loop()
	return
}

func (t *Thread) prepareFrame(method *reflect.Method, args []interface{}) {
	f := stack.NewFrame(method)
	t.stack.Push(f)
	return
}

func (t *Thread) loop() (err error) {
	for t.stack.Current() != nil {
		f := t.stack.Current()
		method := f.GetMethod()
		coder := instructions.NewCodeReader(method.Codes)
		pc := f.GetPC()
		t.PC(pc)
		coder.ResetPC(pc)
		opcode := coder.Read1()
		// get cpcode and operands
		inst := instructions.NewInst(opcode)
		inst.Fetch(coder)
		// set next pc and exec
		f.SetPC(coder.GetPC())
		inst.Exec(f)
	}
	return
}