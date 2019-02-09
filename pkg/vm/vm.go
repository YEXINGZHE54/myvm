package vm

type (
	VM interface {
		Startup(class string, args []string) error
	}
	VMConstructor func(bootPath, classPath string) VM
)

var (
	factory VMConstructor
)

func NewVM(bootPath, classPath string) VM {
	return factory(bootPath, classPath)
}

func Register(f VMConstructor) {
	factory = f
}