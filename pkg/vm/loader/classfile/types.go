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
		name string
		data interface{}
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
	MemberConst struct {
		class u2
		nametype u2
	}
	// attr types
	LineNumber struct {
		pc u2
		line u2
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
		codes []byte
		Exceptions []ExceptionHandle
		Attributes []Attribute
	}
	ExceptionHandle struct {
		start_pc u2
		end_pc u2
		handle_pc u2
		catch_type u2
	}
)