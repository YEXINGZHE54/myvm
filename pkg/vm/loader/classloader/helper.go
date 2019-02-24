package classloader

import (
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/reflect"
	"github.com/YEXINGZHE54/myvm/pkg/vm/loader/classfile"
)

func FileToClass(cf *classfile.ClassFile) (c *reflect.Class, err error) {
	c = new(reflect.Class)
	c.Flag = uint16(cf.AccessFlags)
	c.Name = cf.GetClass(cf.This)
	if cf.Super != 0x00 {
		c.SuperName = cf.GetClass(cf.Super)
	}
	for _, i := range cf.Ifaces{
		c.InterfaceNames = append(c.InterfaceNames, cf.GetClass(i))
	}
	// constants info
	var v interface{}
	var double bool
	for idx := 0; idx < len(cf.Constants); idx = idx + 1 {
		v = cf.Constants[idx]
		double = false
		switch val := v.(type) {
		case classfile.StringConst:
			v = cf.GetUTF8(classfile.ToIdx(val))
		case classfile.ClassConst:
			v = &reflect.ClsRef{cf.GetUTF8(classfile.ToIdx(val)), nil}
		case *classfile.FieldConst:
			v = &reflect.FieldRef{ getCNT(cf, val.Class, val.Nametype),nil}
		case *classfile.MethodConst:
			v = &reflect.MethodRef{ getCNT(cf, val.Class, val.Nametype),nil}
		case *classfile.IfaceMethodConst:
			v = &reflect.MethodRef{ getCNT(cf, val.Class, val.Nametype),nil}
		case classfile.LongConst, classfile.DoubleConst:
			idx = idx + 1 //skip
			double = true
		}
		c.Consts = append(c.Consts, v)
		//influnced by long/double, java ref index will increased, we must adapt to it
		if double {
			c.Consts = append(c.Consts, nil)
		}
	}
	// member info
	c.Fields = NewFields(c, cf)
	c.Methods = NewMethods(c, cf)
	// attr
	c.SourceFile = "Unknown"
	for _, attr := range cf.Attributes {
		if attr.Name == "SourceFile" {
			c.SourceFile = cf.GetUTF8(classfile.ToIdx(attr.Data))
		}
	}
	return
}

func copyMember(tm *reflect.Member, fm *classfile.Member, cls *reflect.Class, cf *classfile.ClassFile) {
	tm.Flag = uint16(fm.AccessFlags)
	tm.Name = cf.GetUTF8(fm.NameIndex)
	tm.Desc = cf.GetUTF8(fm.DescIndex)
	tm.Cls = cls
}

func NewFields(cls *reflect.Class, cf *classfile.ClassFile) (result []*reflect.Field) {
	for _, fi := range cf.Fields {
		field := new(reflect.Field)
		copyMember(&(field.Member), &fi, cls, cf)
		// find constval index for static final field
		for _, attr := range fi.Attributes {
			if attr.Name == "ConstantValue" {
				field.ConstValIndex = int(classfile.ToIdx(attr.Data))
			}
		}
		result = append(result, field)
	}
	return
}

func NewMethods(cls *reflect.Class, cf *classfile.ClassFile) (result []*reflect.Method) {
	for _, fi := range cf.Methods {
		method := new(reflect.Method)
		copyMember(&(method.Member), &fi, cls, cf)
		md, err := method.ParseSignature();
		if err != nil {
			panic(err)
		}
		if method.IsNative() {
			// no need to search code, because empty
			method.MaxStack = 4
			method.MaxLocal = method.ArgSlot
			switch md.Return[0] {
			case 'V': // return
				method.Codes = []byte{0xfe, 0xb1}
			case 'D': // dreturn
				method.Codes = []byte{0xfe, 0xaf}
			case 'F': // freturn
				method.Codes = []byte{0xfe, 0xae}
			case 'J': // lreturn
				method.Codes = []byte{0xfe, 0xad}
			case 'L','[': // areturn
				method.Codes = []byte{0xfe, 0xb0}
			default: //ireturn
				method.Codes = []byte{0xfe, 0xac}
			}
		} else {
			for _, attr := range fi.Attributes {
				if attr.Name == "Code" {
					code := attr.Data.(*classfile.Code)
					method.MaxStack = int(code.MaxStacks)
					method.MaxLocal = int(code.MaxLocals)
					method.Codes = code.Codes
					method.ExceptionTable = make([]*reflect.ExceptionHandle, 0)
					for _, excpt := range code.Exceptions {
						var ref *reflect.ClsRef = nil
						if excpt.CatchType > 0 {
							ref = cls.Consts[excpt.CatchType].(*reflect.ClsRef)
						}
						method.ExceptionTable = append(method.ExceptionTable, &reflect.ExceptionHandle{
							uint16(excpt.StartPC),
							uint16(excpt.EndPC),
							ref,
							uint16(excpt.HandlePC),
						})
					}
					for _, a := range code.Attributes {
						if a.Name == "LineNumberTable" {
							linetable := a.Data.([]classfile.LineNumber)
							method.LineTable = make([]reflect.PCLine, len(linetable))
							for idx, lt := range linetable {
								method.LineTable[idx].PC = int(lt.PC)
								method.LineTable[idx].Number = int(lt.Line)
							}
						}
					}
					break
				}
			}
		}
		result = append(result, method)
	}
	return
}

func getCNT(cf *classfile.ClassFile, cls, nt interface{}) reflect.MemberRef {
	clsname := cf.GetClass(classfile.ToIdx(cls))
	name, desc := cf.GetNameType(classfile.ToIdx(nt))
	return reflect.MemberRef{
		clsname, name, desc,
	}
}

func convertInt32(i interface{}) (v int32) {
	switch val := i.(type) {
	case classfile.IntegerConst:
		v = int32(val)
	default:
		v = i.(int32)
	}
	return
}

func convertFloat32(i interface{}) (v float32) {
	switch val := i.(type) {
	case classfile.FloatConst:
		v = float32(val)
	default:
		v = i.(float32)
	}
	return
}

func convertFloat64(i interface{}) (v float64) {
	switch val := i.(type) {
	case classfile.DoubleConst:
		v = float64(val)
	default:
		v = i.(float64)
	}
	return
}

func convertInt64(i interface{}) (v int64) {
	switch val := i.(type) {
	case classfile.LongConst:
		v = int64(val)
	default:
		v = i.(int64)
	}
	return
}

func setupLoader(ccls *reflect.Class, clsobj *reflect.Object, l *loader) {
	field, err := ccls.LookupInstanceField("classLoader", "Ljava/lang/ClassLoader;")
	if err != nil {
		panic(err)
	}
	clsobj.SetField(field, l.jObj)
}