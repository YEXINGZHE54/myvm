package throwable

import (
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
	"github.com/YEXINGZHE54/myvm/pkg/vm/natives"
)

type (
	Trace struct {
		FileName string
		ClassName string
		MethodName string
		Line int
	}
)

func fillInStackTrace(f *stack.Frame)  {
	this := f.GetLocalRef(0)
	f.PushOpstackRef(this)
	dist := 2 // skip 2 fillInStackTrace frame
	for cls := this.Class; cls != nil; cls = cls.Super { //skip <init> stack frame
		dist = dist + 1
	}
	var traces []*Trace
	frames := f.Stack.Frames()
	for idx := len(frames) - 1 - dist; idx >= 0; idx = idx - 1 {
		m := frames[idx].GetMethod()
		traces = append(traces, &Trace{
			FileName: m.Cls.SourceFile,
			ClassName: m.Cls.Name,
			MethodName: m.Name,
			Line: 0,
		})
	}
	this.Extra = traces
}

func init()  {
	natives.Register("java/lang/Throwable", "fillInStackTrace", "(I)Ljava/lang/Throwable;", fillInStackTrace)
}