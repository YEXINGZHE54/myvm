package class

import (
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/reflect"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
	"github.com/YEXINGZHE54/myvm/pkg/vm/natives"
)

func getDeclaredFields0(f *stack.Frame)  {
	this := f.GetLocalRef(0).Extra.(*reflect.Class)
	public := f.GetLocalVal(1) > 0
	arrcls, err := f.GetMethod().Cls.Loader.LoadClass("[Ljava/lang/reflect/Field;")
	if err != nil {
		panic(err)
	}
	fldcls, err := arrcls.ComponentClass()
	if err != nil {
		panic(err)
	}
	var fields []*reflect.Object
	for _, field := range this.Fields {
		if public && !field.IsPublic() {
			continue
		}
		o, err := fldcls.NewObject()
		if err != nil {
			panic(err)
		}
		fillFiled(field, o)
		fields = append(fields, o)
	}
	arr, err := arrcls.NewArray(len(fields))
	if err != nil {
		panic(err)
	}
	objs := arr.Refs()
	copy(objs, fields)
	f.PushOpstackRef(arr)
}

func fillFiled(field *reflect.Field, o *reflect.Object)  {
	loader := field.Cls.Loader
	for _, f := range o.Class.Fields {
		switch f.Name {
		case "clazz":
			o.SetField(f, field.Cls.ToObject())
		case "slot":
			o.SetField(f, int32(field.SlotId))
		case "name":
			name, err := loader.JString(field.Name)
			if err != nil {
				panic(err)
			}
			o.SetField(f, name)
		case "type":
			clsname, err := reflect.ConvertDescToClassName(field.Desc)
			if err != nil {
				panic(err)
			}
			tcls, err := loader.LoadClass(clsname)
			if err != nil {
				panic(err)
			}
			o.SetField(f, tcls.ToObject())
		case "modifiers":
			o.SetField(f, int32(field.Flag))
		case "signature":
			sig, err := loader.JString("")
			if err != nil {
				panic(err)
			}
			o.SetField(f, sig)
		case "annotations":
			arrcls, err := loader.LoadClass("[B")
			if err != nil {
				panic(err)
			}
			anno, err := arrcls.NewArray(0)
			if err != nil {
				panic(err)
			}
			o.SetField(f, anno)
		}
	}
}

func init()  {
	natives.Register("java/lang/Class","getDeclaredFields0", "(Z)[Ljava/lang/reflect/Field;", getDeclaredFields0)
}