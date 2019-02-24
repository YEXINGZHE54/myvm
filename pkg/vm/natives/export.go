package natives

import (
	"errors"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
)

type (
	NativeMethod func(f *stack.Frame)
)

var (
	ErrorNotFound = errors.New("native method not found")
	nativeMethods = make(map[string]NativeMethod)
)

func LookUpNative(cls, name, desc string) (m NativeMethod, err error) {
	m, ok := nativeMethods[key(cls, name, desc)]
	if ok {
		return
	}
	if name == "registerNatives" && desc == "()V" {
		m = empty
		return 
	}
	if name == "initIDs" && desc == "()V" {
		m = empty
		return
	}
	err = ErrorNotFound
	return
}

func Register(cls, name, desc string, m NativeMethod) {
	nativeMethods[key(cls, name, desc)] = m
}

func key(cls, name, desc string) string {
	return cls + "~" + name + "~" + desc
}

func empty(f *stack.Frame)  {

}