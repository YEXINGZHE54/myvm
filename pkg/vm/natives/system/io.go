package system

import (
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
	"github.com/YEXINGZHE54/myvm/pkg/vm/natives"
)

func setIn0(f *stack.Frame)  {
	o := f.This()
	syscls := f.GetMethod().Cls
	field, err := syscls.LookupStaticField("in", "Ljava/io/InputStream;")
	if err != nil {
		panic(err)
	}
	syscls.SetField(field, o)
}

func setOut0(f *stack.Frame)  {
	o := f.This()
	syscls := f.GetMethod().Cls
	field, err := syscls.LookupStaticField("out", "Ljava/io/PrintStream;")
	if err != nil {
		panic(err)
	}
	syscls.SetField(field, o)
}

func setErr0(f *stack.Frame)  {
	o := f.This()
	syscls := f.GetMethod().Cls
	field, err := syscls.LookupStaticField("err", "Ljava/io/PrintStream;")
	if err != nil {
		panic(err)
	}
	syscls.SetField(field, o)
}

func init()  {
	natives.Register("java/lang/System", "setIn0", "(Ljava/io/InputStream;)V", setIn0)
	natives.Register("java/lang/System", "setOut0", "(Ljava/io/PrintStream;)V", setOut0)
	natives.Register("java/lang/System", "setErr0", "(Ljava/io/PrintStream;)V", setErr0)
}