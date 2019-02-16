package classloader

import (
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/reflect"
)

func init_statics(c *reflect.Class) error {
	// default static field already set to zero/false
	// we only need to set final const value
	for _, field := range c.Fields {
		if !field.IsStatic() || !field.IsFinal() {
			continue // skip non-static or non-final field
		}
		if field.ConstValIndex > 0 { // if const val
			switch field.Desc {
			case "Z", "B", "C", "S", "I":
				c.StaticVars.SetVal(field.SlotId, convertInt32(c.Consts[field.ConstValIndex]))
			case "J":
				c.StaticVars.SetLong(field.SlotId, convertInt64(c.Consts[field.ConstValIndex]))
			case "F":
				c.StaticVars.SetFloat(field.SlotId, convertFloat32(c.Consts[field.ConstValIndex]))
			case "D":
				c.StaticVars.SetDouble(field.SlotId, convertFloat64(c.Consts[field.ConstValIndex]))
			case "Ljava/lang/String;":
				o, err := c.Loader.JString(c.Consts[field.ConstValIndex].(string))
				if err != nil {
					return err
				}
				c.StaticVars.SetRef(field.SlotId, o)
			}
		}
	}
	return nil
}
