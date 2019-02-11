package reflect

func (f *Field) IsStatic() bool {
	return f.Flag & ACCESS_STATIC > 0
}

func (f *Field) IsFinal() bool {
	return f.Flag & ACCESS_FINAL > 0
}

func (f *Field) IsLongDouble() bool {
	return f.Desc == "J" || f.Desc == "D"
}