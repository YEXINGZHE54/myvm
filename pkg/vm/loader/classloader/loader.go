package classloader

import (
	"strings"
	"github.com/YEXINGZHE54/myvm/pkg/vm/loader/classfile"
	"github.com/YEXINGZHE54/myvm/pkg/vm/loader/classpath"
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/reflect"
)

type (
	loader struct {
		classes map[string]*reflect.Class
		cp *classpath.ClassPath
	}
)

func NewLoader(bootPath, classPath string) (l reflect.Loader) {
	l = &loader{make(map[string]*reflect.Class),classpath.ParseOption(bootPath, classPath)}
	return
}

func (l *loader) LoadClass(cls string) (c *reflect.Class, err error) {
	// try cache
	c, ok := l.classes[cls]
	if ok {
		return
	}
	// read classfile
	clsname := strings.Replace(cls, ".", "/", -1)
	data, _, err := l.cp.ReadClass(clsname)
	if err != nil {
		return
	}
	cf, err := classfile.Parse(data)
	if err != nil {
		return
	}
	// verify, skip it now
	// convert to reflect.class
	c, err = FileToClass(cf)
	if err != nil {
		return
	}
	c.Loader = l
	// resolve super and interfaces
	if len(c.SuperName) > 0 {
		c.Super, err = l.LoadClass(c.SuperName)
		if err != nil {
			return
		}
	}
	for _, iface := range c.InterfaceNames {
		ifc, err := l.LoadClass(iface)
		if err != nil {
			return nil, err
		}
		c.Interfaces = append(c.Interfaces, ifc)
	}
	// prepare
	
	// initialzie

	// finnaly record it
	l.classes[cls] = c
	return
}

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
	// member info
	c.Fields = NewFields(c, cf)
	c.Methods = NewMethods(c, cf)
	// constants info
	var v interface{}
	for idx := 1; idx < len(cf.Constants); idx = idx + 1 {
		v = cf.Constants[idx]
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
		}
		c.Consts = append(c.Consts, v)
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
		result = append(result, field)
	}
	return
}

func NewMethods(cls *reflect.Class, cf *classfile.ClassFile) (result []*reflect.Method) {
	for _, fi := range cf.Methods {
		method := new(reflect.Method)
		copyMember(&(method.Member), &fi, cls, cf)
		code := fi.GetCode()
		if code != nil {
			method.MaxStack = int(code.MaxStacks)
			method.MaxLocal = int(code.MaxLocals)
			method.Codes = code.Codes
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