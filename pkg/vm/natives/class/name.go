package class

import (
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/reflect"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
	"github.com/YEXINGZHE54/myvm/pkg/vm/natives"
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

func convert(name string) string {
	return strings.Replace(name, "/", ".", -1)
}

func init()  {
	natives.Register("java/lang/Class", "getName0", "()Ljava/lang/String;", getName0)
}