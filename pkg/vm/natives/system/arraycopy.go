package system

import (
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/reflect"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
	"github.com/YEXINGZHE54/myvm/pkg/vm/natives"
)

func arraycopy(f *stack.Frame)  {
	src := f.GetLocalRef(0)
	srcPos := f.GetLocalVal(1)
	dest := f.GetLocalRef(2)
	destPos := f.GetLocalVal(3)
	length := f.GetLocalVal(4)
	if src == nil || dest == nil {
		panic("NullPointerException")
	}
	if srcPos < 0 || destPos < 0 || length < 0 {
		panic("invalid argument of arraycopy")
	}
	if !src.Class.IsClass() || !dest.Class.IsArray() {
		panic("array copy must be applied to array")
	}
	srcComp, err := src.Class.ComponentClass()
	if err != nil {
		panic(err)
	}
	destComp, err := src.Class.ComponentClass()
	if err != nil {
		panic(err)
	}
	if srcComp.IsPrimitive() && destComp.IsPrimitive() && src.Class != dest.Class {
		panic("array type not match")
	}
	reflect.Copy(src, dest, srcPos, destPos, length)
}

func init()  {
	natives.Register("java/lang/System", "arraycopy", "(Ljava/lang/Object;ILjava/lang/Object;II)V", arraycopy)
}