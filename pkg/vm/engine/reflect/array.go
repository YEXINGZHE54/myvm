package reflect


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
