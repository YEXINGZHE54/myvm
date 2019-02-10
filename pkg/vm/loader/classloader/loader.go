package classloader

import (
	"strings"
	"myvm/pkg/vm/loader/classfile"
	"myvm/pkg/vm/loader/classpath"
	"myvm/pkg/vm/engine/reflect"
)

type (
	loader struct {
		cp *classpath.ClassPath
	}
)

func NewLoader(bootPath, classPath string) (l reflect.Loader) {
	l = &loader{classpath.ParseOption(bootPath, classPath)}
	return
}

func (l *loader) LoadClass(cls string) (c *reflect.Class, err error) {
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

	// prepare
	
	// initialzie

	return
}

func FileToClass(cf *classfile.ClassFile) (c *reflect.Class, err error) {
	c = new(reflect.Class)
	c.Flag = uint16(cf.AccessFlags)
	c.Name = cf.GetClass(cf.This)
	c.SuperName = cf.GetClass(cf.Super)
	for _, i := range cf.Ifaces{
		c.InterfaceNames = append(c.InterfaceNames, cf.GetClass(i))
	}
	// member info
	c.Fields = NewFields(c, cf)
	c.Methods = NewMethods(c, cf)
	// constants info
	
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
		method.MaxStack = int(code.MaxStacks)
		method.MaxLocal = int(code.MaxLocals)
		method.Codes = code.Codes
		result = append(result, method)
	}
	return
}