package load

import (
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/instructions"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
)

const (
	aload0_op = 0x2a
	aload1_op = 0x2b
	aload2_op = 0x2c
	aload3_op = 0x2d
)

type (
	ALoad0 struct{
	}
	ALoad1 struct {
	}
	ALoad2 struct {
	}
	ALoad3 struct {
	}
)

func (i *ALoad0) Clone() instructions.Inst {
	return i
}

func (i *ALoad0) Fetch(coder *instructions.CodeReader) {

}

func (i *ALoad0) Exec(f *stack.Frame) {
	aload(f, 0)
}


func (i *ALoad1) Clone() instructions.Inst {
	return i
}

func (i *ALoad1) Fetch(coder *instructions.CodeReader) {

}

func (i *ALoad1) Exec(f *stack.Frame) {
	aload(f, 1)
}


func (i *ALoad2) Clone() instructions.Inst {
	return i
}

func (i *ALoad2) Fetch(coder *instructions.CodeReader) {

}

func (i *ALoad2) Exec(f *stack.Frame) {
	aload(f, 2)
}


func (i *ALoad3) Clone() instructions.Inst {
	return i
}

func (i *ALoad3) Fetch(coder *instructions.CodeReader) {

}

func (i *ALoad3) Exec(f *stack.Frame) {
	aload(f, 3)
}

func aload(f *stack.Frame, idx int)  {
	f.PushOpstackRef(f.GetLocalRef(idx))
}

func init() {
	instructions.Register(aload0_op, &ALoad0{})
	instructions.Register(aload1_op, &ALoad1{})
	instructions.Register(aload2_op, &ALoad2{})
	instructions.Register(aload3_op, &ALoad3{})
}
