package classfile

type (
	u1 uint8
	u2 uint16
	u4 uint32
	u8 uint64
	ClassFile struct {
		Magic u4
		Minor u2
		Major u2
		Constants []Constant
		AccessFlags u2
		This u2
		Super u2
		Ifaces []u2
		Fields []Member
		Methods []Member
		Attributes []Attribute
	}
	Constant interface{}
	Member struct {
		AccessFlags u2
		NameIndex u2
		DescIndex u2
		Attributes []Attribute
	}
	Attribute struct {
		Name string
		Data interface{}
	}
	// constant types
	NullConst struct{}
	IntegerConst int32
	FloatConst float32
	LongConst int64
	DoubleConst float64
	UTF8Const string
	StringConst u2
	ClassConst u2
	NameTypeConst struct {
		name u2
		desc u2
	}
	FieldConst struct {
		Class u2
		Nametype u2
	}
	MethodConst struct {
		Class u2
		Nametype u2
	}
	IfaceMethodConst struct {
		Class u2
		Nametype u2
	}
	// attr types
	LineNumber struct {
		PC u2
		Line u2
	}
	LocalVariable struct {
		pc u2
		pclen u2
		name u2
		desc u2
		index u2
	}
	Code struct {
		MaxStacks u2
		MaxLocals u2
		Codes []byte
		Exceptions []ExceptionHandle
		Attributes []Attribute
	}
	ExceptionHandle struct {
		StartPC u2
		EndPC u2
		HandlePC u2
		CatchType u2
	}
)

func ToIdx(v interface{}) u2 {
	switch val := v.(type) {
	case StringConst:
		return u2(val)
	case ClassConst:
		return u2(val)
	default:
	}
	return v.(u2)
}