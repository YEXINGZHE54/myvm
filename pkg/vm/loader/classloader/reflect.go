package classloader

func (l *loader) prepareReflect() (err error){
	ccls, err := l.LoadClass("java/lang/Class")
	if err != nil {
		return
	}
	for _, cls := range l.classes {
		cls.ClsObj, err = ccls.NewObject()
		if err != nil {
			return
		}
		cls.ClsObj.Extra = cls
		// sets up loader
		setupLoader(ccls, cls.ClsObj, l)
	}
	return
}