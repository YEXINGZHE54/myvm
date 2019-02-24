package reflect

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrorFieldNotFound = errors.New("field not found")
	ErrorInvalidClassDesc = errors.New("invalid class descriptor")
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
	cname, err := ConvertDescToClassName(name[1:])
	if err != nil {
		return
	}
	cls, err = c.Loader.LoadClass(cname)
	return
}

func ConvertDescToClassName(desc string) (cname string, err error) {
	switch desc[0] {
	case 'Z':
		cname = "boolean"
	case 'B':
		cname = "byte"
	case 'C':
		cname = "char"
	case 'S':
		cname = "short"
	case 'I':
		cname = "int"
	case 'J':
		cname = "long"
	case 'F':
		cname = "float"
	case 'D':
		cname = "double"
	case 'L': //refs
		cname = desc[1:len(desc)-1]
	case '[': //array
		cname = desc
	default:
		err = ErrorInvalidClassDesc
	}
	return
}

func (c *Class) ExtendsOrSame(other *Class) bool {
	for cls := c; cls != nil; cls = cls.Super {
		if cls == other {
			return true
		}
	}
	return false
}

func (c *Class) Implements(iface *Class) bool {
	for cls := c; cls != nil; cls = cls.Super {
		for _, ifc := range cls.Interfaces {
			for iplted := ifc; iplted != nil; iplted = iplted.Super {
				if iplted == iface {
					return true
				}
			}
		}
	}
	return false
}

func (c *Class) LookupInstanceField(name, desc string) (f *Field, err error) {
	for cls := c; cls != nil; cls = cls.Super {
		for _, f = range cls.Fields {
			if f.Name == name && f.Desc == desc && !f.IsStatic() {
				return
			}
		}
	}
	err = ErrorFieldNotFound
	return
}

func (c *Class) LookupStaticField(name, desc string) (f *Field, err error) {
	for _, f = range c.Fields {
		if f.Name == name && f.Desc == desc && f.IsStatic() {
			return
		}
	}
	err = ErrorFieldNotFound
	return
}

func (c *Class) GetField(f *Field) interface{} {
	return c.StaticVars[f.SlotId]
}

/* JVM SE8, checkcast
if S is non-array class, then:
1. if T is class, then S is the same to T, or extends T
2. if T is interface, then S must implement T
if S is interface, then:
1. T == java/lang/Object
2. if T is interface, then S is the same to T or extends T
(this case happens when array of interface item, like: []List []Collection ...)
if S is array class, then
1. T == java/lang/Object
2. if T is interface, T must be the interfaces implemented by arrays
3. if T is array class, then:
-- SC and TC are the same primitive class, or
-- SC and tC are reference and call recursive on checkcast(SC, TC)
 */
func CanCastTo(S, T *Class) bool {
	if S.IsClass() && !S.IsArray() {
		// 1
		if T.IsClass() {
			return S.ExtendsOrSame(T)
		}
		// 2
		return S.Implements(T)
	} else if S.IsInterface() {
		// 1
		if T.IsClass() {
			return T.Name == "java/lang/Object"
		}
		// 2
		return S.ExtendsOrSame(T)
	} else { // S is array class
		// 1
		if T.IsClass() && !T.IsArray() {
			return T.Name == "java/lang/Object"
		}
		// 2
		if T.IsInterface() {
			return S.Implements(T)
		}
		// 3
		sc, err := S.ComponentClass()
		if err != nil {
			panic(err)
		}
		tc, err := T.ComponentClass()
		if err != nil {
			panic(err)
		}
		return CanCastTo(sc, tc)
	}
	return false
}

// only accept: int32, int64, float32, float64, *Object
func (c *Class) SetField(f *Field, val interface{}) {
	switch v := val.(type) {
	case int32:
		c.StaticVars.SetVal(f.SlotId, v)
	case float32:
		c.StaticVars.SetFloat(f.SlotId, v)
	case int64:
		c.StaticVars.SetLong(f.SlotId, v)
	case float64:
		c.StaticVars.SetDouble(f.SlotId, v)
	case *Object:
		c.StaticVars.SetRef(f.SlotId, v)
	default:
		panic(fmt.Sprintf("unexpected type %T", val))
	}
}