package reflect

import (
	"errors"
	"strings"
)

var (
	ErrorFieldNotFound = errors.New("field not found")
)

func (c *Class) GetMain() (m *Method, err error) {
	return c.GetStatic("main", "([Ljava/lang/String;)V")
}

func (c *Class) GetClinit() (m *Method, err error) {
	return c.GetStatic("<clinit>", "()V")
}

func (c *Class) IsSuperOf(other *Class) bool {
	for super := other.Super; super != nil; super = super.Super {
		if c == super {
			return true
		}
	}
	return false
}

func (c *Class) GetPackageName() string {
	idx := strings.LastIndex(c.Name, "/")
	if idx < 0 {
		return ""
	}
	return c.Name[:idx]
}

// return class whose component is current class
func (c *Class) ArrayClass() (cls *Class, err error) {
	cname := "[L" + c.Name + ";"
	return c.Loader.LoadClass(cname)
}

// return component class if current class is array class
func (c *Class) ComponentClass() (cls *Class, err error) {
	name := c.Name
	if name[0] != '[' {
		err = ErrorInvalidArrayClassName
		return
	}
	var cname string
	switch name[1] {
	case 'Z':
		cname = "java/lang/Boolean"
	case 'B':
		cname = "java/lang/Byte"
	case 'C':
		cname = "java/lang/Character"
	case 'S':
		cname = "java/lang/Short"
	case 'I':
		cname = "java/lang/Integer"
	case 'J':
		cname = "java/lang/Long"
	case 'F':
		cname = "java/lang/Float"
	case 'D':
		cname = "java/lang/Double"
	case 'L': //refs
		cname = name[2:len(name)-1]
	case '[': //multi array
		cname = name[1:]
	default:
		err = ErrorInvalidArrayClassName
		return
	}
	cls, err = c.Loader.LoadClass(cname)
	return
}

func (c *Class) GetInstanceField(name, desc string) (f *Field, err error) {
	for _, f = range c.Fields {
		if f.Name == name && f.Desc == desc && !f.IsStatic() {
			return
		}
	}
	err = ErrorFieldNotFound
	return
}