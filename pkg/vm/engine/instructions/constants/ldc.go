package constants

import (
	"github.com/YEXINGZHE54/myvm/pkg/utils"
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/instructions"
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/reflect"
	"github.com/YEXINGZHE54/myvm/pkg/vm/loader/classfile"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
)

const (
	ldc_op = 0x12
	ldcw_op = 0x13
	ldc2w_op = 0x14
)

type (
	LdcInst struct {
		idx uint8
	}
	LdcwInst struct {
		idx uint16
	}
	Ldc2wInst struct {
		idx uint16
	}
)

func (i *LdcInst) Clone() instructions.Inst {
	return &LdcInst{}
}

func (i *LdcInst) Fetch(coder *instructions.CodeReader) {
	i.idx = coder.Read1()
}

func (i *LdcInst) Exec(f *stack.Frame) {
	utils.Log("executing instruction ldc")
	ldc(f, int(i.idx))
}

func (i *LdcwInst) Clone() instructions.Inst {
	return &LdcwInst{}
}

func (i *LdcwInst) Fetch(coder *instructions.CodeReader) {
	i.idx = coder.Read2()
}

func (i *LdcwInst) Exec(f *stack.Frame) {
	utils.Log("executing instruction ldcw")
	ldc(f, int(i.idx))
}

func (i *Ldc2wInst) Clone() instructions.Inst {
	return &Ldc2wInst{}
}

func (i *Ldc2wInst) Fetch(coder *instructions.CodeReader) {
	i.idx = coder.Read2()
}

func (i *Ldc2wInst) Exec(f *stack.Frame) {
	utils.Log("executing instruction ldc2w")
	ldc2(f, int(i.idx))
}

func ldc(f *stack.Frame, idx int) {
	cls := f.GetMethod().Cls
	switch val := cls.Consts[idx].(type) {
	case classfile.IntegerConst:
		f.PushOpstackVal(int32(val))
	case classfile.FloatConst:
		f.PushOpstackFloat(float32(val))
	case string:
		o, err := f.GetMethod().Cls.Loader.JString(val)
		if err != nil {
			panic(err)
		}
		f.PushOpstackRef(o)
	case *reflect.ClsRef:
		var err error
		c := val.Ref
		if c == nil {
			c, err = cls.Loader.LoadClass(val.Name)
			if err != nil {
				panic(err)
			}
			val.Ref = c
		}
		f.PushOpstackRef(c.ToObject())
	default:
		panic("unsupported ldc")
	}
}

func ldc2(f *stack.Frame, idx int)  {
	cls := f.GetMethod().Cls
	switch val := cls.Consts[idx].(type) {
	case classfile.LongConst:
		f.PushOpstackLong(int64(val))
	case classfile.DoubleConst:
		f.PushOpstackDouble(float64(val))
	default:
		panic("unsupported ldc")
	}
}

func init() {
	instructions.Register(ldc_op, &LdcInst{})
	instructions.Register(ldcw_op, &LdcwInst{})
	instructions.Register(ldc2w_op, &Ldc2wInst{})
}