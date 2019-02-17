package classloader

func (l *loader) prepareReflect() (err error){
	c, err := l.LoadClass("java/lang/Class")
	if err != nil {
		return
	}
	for _, cls := range l.classes {
		cls.ClsObj, err = c.NewObject()
		if err != nil {
			return
		}
		cls.ClsObj.Extra = cls
	}
	return
}