package class

import (
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/reflect"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
	"github.com/YEXINGZHE54/myvm/pkg/vm/natives"
	thread2 "github.com/YEXINGZHE54/myvm/pkg/vm/engine/thread"
	"strings"
)

func getName0(f *stack.Frame)  {
	this := f.GetLocalRef(0) // class object
	clsname := convert(this.Extra.(*reflect.Class).Name)
	o, err := f.GetMethod().Cls.Loader.JString(clsname)
	if err != nil {
		panic(err)
	}
	f.PushOpstackRef(o)
}

func forName0(f *stack.Frame)  {
	cname := f.This().GoString()
	init := f.GetLocalVal(1)
	cls, err := f.GetMethod().Cls.Loader.LoadClass(cname)
	if err != nil {
		panic(err)
	}
	if init > 0 && !cls.Started { // add initclass frame and revert forName0 pc
		thread := f.Stack.Thread().(*thread2.Thread)
		f.SetPC(thread.GetPC())
		thread.InitClass(cls)
		return
	}
	f.PushOpstackRef(cls.ToObject())
}

func convert(name string) string {
	return strings.Replace(name, "/", ".", -1)
}

func init()  {
	natives.Register("java/lang/Class", "getName0", "()Ljava/lang/String;", getName0)
	natives.Register("java/lang/Class", "forName0", "(Ljava/lang/String;ZLjava/lang/ClassLoader;Ljava/lang/Class;)Ljava/lang/Class;", forName0)
}