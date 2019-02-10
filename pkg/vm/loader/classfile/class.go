package classfile

func (f *ClassFile) GetConst(idx u2) Constant {
	return f.Constants[idx]
}

func (f *ClassFile) GetUTF8(idx u2) string {
	return string(f.GetConst(idx).(UTF8Const))
}

func (f *ClassFile) GetNameType(idx u2) (string, string) {
	nt := f.GetConst(idx).(*NameTypeConst)
	name := f.GetUTF8(nt.name)
	_type := f.GetUTF8(nt.desc)
	return name, _type
}

func (f *ClassFile) GetClass(idx u2) string {
	clstype := f.GetConst(idx).(ClassConst)
	return f.GetUTF8(u2(clstype))
}

func (f *ClassFile) GetMain() *Member {
	for _, m := range f.Methods {
		if f.GetUTF8(m.NameIndex) == "main" && 
		f.GetUTF8(m.DescIndex) == "([Ljava/lang/String;)V" {
			return &m
		}
	}
	return nil
}

func (m *Member) GetCode() *Code {
	for _, attr := range m.Attributes {
		if attr.name == "Code" {
			return attr.data.(*Code)
		}
	}
	return nil
}

func (c *Code) GetReader() *CodeReader {
	return NewCodeReader(c.codes)
}