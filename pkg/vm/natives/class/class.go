package class

import (
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/reflect"
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
	"github.com/YEXINGZHE54/myvm/pkg/vm/natives"
)

func getDeclaredConstructors0(f *stack.Frame)  {
	this := f.This().Extra.(*reflect.Class)
	public := f.GetLocalVal(1)
	arrcls, err := f.GetMethod().Cls.Loader.LoadClass("[Ljava/lang/reflect/Constructor;")
	if err != nil {
		panic(err)
	}
	compcls, err := arrcls.ComponentClass()
	if err != nil {
		panic(err)
	}
	var ctrs []*reflect.Object
	ctormethods := this.GetContructors(public > 0)
	for _, method := range ctormethods {
		obj, err := compcls.NewObject()
		if err != nil {
			panic(err)
		}
		fillContructor(method, obj)
		ctrs = append(ctrs, obj)
	}
	arr, err := arrcls.NewArray(len(ctrs))
	if err != nil {
		panic(err)
	}
	copy(arr.Refs(), ctrs)
	f.PushOpstackRef(arr)
}

func fillContructor(ctor *reflect.Method, o *reflect.Object)  {
	loader := ctor.Cls.Loader
	arrcls, err := loader.LoadClass("[Ljava/lang/Class;")
	if err != nil {
		panic(err)
	}
	for _, f := range o.Class.Fields {
		switch f.Name {
		case "clazz":
			o.SetField(f, ctor.Cls.ToObject())
		case "slot":
			o.SetField(f, int32(0))
		case "parameterTypes":
			var types []*reflect.Object
			md, err := ctor.ParseSignature()
			if err != nil {
				panic(err)
			}
			for _, arg := range md.Args {
				aname, err := reflect.ConvertDescToClassName(arg)
				if err != nil {
					panic(err)
				}
				acls, err := loader.LoadClass(aname)
				if err != nil {
					panic(err)
				}
				types = append(types, acls.ToObject())
			}
			arr, err := arrcls.NewArray(len(types))
			if err != nil {
				panic(err)
			}
			copy(arr.Refs(), types)
			o.SetField(f, arr)
		case "exceptionTypes":
			var types []*reflect.Object
			for _, exc := range ctor.ExceptionTable {
				if exc.Caught.Ref == nil {
					err := loader.ResolveClass(exc.Caught)
					if err != nil {
						panic(err)
					}
				}
				types = append(types, exc.Caught.Ref.ToObject())
			}
			arr, err := arrcls.NewArray(len(types))
			if err != nil {
				panic(err)
			}
			copy(arr.Refs(), types)
			o.SetField(f, arr)
		case "modifiers":
			o.SetField(f, int32(ctor.Flag))
		case "signature":
			sig, err := loader.JString(ctor.Desc)
			if err != nil {
				panic(err)
			}
			o.SetField(f, sig)
		}
	}
	o.Extra = ctor // save method; read carefully, because of Pattern of Copy Accessor
}

func getModifiers(f *stack.Frame)  {
	this := f.This().Extra.(*reflect.Class)
	f.PushOpstackVal(int32(this.Flag))
}

func init()  {
	natives.Register("java/lang/Class", "getDeclaredConstructors0", "(Z)[Ljava/lang/reflect/Constructor;", getDeclaredConstructors0)
	natives.Register("java/lang/Class", "getModifiers", "()I", getModifiers)
}