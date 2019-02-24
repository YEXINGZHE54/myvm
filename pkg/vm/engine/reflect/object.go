package reflect

import (
	"errors"
	"fmt"
	"github.com/YEXINGZHE54/myvm/pkg/utils"
)

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

// convert to class object
func (c *Class) ToObject() (o *Object) {
	return c.ClsObj
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
	default:
		panic(fmt.Sprintf("unexpected type %T", val))
	}
}

// only accept: int32, int64, float32, float64, *Object
func (o *Object) GetField(f *Field) (v interface{}) {
	return o.Fields()[f.SlotId]
}

func (o *Object) GoString() string {
	field, err := o.Class.LookupInstanceField("value", "[C")
	if err != nil {
		panic(err)
	}
	chars := o.GetField(field).(*Object).Chars()
	if err != nil {
		panic(err)
	}
	return utils.UTF16ToString(chars)
}

func (o *Object) Clone() *Object {
	var data interface{}
	switch v := o.fields.(type) {
	case []int8:
		arr := make([]int8, len(v))
		copy(arr, v)
		data = arr
	case []int16:
		arr := make([]int16, len(v))
		copy(arr, v)
		data = arr
	case []uint16:
		arr := make([]uint16, len(v))
		copy(arr, v)
		data = arr
	case []int32:
		arr := make([]int32, len(v))
		copy(arr, v)
		data = arr
	case []int64:
		arr := make([]int64, len(v))
		copy(arr, v)
		data = arr
	case []*Object:
		arr := make([]*Object, len(v))
		copy(arr, v)
		data = arr
	case Slots:
		arr := make(Slots, len(v))
		copy(arr, v)
		data = arr
	}
	return &Object{
		data,
		o.Class,
		o.Extra,
	}
}