package rule

type Rule uint32

type Rules interface {
	IsSet(pos int) bool
	SetByte(pos int) Rule
}

func (r Rule) IsSet(pos int) bool {
	if pos < 0 || pos > 31 {
		return false
	}
	return (r & 1 << pos) > 0
}

func (r Rule) SetByte(pos int) Rule {
	if pos < 0 || pos > 31 {
		return r
	}
	return r | 1<<pos
}
