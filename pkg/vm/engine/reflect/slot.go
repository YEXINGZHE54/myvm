package reflect

func (s Slots) SetVal(idx int, val int32) {
	s[idx] = val
}

func (s Slots) GetVal(idx int) int32 {
	return GetVal(s[idx])
}

func (s Slots) SetLong(idx int, val int64) {
	s[idx] = val
}

func (s Slots) GetLong(idx int) int64 {
	return GetLong(s[idx])
}

func (s Slots) SetFloat(idx int, val float32) {
	s[idx] = val
}

func (s Slots) GetFloat(idx int) float32 {
	return GetFloat(s[idx])
}

func (s Slots) SetDouble(idx int, val float64) {
	s[idx] = val
}

func (s Slots) GetDouble(idx int) float64 {
	return GetDouble(s[idx])
}

func (s Slots) SetRef(idx int, ref *Object) {
	s[idx] = ref
}

func (s Slots) GetRef(idx int) *Object {
	return GetRef(s[idx])
}