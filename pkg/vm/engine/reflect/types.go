package reflect

type (
	Loader interface {
		LoadClass(cls string) (c *Class, err error)
	}
	Object struct {
		
	}
	Class struct {
		Flag uint16
		Name string
		SuperName string
		InterfaceNames []string
		Consts []interface{}
		Fields []*Field
		Methods []*Method
		Loader Loader
		Super *Class
		Interfaces []*Class
		instanceSlotCount uint
		staticSlotCount uint
		staticVars interface{}
	}
	Member struct {
		Flag uint16
		Name string
		Desc string
		Cls *Class
	}
	Field struct {
		Member
	}
	Method struct {
		Member
		MaxStack int
		MaxLocal int
		Codes []byte
	}
	ClsRef struct {
		Name string
		Ref *Class
	}
	MemberRef struct {
		ClsName string
		Name string
		Desc string
	}
	FieldRef struct {
		MemberRef
		Ref *Field
	}
	MethodRef struct {
		MemberRef
		Ref *Method
	}
)

const (
	ACCESS_PUBLIC = 0x0001
	ACCESS_PRIVATE = 0x0002
	ACCESS_PROTECTED = 0x0004
	ACCESS_STATIC = 0x0008
	ACCESS_FINAL = 0x00010
	ACCESS_SUPER = 0x0020
	ACCESS_SYNC = 0x0020
	ACCESS_VOLATILE = 0x0040
	ACCESS_BRIDGE = 0x0040
	ACCESS_TRANSIENT = 0x0080
	ACCESS_VARARGS = 0x0080
	ACCESS_NATIVE = 0x0100
	ACCESS_INTERFACE = 0x0200
	ACCESS_ABSTRACT = 0x0400
	ACCESS_STRICT = 0x0800
	ACCESS_SYNTHETIC = 0x1000
	ACCESS_ANNOT = 0x2000
	ACCESS_ENUM = 0x1000
)