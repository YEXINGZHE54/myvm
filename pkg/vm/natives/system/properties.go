package system

import (
	"github.com/YEXINGZHE54/myvm/pkg/vm/memory/stack"
	"github.com/YEXINGZHE54/myvm/pkg/vm/natives"
	"runtime"
)

// static method
func initProperties(f *stack.Frame)  {
	props := f.This()
	f.PushOpstackRef(props)
	method, err := props.Class.LookupMethod("setProperty", "(Ljava/lang/String;Ljava/lang/String;)Ljava/lang/Object;")
	if err != nil {
		panic(err)
	}
	for k, v := range _sysProps() {
		key, err := props.Class.Loader.JString(k)
		if err != nil {
			panic(err)
		}
		val, err := props.Class.Loader.JString(v)
		if err != nil {
			panic(err)
		}
		newf := stack.NewFrame(method)
		f.Stack.Push(newf)
		newf.SetLocalRef(0, props)
		newf.SetLocalRef(1, key)
		newf.SetLocalRef(2, val)
		newf.ExtendStack(1) //mock, a holder position to store result from method
	}
}

func _sysProps() map[string]string {
	return map[string]string{
		"java.version":         "1.8.0",
		"java.vendor":          "jvm.go",
		"java.vendor.url":      "https:///github.com/YEXINGZHE54/myvm",
		"java.home":            "", //TODO
		"java.class.version":   "52.0",
		"java.class.path":      "", //TODO
		"java.awt.graphicsenv": "sun.awt.CGraphicsEnvironment",
		"os.name":              runtime.GOOS,
		"os.arch":              runtime.GOARCH,
		"os.version":           "",             // todo
		"file.separator":       "/",            // todo os.PathSeparator
		"path.separator":       ":",            // todo os.PathListSeparator
		"line.separator":       "\n",           // todo
		"user.name":            "",             // todo
		"user.home":            "",             // todo
		"user.dir":             ".",            // todo
		"user.country":         "CN",           // todo
		"file.encoding":        "UTF-8",
		"sun.stdout.encoding":  "UTF-8",
		"sun.stderr.encoding":  "UTF-8",
	}
}

func init()  {
	natives.Register("java/lang/System", "initProperties", "(Ljava/util/Properties;)Ljava/util/Properties;", initProperties)
}