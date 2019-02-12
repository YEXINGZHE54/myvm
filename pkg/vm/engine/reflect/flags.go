package reflect

func (c *Class) IsClass() bool {
	return c.Flag & ACCESS_INTERFACE == 0
}

func (c *Class) IsInterface() bool {
	return c.Flag & ACCESS_INTERFACE > 0
}

func (f *Field) IsStatic() bool {
	return f.Flag & ACCESS_STATIC > 0
}

func (f *Field) IsFinal() bool {
	return f.Flag & ACCESS_FINAL > 0
}

func (f *Field) IsLongDouble() bool {
	return f.Desc == "J" || f.Desc == "D"
}

func (m *Method) IsPrivate() bool {
	return m.Flag & ACCESS_PRIVATE > 0
}

func (f *Method) IsStatic() bool {
	return f.Flag & ACCESS_STATIC > 0
}

func (f *Method) IsAbstract() bool {
	return f.Flag & ACCESS_ABSTRACT > 0
}