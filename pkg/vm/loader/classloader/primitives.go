package classloader

import "github.com/YEXINGZHE54/myvm/pkg/vm/engine/reflect"

var (
	prims = []string{
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
)

func (l *loader) loadPrim(clsname string) (c *reflect.Class) {
	return &reflect.Class{
		Flag: reflect.ACCESS_PUBLIC,
		Name: clsname,
		Loader: l,
		Started: true,
	}
}

func (l *loader) loadPrims() (err error) {
	ccls, err := l.LoadClass("java/lang/Class")
	if err != nil {
		return
	}
	for _, cname := range prims {
		c := l.loadPrim(cname)
		// setup reflect info
		c.ClsObj, err = ccls.NewObject()
		if err != nil {
			return
		}
		c.ClsObj.Extra = c
		// sets up loader
		setupLoader(ccls, c.ClsObj, l)
		l.classes[cname] = c
	}
	return
}