package objects

import (
	"github.com/YEXINGZHE54/myvm/pkg/utils"
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/instructions"
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/reflect"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
)

const (
	new_op = 0xbb
	getfield_op = 0xb4
	putfield_op = 0xb5
)

type (
	NewInst struct {
		idx uint16
	}
	PutInst struct {
		idx uint16
	}
	GetInst struct {
		idx uint16
	}
)

func (i *NewInst) Clone() instructions.Inst {
	return &NewInst{}
}

func (i *NewInst) Fetch(coder *instructions.CodeReader) {
	i.idx = coder.Read2()
}

func (i *NewInst) Exec(f *stack.Frame) {
	cls := f.GetMethod().Cls
	ref := cls.Consts[i.idx].(*reflect.ClsRef)
	utils.Log("executing instruction new, %s", ref.Name)
	err := cls.Loader.ResolveClass(ref)
	if err != nil {
		panic(err)
	}
	// check class init
	inited, err := init_class(f, cls)
	if err != nil {
		panic(err)
	}
	if !inited {
		revertPC(f)
		return
	}
	o, err := ref.Ref.NewObject()
	if err != nil {
		panic(err)
	}
	f.PushOpstackRef(o)
}

func (i *GetInst) Clone() instructions.Inst {
	return &GetInst{}
}

func (i *GetInst) Fetch(coder *instructions.CodeReader) {
	i.idx = coder.Read2()
}

func (i *GetInst) Exec(f *stack.Frame) {
	cls := f.GetMethod().Cls
	ref := cls.Consts[i.idx].(*reflect.FieldRef)
	utils.Log("executing instruction getfield, %s.%s%s", ref.ClsName, ref.Name, ref.Desc)
	err := cls.Loader.ResolveField(ref)
	if err != nil {
		panic(err)
	}
	obj := f.PopOpstackRef()
	switch ref.Ref.Desc[0] {
	case 'Z','B','C','S','I':
		f.PushOpstackVal(obj.Fields().GetVal(ref.Ref.SlotId))
	case 'F':
		f.PushOpstackFloat(obj.Fields().GetFloat(ref.Ref.SlotId))
	case 'J':
		f.PushOpstackLong(obj.Fields().GetLong(ref.Ref.SlotId))
	case 'D':
		f.PushOpstackDouble(obj.Fields().GetDouble(ref.Ref.SlotId))
	case 'L','[': //object or array
		f.PushOpstackRef(obj.Fields().GetRef(ref.Ref.SlotId))
	}
}

func (i *PutInst) Clone() instructions.Inst {
	return &PutInst{}
}

func (i *PutInst) Fetch(coder *instructions.CodeReader) {
	i.idx = coder.Read2()
}

func (i *PutInst) Exec(f *stack.Frame) {
	cls := f.GetMethod().Cls
	ref := cls.Consts[i.idx].(*reflect.FieldRef)
	utils.Log("executing instruction putfield, %s.%s%s", ref.ClsName, ref.Name, ref.Desc)
	err := cls.Loader.ResolveField(ref)
	if err != nil {
		panic(err)
	}
	switch ref.Ref.Desc[0] {
	case 'Z','B','C','S','I':
		v := f.PopOpstackVal()
		f.PopOpstackRef().Fields().SetVal(ref.Ref.SlotId, v)
	case 'F':
		v := f.PopOpstackFloat()
		f.PopOpstackRef().Fields().SetFloat(ref.Ref.SlotId, v)
	case 'J':
		v := f.PopOpstackLong()
		f.PopOpstackRef().Fields().SetLong(ref.Ref.SlotId, v)
	case 'D':
		v := f.PopOpstackDouble()
		f.PopOpstackRef().Fields().SetDouble(ref.Ref.SlotId, v)
	case 'L','[': //object or array
		v := f.PopOpstackRef()
		f.PopOpstackRef().Fields().SetRef(ref.Ref.SlotId, v)
	}
}

func init() {
	instructions.Register(new_op, &NewInst{})
	instructions.Register(putfield_op, &PutInst{})
	instructions.Register(getfield_op, &GetInst{})
}
