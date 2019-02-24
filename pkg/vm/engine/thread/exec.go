package thread

import (
	"fmt"
	"github.com/YEXINGZHE54/myvm/pkg/utils"
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/instructions"
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/reflect"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
	"github.com/YEXINGZHE54/myvm/pkg/vm/natives/throwable"
	"github.com/kr/pretty"
	"runtime"
)

func (t *Thread) Run() (err error) {
	err = t.initVMClass()
	if err != nil {
		return
	}
	return t.runMain()
}

func (t *Thread) initVMClass() (err error) {
	cls, err := t.vm.LoadClass("sun/misc/VM")
	if err != nil {
		return
	}
	t.InitClass(cls)
	return t.loop()
}

func (t *Thread) runMain() (err error) {
	// prepare class method
	c, err := t.vm.LoadClass(t.class)
	if err != nil {
		return
	}
	main, err := c.GetMain()
	if err != nil {
		return
	}
	// prepare args
	strclass, err := t.vm.LoadClass("java/lang/String")
	if err != nil {
		return
	}
	strarrcls, err := strclass.ArrayClass()
	if err != nil {
		return
	}
	args, err := strarrcls.NewArray(len(t.args))
	if err != nil {
		return
	}
	refs := args.Refs()
	for idx, arg := range t.args {
		refs[idx], err = c.Loader.JString(arg)
		if err != nil {
			return
		}
	}
	// run main method
	f := t.prepareFrame(main)
	f.SetLocalRef(0, args)
	// run clinit firstly
	t.InitClass(c)
	err = t.loop()
	return
}

func (t *Thread) InitClass(cls *reflect.Class) {
	cls.Started = true
	// push clinit method
	clinit, err := cls.GetClinit()
	if err == nil && clinit != nil {
		t.prepareFrame(clinit)
	}
	// init super class if not started
	if cls.Super != nil && !cls.Started {
		t.InitClass(cls.Super)
	}
	return
}

func (t *Thread) prepareFrame(method *reflect.Method) (f *stack.Frame) {
	f = stack.NewFrame(method)
	t.stack.Push(f)
	return
}

func (t *Thread) loop() (err error) {
	defer func() {
		if r := recover(); r != nil {
			if v, ok := r.(*reflect.Object); ok {
				fmt.Println("Exception: ", v.Class.Name)
				traces := v.Extra.([]*throwable.Trace)
				for _, tr := range traces {
					fmt.Printf("at %s.%s(%s:%d)\n", tr.ClassName, tr.MethodName, tr.FileName, tr.Line)
				}
			} else {
				t.Dump()
				pretty.Println(r)
			}
			var buf [40960]byte
			n := runtime.Stack(buf[:], true)
			println(string(buf[:n]))
		}
	}()
	for t.stack.Current() != nil {
		f := t.stack.Current()
		method := f.GetMethod()
		coder := instructions.NewCodeReader(method.Codes)
		pc := f.GetPC()
		t.PC(pc)
		coder.ResetPC(pc)
		opcode := coder.Read1()
		// get cpcode and operands
		utils.Log("pc = %d, op_code = %d", pc, opcode)
		inst := instructions.NewInst(opcode)
		if inst == nil {
			fmt.Printf("opcode %d not implemented\n", opcode)
		}
		inst.Fetch(coder)
		// set next pc and exec
		f.SetPC(coder.GetPC())
		inst.Exec(f)
	}
	return
}