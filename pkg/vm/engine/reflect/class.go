package reflect

import "errors"

var (
	ErrorMethodNotFound = errors.New("reflect: method not found")
)

func (c *Class) GetMain() (m *Method, err error) {
	for _, m := range c.Methods {
		if m.Name == "main" &&
			m.Desc == "([Ljava/lang/String;)V" {
			return m, nil
		}
	}
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