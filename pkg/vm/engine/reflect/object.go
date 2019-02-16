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

// create one dimension array object
func (c *Class) NewArray(count int) (o *Object, err error) {
	return c.newMultiArray(c.Name, []int{count})
}

// create multi-dimension array object
func (c *Class) NewMArray(counts []int) (o *Object, err error) {
	return c.newMultiArray(c.Name, counts)
}

func (c *Class) newMultiArray(name string, counts []int) (o *Object, err error) {
	if len(name) < 2 || name[0] != '[' {
		err = ErrorInvalidArrayClassName
		return
	}
	if len(counts) == 0 {
		err = ErrorInvalidArrayCount
		return
	}
	cls, err := c.Loader.LoadClass(name)
	if err != nil {
		return nil, err
	}
	switch name[1] {
	case 'Z', 'B':
		o = &Object{make([]int8, counts[0]), cls}
	case 'C':
		o = &Object{make([]uint16, counts[0]), cls}
	case 'S':
		o = &Object{make([]int16, counts[0]), cls}
	case 'I':
		o = &Object{make([]int32, counts[0]), cls}
	case 'J':
		o = &Object{make([]int64, counts[0]), cls}
	case 'F':
		o = &Object{make([]float32, counts[0]), cls}
	case 'D':
		o = &Object{make([]float64, counts[0]), cls}
	case 'L': //refs
		o = &Object{make([]*Object, counts[0]), cls}
	case '[': //multi array
		refs := make([]*Object, counts[0])
		for idx := range refs {
			refs[idx], err = c.newMultiArray(name[1:], counts[1:])
			if err != nil {
				return
			}
		}
		o = &Object{refs, cls}
	default:
		err = ErrorInvalidArrayClassName
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