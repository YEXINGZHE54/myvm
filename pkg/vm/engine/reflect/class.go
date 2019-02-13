package reflect

import (
	"strings"
)

func (c *Class) GetMain() (m *Method, err error) {
	return c.GetStatic("main", "([Ljava/lang/String;)V")
}

func (c *Class) GetClinit() (m *Method, err error) {
	return c.GetStatic("<clinit>", "()V")
}

func (c *Class) NewObject() (o *Object, err error) {
	// create instance of class
	o = &Object{
		Class:c,
		Fields:make(Slots, c.InstanceSlotCount),
	}
	return
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