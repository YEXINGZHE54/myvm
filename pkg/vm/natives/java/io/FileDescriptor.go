package io

import (
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
	"github.com/YEXINGZHE54/myvm/pkg/vm/natives"
)

func initIDs(f *stack.Frame)  {

}

func init()  {
	natives.Register("java/io/FileDescriptor", "initIDs", "()V", initIDs)
}
