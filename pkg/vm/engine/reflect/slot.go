package reflect

func (s Slots) SetVal(idx int, val int32) {
	s[idx].Val = val
}

func (s Slots) GetVal(idx int) int32 {
	return s[idx].Val
}

func (s Slots) SetLong(idx int, val int64) {
	s[idx].Val = int32(val)
}

func (s Slots) GetLong(idx int) int64 {
	return int64(s[idx].Val)
}

func (s Slots) SetFloat(idx int, val float32) {
	s[idx].Val = int32(val)
}

func (s Slots) GetFloat(idx int) float32 {
	return float32(s[idx].Val)
}

func (s Slots) SetDouble(idx int, val float64) {
	s[idx].Val = int32(val)
}

func (s Slots) GetDouble(idx int) float64 {
	return float64(s[idx].Val)
}

func (s Slots) SetRef(idx int, ref *Object) {
	s[idx].Ref = ref
}

func (s Slots) GetRef(idx int) *Object {
	return s[idx].Ref
}