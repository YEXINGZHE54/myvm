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
	if ref.Ref == nil {
		var cls *reflect.Class
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
		return ErrorFieldNotFound
	}
	return
}

func (l *loader) ResolveMethod(ref *reflect.MethodRef) (err error) {
	if ref.Ref == nil {
		var cls *reflect.Class
		cls, err = l.LoadClass(ref.ClsName)
		if err != nil {
			return err
		}
		ref.Ref, err = cls.LookupMethod(ref.Name, ref.Desc)
		if err != nil {
			return
		}
		err = ref.Ref.ParseSignature()
		if err != nil {
			return
		}
	}
	return
}

func (l *loader) ResolveIfaceMethod(ref *reflect.MethodRef) (err error) {
	if ref.Ref == nil {
		var cls *reflect.Class
		cls, err = l.LoadClass(ref.ClsName)
		if err != nil {
			return err
		}
		ref.Ref, err = cls.LookupIfaceMethod(ref.Name, ref.Desc)
		if err != nil {
			return
		}
		err = ref.Ref.ParseSignature()
		if err != nil {
			return
		}
	}
	return
}

func (l *loader) LoadClass(cls string) (c *reflect.Class, err error) {
	// try cache
	c, ok := l.classes[cls]
	if ok {
		return
	}
	if len(cls) > 1 && cls[0] == '[' { //array class
		c, err = l.loadArrayClass(cls)
	} else {
		c, err = l.loadObjectClass(cls)
	}
	if err != nil {
		return
	}
	// finnaly record it
	l.classes[cls] = c
	return
}

func (l *loader) loadObjectClass(cls string) (c *reflect.Class, err error) {
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
	return
}

func (l *loader) loadArrayClass(cls string) (c *reflect.Class, err error) {
	supername := "java/lang/Object"
	ifacenames := []string{"java/lang/Cloneable", "java/io/Serializable"}
	super, err := l.LoadClass(supername)
	if err != nil {
		return
	}
	var ifaces []*reflect.Class
	for _, ifn := range ifacenames {
		iface, err := l.LoadClass(ifn)
		if err != nil {
			return nil, err
		}
		ifaces = append(ifaces, iface)
	}
	c = &reflect.Class{
		Flag: reflect.ACCESS_PUBLIC,
		Name: cls,
		SuperName: supername,
		InterfaceNames: ifacenames,
		Loader: l,
		Super: super,
		Interfaces: ifaces,
		Started: true, //skip <clinit>
	}
	// load componet class
	_, err = c.ComponentClass()
	return
}