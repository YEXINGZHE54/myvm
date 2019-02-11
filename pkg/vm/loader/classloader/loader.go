package classloader

import (
	"errors"
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

var (
	ErrorFieldNotFound = errors.New("field not found")
	ErrorMethodNotFound = errors.New("method not found")
)

func NewLoader(bootPath, classPath string) (l reflect.Loader) {
	l = &loader{make(map[string]*reflect.Class),classpath.ParseOption(bootPath, classPath)}
	return
}

func (l *loader) readClass(cls string) (cf *classfile.ClassFile, err error) {
	clsname := strings.Replace(cls, ".", "/", -1)
	data, _, err := l.cp.ReadClass(clsname)
	if err != nil {
		return
	}
	cf, err = classfile.Parse(data)
	return
}

func (l *loader) resolveSuper(c *reflect.Class) (err error) {
	if len(c.SuperName) > 0 {
		c.Super, err = l.LoadClass(c.SuperName)
		if err != nil {
			return
		}
	}
	for _, iface := range c.InterfaceNames {
		ifc, err := l.LoadClass(iface)
		if err != nil {
			return err
		}
		c.Interfaces = append(c.Interfaces, ifc)
	}
	return
}

func (l *loader) ResolveClass(clsref *reflect.ClsRef) (err error) {
	if clsref.Ref == nil {
		clsref.Ref, err = l.LoadClass(clsref.Name)
	}
	return
}

func (l *loader) ResolveField(ref *reflect.FieldRef) (err error) {
	var cls *reflect.Class
	if ref.Ref == nil {
		cls, err = l.LoadClass(ref.ClsName)
		if err != nil {
			return err
		}
		for _, field := range cls.Fields {
			if field.Name == ref.Name && field.Desc == ref.Desc {
				ref.Ref = field
				return
			}
		}
	}
	return ErrorFieldNotFound
}

func (l *loader) ResolveMethod(ref *reflect.MethodRef) (err error) {
	var cls *reflect.Class
	if ref.Ref == nil {
		cls, err = l.LoadClass(ref.ClsName)
		if err != nil {
			return err
		}
		for _, field := range cls.Methods {
			if field.Name == ref.Name && field.Desc == ref.Desc {
				ref.Ref = field
				return
			}
		}
	}
	return ErrorMethodNotFound
}

func (l *loader) LoadClass(cls string) (c *reflect.Class, err error) {
	// try cache
	c, ok := l.classes[cls]
	if ok {
		return
	}
	// read classfile
	cf, err := l.readClass(cls)
	if err != nil {
		return
	}
	// verify
	err = verify(c)
	if err != nil {
		return
	}
	// convert to reflect.class
	c, err = FileToClass(cf)
	if err != nil {
		return
	}
	c.Loader = l
	// resolve super and interfaces
	err = l.resolveSuper(c)
	if err != nil {
		return
	}
	// prepare
	prepare(c)
	// initialzie
	init_statics(c)
	// finnaly record it
	l.classes[cls] = c
	return
}