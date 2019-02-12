package reflect

import "errors"

var (
	ErrorMethodNotFound = errors.New("reflect: method not found")
	ErrorClassMethodMatch = errors.New("reflect: class method in interface")
	ErrorIfaceMethodMatch = errors.New("reflect: interface method in class")
)

func (c *Class) GetMain() (m *Method, err error) {
	return c.GetStatic("main", "([Ljava/lang/String;)V")
}

func (c *Class) GetClinit() (m *Method, err error) {
	return c.GetStatic("<clinit>", "()V")
}

func (c *Class) GetStatic(name, desc string) (m *Method, err error) {
	for _, m := range c.Methods {
		if m.Name == name && m.Desc == desc {
			return m, nil
		}
	}
	err = ErrorMethodNotFound
	return
}

/* Lookup ClassMethod specified in JVM 8
* 1. return error if c is interface
* 2. look up method in c and superclasses (not searching for interface{})
* 3. [MaxSpecific SuperInterface] and non-abstract method
* 4. lookup in super interface and non-static and non-private
* 5. Fail
*/
func (c *Class) LookupMethod(name, desc string) (m *Method, err error) {
	// 1
	if !c.IsClass() {
		err = ErrorClassMethodMatch
		return
	}
	// 2
	m, err = lookupSuperClassMethod(c, name, desc)
	if err == nil {
		return
	}
	// 3
	for _, iface := range c.Interfaces {
		m, err = lookupMaxSpecificMethod(iface, name, desc)
		if err == nil {
			return
		}
	}
	// 4
	for _, iface := range c.Interfaces {
		m, err = lookupSuperInterfaceMethod(iface, name, desc)
		if err == nil {
			return
		}
	}
	// 5
	err = ErrorMethodNotFound
	return
}

/* Lookup IfaceMethod specified in JVM8
* 1. error if c is class
* 2. lookup in c
* 3. lookup in Object, public and not-static
* 4. [MaxSpecific SuperInterface] and non-abstract method
* 5. lookup in super interface and non-static and non-private
* 6. Fail
 */
func (c *Class) LookupIfaceMethod(name, desc string) (m *Method, err error) {
	// 1
	if !c.IsInterface() {
		err = ErrorIfaceMethodMatch
		return
	}
	// 2
	for _, m = range c.Methods {
		if m.Name == name && m.Desc == desc {
			return
		}
	}
	// 3
	// skip it
	// 4
	m, err = lookupMaxSpecificMethod(c, name, desc)
	if err == nil {
		return
	}
	// 5
	m, err = lookupSuperInterfaceMethod(c, name, desc)
	if err == nil {
		return
	}
	// 6
	err = ErrorMethodNotFound
	return
}

func (c *Class) NewObject() (o *Object, err error) {
	// create instance of class
	o = &Object{
		Class:c,
		Fields:make(Slots, c.InstanceSlotCount),
	}
	return
}

// look up method in class and super classes
func lookupSuperClassMethod(c *Class, name, desc string) (m *Method, err error) {
	for cls := c; cls != nil; cls = cls.Super {
		for _, m = range cls.Methods {
			// case1: method defined in Class c
			// case2: method defined in SuperClass of c, but it is non-private
			if m.Name == name && m.Desc == desc && (!m.IsPrivate() || c == cls) {
				return
			}
		}
	}
	err = ErrorMethodNotFound
	return
}

// lookup [MaxSpecific SuperInterface] and non-abstract method which means:
// method is defined in iface or SuperIface, is non-private, not-static, non-abstract, no duplicate
func lookupMaxSpecificMethod(iface *Class, name, desc string) (m *Method, err error) {
	for ifc := iface; ifc != nil; ifc = ifc.Super {
		for _, m = range ifc.Methods {
			if m.IsPrivate() || m.IsStatic() || m.IsAbstract() {
				continue
			}
			return
		}
	}
	err = ErrorMethodNotFound
	return
}

// lookup in super interface and non-static and non-private
func lookupSuperInterfaceMethod(iface *Class, name, desc string) (m *Method, err error) {
	for ifc := iface.Super; ifc != nil; ifc = ifc.Super {
		for _, m = range ifc.Methods {
			if m.IsPrivate() || m.IsStatic() {
				continue
			}
			return
		}
	}
	err = ErrorMethodNotFound
	return
}