package structures

type BaseStruct struct {
	// импортируемые
	Integer int
	Str     string

	// приватные
	integer int
	str     string
}

// SetInteger импортируемый метод от объекта
func (s *BaseStruct) SetInteger(i int) {
	s.Integer = i
	s.integer = i
}

// SetString импортируемый метод от значения
func (s BaseStruct) SetString(str string) {
	s.Str = str
	s.setPrivateString(str)
}

// setPrivateString приватный метод структуры
func (s *BaseStruct) setPrivateString(str string) {
	s.str = str
}
