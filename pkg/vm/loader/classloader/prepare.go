package classloader

import "github.com/YEXINGZHE54/myvm/pkg/vm/engine/reflect"

func prepare(c *reflect.Class)  {
	// calc instance field slots
	slotId := 0
	if c.Super != nil && c.Super.InstanceSlotCount > 0 {
		slotId = c.Super.InstanceSlotCount
	}
	for _, field := range c.Fields {
		if field.IsStatic() {
			continue //skip static vars
		}
		field.SlotId = slotId
		slotId = slotId + 1
		if field.IsLongDouble() {
			slotId = slotId + 1 // long/double need 2 slot
		}
	}
	c.InstanceSlotCount = slotId
	// calc static field slots
	slotId = 0
	for _, field := range c.Fields {
		if field.IsStatic() {
			//only static vars
			field.SlotId = slotId
			slotId = slotId + 1
			if field.IsLongDouble() {
				slotId = slotId + 1 // long/double need 2 slot
			}
		}
	}
	c.StaticVars = make(reflect.Slots, slotId)
}
