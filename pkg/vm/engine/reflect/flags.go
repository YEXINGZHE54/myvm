package reflect

var (
	primitives map[string]bool = make(map[string]bool)
)

func (c *Class) IsClass() bool {
	return c.Flag & ACCESS_INTERFACE == 0
}

func (c *Class) IsInterface() bool {
	return c.Flag & ACCESS_INTERFACE > 0
}

func (c *Class) IsSuperSet() bool {
	return c.Flag & ACCESS_SUPER > 0
}

func (c *Class) IsArray() bool {
	return c.Name[0] == '['
}

func (c *Class) IsPrimitive() bool {
	return primitives[c.Name]
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

func (f *Field) IsPublic() bool {
	return f.Flag & ACCESS_PUBLIC > 0
}

func (m *Method) IsPrivate() bool {
	return m.Flag & ACCESS_PRIVATE > 0
}

func (m *Method) IsStatic() bool {
	return m.Flag & ACCESS_STATIC > 0
}

func (m *Method) IsAbstract() bool {
	return m.Flag & ACCESS_ABSTRACT > 0
}

func (m *Method) IsProtected() bool {
	return m.Flag & ACCESS_PROTECTED > 0
}

func (m *Method) IsPublic() bool {
	return m.Flag & ACCESS_PUBLIC > 0
}

func (m *Method) IsNative() bool {
	return m.Flag & ACCESS_NATIVE > 0
}

func init()  {
	prims := []string{
		"void",
		"boolean",
		"byte",
		"char",
		"short",
		"int",
		"long",
		"float",
		"double",
	}
	for _, p := range prims {
		primitives[p] = true
	}
}