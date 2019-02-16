package reflect

import "errors"

var (
	ErrorInvalidArrayClassName = errors.New("invalid array class name")
	ErrorInvalidArrayCount = errors.New("invalid array count")
)

// create object
func (c *Class) NewObject() (o *Object, err error) {
	// create instance of class
	o = &Object{
		Class:c,
		fields:make(Slots, c.InstanceSlotCount),
	}
	return
}

//TODO
// convert to class object
func (c *Class) ToObject() (o *Object, err error) {
	return
}

func (o *Object) Fields() Slots {
	return o.fields.(Slots)
}

// only accept: int32, int64, float32, float64, *Object
func (o *Object) SetField(f *Field, val interface{}) {
	switch v := val.(type) {
	case int32:
		o.Fields().SetVal(f.SlotId, v)
	case float32:
		o.Fields().SetFloat(f.SlotId, v)
	case int64:
		o.Fields().SetLong(f.SlotId, v)
	case float64:
		o.Fields().SetDouble(f.SlotId, v)
	case *Object:
		o.Fields().SetRef(f.SlotId, v)
	}
}

// only accept: int32, int64, float32, float64, *Object
func (o *Object) GetField(f *Field) (v interface{}) {
	slot := o.Fields()[f.SlotId]
	switch f.Desc[0] {
	case 'L','[':
		v = slot.Ref
	default:
		v = slot.Val
	}
	return
}