package io

import "github.com/YEXINGZHE54/myvm/pkg/vm/natives"

func init()  {
	natives.Register("java/io/FileInputStream", "initIDs", "()V", initIDs)
}