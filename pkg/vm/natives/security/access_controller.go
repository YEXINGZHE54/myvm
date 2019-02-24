package security

import (
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
	"github.com/YEXINGZHE54/myvm/pkg/vm/natives"
)

func doPrivileged(f *stack.Frame) {
	action := f.GetLocalRef(0)
	method, err := action.Class.LookupMethod("run", "()Ljava/lang/Object;")
	if err != nil {
		panic(err)
	}
	newf := stack.NewFrame(method)
	f.Stack.Push(newf)
	newf.SetLocalRef(0, action)
}

func init()  {
	natives.Register("java/security/AccessController", "doPrivileged", "(Ljava/security/PrivilegedAction;)Ljava/lang/Object;", doPrivileged)
	natives.Register("java/security/AccessController", "doPrivileged", "(Ljava/security/PrivilegedExceptionAction;)Ljava/lang/Object;", doPrivileged)
}