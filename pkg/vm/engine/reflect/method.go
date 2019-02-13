package reflect

import (
	"errors"
	"strings"
)

var (
	ErrorMethodNotFound = errors.New("reflect: method not found")
	ErrorClassMethodMatch = errors.New("reflect: class method in interface")
	ErrorIfaceMethodMatch = errors.New("reflect: interface method in class")
	ErrorBadMethodDescriptor = errors.New("bad method descriptor")
)

func (m *Method) ParseSignature() (err error) {
	md, err := parseSignature(m.Desc)
	if err != nil {
		return
	}
	m.ArgSlot = 0
	for _, arg := range md.Args {
		switch arg {
		case "D","J":
			m.ArgSlot = m.ArgSlot+2
		default:
			m.ArgSlot = m.ArgSlot+1
		}
	}
	// if method is not static, [this] ref will be added as an argument
	if !m.IsStatic() {
		m.ArgSlot = m.ArgSlot + 1
	}
	return
}


func (c *Class) GetStatic(name, desc string) (m *Method, err error) {
	return lookupCurrent(c, name, desc)
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
	m, err = lookupCurrent(c, name, desc)
	if err == nil {
		return
	}
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
	m, err = lookupCurrent(c, name, desc)
	if err != nil {
		return
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

/*
only instance method is searched
1. lookup in c
2. if c is class, search its super and further
3. if c is interface, search public method in Object
4. search maximally-specific method (§5.4.3.3) in the superinterfaces of C and not abstract
 */
func (c *Class) LookupSpecialMethod(m *Method) (newm *Method, err error) {
	name, desc := m.Name, m.Desc
	// 1
	m, err = lookupCurrent(c, name, desc)
	if err == nil && !m.IsStatic() {
		return
	}
	// 2
	if c.IsClass() {
		m, err = lookupSuperClassMethod(c, name, desc)
		if err == nil && !m.IsStatic() {
			return
		}
	}
	// 3, skip
	// 4
	if c.IsInterface() {
		m, err = lookupMaxSpecificMethod(c, name, desc)
		if err == nil && !m.IsStatic() {
			return
		}
	}
	err = ErrorMethodNotFound
	return
}


// look up method in class and super classes
func lookupSuperClassMethod(c *Class, name, desc string) (m *Method, err error) {
	for cls := c.Super; cls != nil; cls = cls.Super {
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

// lookup in current class
func lookupCurrent(c *Class, name, desc string) (m *Method, err error)  {
	for _, m = range c.Methods {
		if m.Name == name && m.Desc == desc {
			return
		}
	}
	err = ErrorMethodNotFound
	return
}

func parseSignature(sig string) (d *MethodDescriptor, err error) {
	l := len(sig)
	d = new(MethodDescriptor)
	idx := 0
	// check begin args
	if idx >= l || sig[idx] != '(' {
		err = ErrorBadMethodDescriptor
		return
	}
	idx = idx + 1
	// parse args
	args:
	for ; idx < l; idx = idx + 1 {
		switch sig[idx] {
		case 'B','C','D','F','I','J','S','Z':
			d.Args = append(d.Args, string(sig[idx]))
		case 'L':
			next := strings.Index(sig[idx:], ";")
			if next < 0 {
				err = ErrorBadMethodDescriptor
				return
			}
			d.Args = append(d.Args, sig[idx+1:idx+next])
			idx = idx + next
		case '[':
			continue
		case ')': //end of args detected, break args loop
			break args
		default:
			err = ErrorBadMethodDescriptor
			return
		}
	}
	// check end of args
	if idx >= l || sig[idx] != ')' {
		err = ErrorBadMethodDescriptor
		return
	}
	idx = idx + 1
	// parse return
	for ; idx < l; idx = idx + 1 {
		switch sig[idx] {
		case 'B','C','D','F','I','J','S','Z','V':
			d.Return = string(sig[idx])
		case 'L':
			next := strings.Index(sig[idx:], ";")
			if next < 0 {
				err = ErrorBadMethodDescriptor
				return
			}
			d.Return = sig[idx+1:idx+next]
			idx = idx + next
		case '[':
			continue
		default:
			err = ErrorBadMethodDescriptor
			return
		}
		break //if got return type, break the loop
	}
	idx = idx + 1
	// check end
	if idx != l {
		err = ErrorBadMethodDescriptor
		return
	}
	return
}
