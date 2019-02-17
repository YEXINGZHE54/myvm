package classfile

import (
	"errors"
	"math"
)

const (
	MAGIC = 0xCAFEBABE
	CONSTANT_Utf8 = 1
	CONSTANT_Integer = 3
	CONSTANT_Float = 4
	CONSTANT_Long = 5
	CONSTANT_Double = 6
	CONSTANT_Class = 7
	CONSTANT_String = 8
	CONSTANT_Field = 9
	CONSTANT_Method = 10
	CONSTANT_InterfaceMethod = 11
	CONSTANT_NameAndType = 12
	CONSTANT_MethodHandle = 15
	CONSTANT_MethodType = 16
	CONSTANT_InvokeDynamic = 18
)

var (
	ErrorMagic = errors.New("magic number not match")
	ErrorVersion = errors.New("version not supported")
	ErrorConst = errors.New("Unsupported constant type")
)

func Parse(data []byte) (f *ClassFile, err error) {
	r := NewReader(data)
	f = new(ClassFile)
	f.Magic, err = f.parseMagic(r)
	if err != nil {
		return
	}
	f.Minor, f.Major, err = f.parseVersion(r)
	if err != nil {
		return
	}
	f.Constants, err = f.parseConstant(r)
	if err != nil {
		return
	}
	f.AccessFlags = f.parseFlag(r)
	f.This, f.Super = f.parseClass(r)
	f.Ifaces = f.parseIfaces(r)
	f.Fields = f.parseMembers(r)
	f.Methods = f.parseMembers(r)
	f.Attributes = f.parseAttributes(r)
	return
}

func (f *ClassFile) parseMagic(r *Reader) (u4, error) {
	mag := r.read4()
	if mag != MAGIC {
		return 0, ErrorMagic
	}
	return mag, nil
}

func (f *ClassFile) parseVersion(r *Reader) (minor, major u2, err error) {
	minor = r.read2()
	major = r.read2()
	if major > 52 || (major > 45 && major <= 52 && minor > 0) {
		err = ErrorVersion
		return
	}
	return
}

func (f *ClassFile) parseConstant(r *Reader) (consts []Constant, err error) {
	consts = make([]Constant, int(r.read2()))
	consts[0] = NullConst{}
	tag := 0
	for i := 1; i < len(consts); i = i + 1 {
		tag = int(r.read1())
		switch tag {
		case CONSTANT_Integer:
			consts[i] = IntegerConst(r.read4())
		case CONSTANT_Float:
			consts[i] = FloatConst(math.Float32frombits(uint32(r.read4())))
		case CONSTANT_Long:
			consts[i] = LongConst(r.read8())
			i = i + 1
		case CONSTANT_Double:
			consts[i] = DoubleConst(math.Float64frombits(uint64(r.read8())))
			i = i + 1
		case CONSTANT_Utf8:
			consts[i] = UTF8Const(r.readBytes(int(r.read2())))
		case CONSTANT_String:
			consts[i] = StringConst(r.read2())
		case CONSTANT_Class:
			consts[i] = ClassConst(r.read2())
		case CONSTANT_NameAndType:
			consts[i] = &NameTypeConst{r.read2(), r.read2()}
		case CONSTANT_Field:
			consts[i] = &FieldConst{r.read2(), r.read2()}
		case CONSTANT_Method:
			consts[i] = &MethodConst{r.read2(), r.read2()}
		case CONSTANT_InterfaceMethod:
			consts[i] = &IfaceMethodConst{r.read2(), r.read2()}
		case CONSTANT_MethodHandle:
			r.read1()
			r.read2()
		case CONSTANT_MethodType:
			r.read2()
		case CONSTANT_InvokeDynamic:
			r.read2()
			r.read2()
		default:
			println("const type: ")
			println(r.index)
			println(tag)
			err = ErrorConst
			return
		}
	}
	return
}

func (f *ClassFile) parseFlag(r *Reader) (flag u2) {
	flag = r.read2()
	return
}

func (f *ClassFile) parseClass(r *Reader) (this, super u2) {
	this = r.read2()
	super = r.read2()
	return
}

func (f *ClassFile) parseIfaces(r *Reader) (ifs []u2) {
	cnt := int(r.read2())
	for i := 0; i < cnt; i = i + 1 {
		ifs = append(ifs, r.read2())
	}
	return
}

func (f *ClassFile) parseMembers(r *Reader) (fields []Member) {
	mlen := int(r.read2())
	for i := 0; i < mlen; i = i + 1 {
		fields = append(fields, Member{
			AccessFlags: r.read2(),
			NameIndex: r.read2(),
			DescIndex: r.read2(),
			Attributes: f.parseAttributes(r),
		})
	}
	return
}

func (f *ClassFile) parseAttributes(r *Reader) (attrs []Attribute) {
	attrlen := int(r.read2())
	for i := 0; i < attrlen; i = i + 1 {
		name := f.GetUTF8(r.read2())
		len := int(r.read4())
		switch name {
		case "Code":
			code := &Code{
				r.read2(),
				r.read2(),
				r.readBytes(int(r.read4())),
				f.parseExceptionHandler(r),
				f.parseAttributes(r),
			}
			attrs = append(attrs, Attribute{name, code})
		case "SourceFile":
			attrs = append(attrs, Attribute{name, r.read2()})
		case "ConstantValue":
			attrs = append(attrs, Attribute{name, r.read2()})
		case "Deprecated", "Synthetic":
			attrs = append(attrs, Attribute{name, nil})
		case "Exceptions":
			data := make([]u2, int(r.read2()))
			for i := range data {
				data[i] = r.read2()
			}
			attrs = append(attrs, Attribute{name, data})
		case "LineNumberTable":
			data := make([]LineNumber, int(r.read2()))
			for i := range data {
				data[i] = LineNumber{
					r.read2(), r.read2(),
				}
			}
			attrs = append(attrs, Attribute{name, data})
		case "LocalVariableTable":
			data := make([]LocalVariable, int(r.read2()))
			for i := range data {
				data[i] = LocalVariable{
					r.read2(), r.read2(), r.read2(), r.read2(),r.read2(),
				}
			}
			attrs = append(attrs, Attribute{name, data})
		default:
			// skip unknown attrs
			r.readBytes(len)
		}
	}
	return
}

func (f *ClassFile) parseExceptionHandler(r *Reader) []ExceptionHandle {
	data := make([]ExceptionHandle, int(r.read2()))
	for i := range data {
		data[i] = ExceptionHandle{
			r.read2(),
			r.read2(),
			r.read2(),
			r.read2(),
		}
	}
	return data
}