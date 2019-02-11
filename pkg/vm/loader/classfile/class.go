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