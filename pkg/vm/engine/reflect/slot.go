package reflect

func (s Slots) SetVal(idx uint, val int32) {
	s[idx].Val = val
}

func (s Slots) GetVal(idx uint) int32 {
	return s[idx].Val
}

func (s Slots) SetLong(idx uint, val int64) {
	s[idx].Val = int32(val)
}

func (s Slots) GetLong(idx uint) int64 {
	return int64(s[idx].Val)
}

func (s Slots) SetDouble(idx uint, val float64) {
	s[idx].Val = int32(val)
}

func (s Slots) GetDouble(idx uint) float64 {
	return float64(s[idx].Val)
}

func (s Slots) SetRef(idx uint, ref *Object) {
	s[idx].Ref = ref
}

func (s Slots) GetRef(idx uint) *Object {
	return s[idx].Ref
}