package thread

import (
	"myvm/pkg/vm/engine/instructions"
	"myvm/pkg/vm/memory/stack"
	"myvm/pkg/vm/loader/classfile"
	"github.com/kr/pretty"
)

func (t *Thread) Run() (err error) {
	c, err := t.vm.LoadClass(t.class)
	if err != nil {
		return
	}
	pretty.Println(c)
	//main := cf.GetMain()
	//err = t.interpret(main, nil)
	return
}

func (t *Thread) interpret(method *classfile.Member, args []interface{}) error {
	code := method.GetCode()
	coder := code.GetReader()
	f := stack.NewFrame(int(code.MaxStacks), int(code.MaxLocals))
	// running bytecode
	for !f.Exited() {
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
	return nil
}