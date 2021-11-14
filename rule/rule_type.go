package rule

type Rule uint32

// ValidRuleElementFunc функция контроля входных данных.
type ValidRuleElementFunc func(value uint32) bool

type Rules interface {
	IsSet(pos int) bool
	SetByte(pos int) Rule
	IsValid(value uint32, validateFunc ValidRuleElementFunc) bool
	AddData(value uint32) Rule
}

//IsSet проверка установлен ли данный бит
func (r Rule) IsSet(pos int) bool {
	if pos < 0 || pos > 31 {
		return false
	}
	return (r & 1 << pos) > 0
}

//SetByte установить позицию в линейке
func (r Rule) SetByte(pos int) Rule {
	if pos < 0 || pos > 31 {
		return r
	}
	return r | Rule(1<<pos)
}

//AddData добавить данные в линейку
func (r Rule) AddData(value uint32) Rule {
	return r | Rule(value)
}

//IsValid проверка на валидность входящих данных для текущей линейки
func (r Rule) IsValid(value uint32, validateFunc ValidRuleElementFunc) bool {
	return validateFunc(value)
}
