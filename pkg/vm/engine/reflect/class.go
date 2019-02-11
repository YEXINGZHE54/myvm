package reflect

import "errors"

var (
	ErrorMethodNotFound = errors.New("reflect: method not found")
)

func (f *Class) GetMain() (m *Method, err error) {
	for _, m := range f.Methods {
		if m.Name == "main" &&
			m.Desc == "([Ljava/lang/String;)V" {
			return m, nil
		}
	}
	err = ErrorMethodNotFound
	return
}