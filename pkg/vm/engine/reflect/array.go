package reflect

// create one dimension array object
func (c *Class) NewArray(count int) (o *Object, err error) {
	return c.newMultiArray(c.Name, []int{count})
}

// create multi-dimension array object
func (c *Class) NewMArray(counts []int) (o *Object, err error) {
	return c.newMultiArray(c.Name, counts)
}

func (c *Class) ArrayFrom(val interface{}) (o *Object, err error) {
	return &Object{val, c, nil}, nil
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
		o = &Object{make([]int8, counts[0]), cls, nil}
	case 'C':
		o = &Object{make([]uint16, counts[0]), cls, nil}
	case 'S':
		o = &Object{make([]int16, counts[0]), cls, nil}
	case 'I':
		o = &Object{make([]int32, counts[0]), cls, nil}
	case 'J':
		o = &Object{make([]int64, counts[0]), cls, nil}
	case 'F':
		o = &Object{make([]float32, counts[0]), cls, nil}
	case 'D':
		o = &Object{make([]float64, counts[0]), cls, nil}
	case 'L': //refs
		o = &Object{make([]*Object, counts[0]), cls, nil}
	case '[': //multi array
		refs := make([]*Object, counts[0])
		for idx := range refs {
			refs[idx], err = c.newMultiArray(name[1:], counts[1:])
			if err != nil {
				return
			}
		}
		o = &Object{refs, cls, nil}
	default:
		err = ErrorInvalidArrayClassName
	}
	return
}

func (o *Object) Bytes() []int8 {
	return o.fields.([]int8)
}

func (o *Object) Shorts() []int16 {
	return o.fields.([]int16)
}

func (o *Object) Ints() []int32 {
	return o.fields.([]int32)
}

func (o *Object) Longs() []int64 {
	return o.fields.([]int64)
}

func (o *Object) Chars() []uint16 {
	return o.fields.([]uint16)
}

func (o *Object) Floats() []float32 {
	return o.fields.([]float32)
}

func (o *Object) Doubles() []float64 {
	return o.fields.([]float64)
}

func (o *Object) Refs() []*Object {
	return o.fields.([]*Object)
}

func (o *Object) ArrLength() int {
	switch val := o.fields.(type) {
	case []int8:
		return len(val)
	case []int16:
		return len(val)
	case []int32:
		return len(val)
	case []int64:
		return len(val)
	case []uint8:
		return len(val)
	case []float32:
		return len(val)
	case []float64:
		return len(val)
	case []*Object:
		return len(val)
	default:
		panic(val)
	}
}
