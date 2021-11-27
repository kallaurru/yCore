package rule

// RulesManager можно собрать правильные значения и выгрузить линейку
type RulesManager struct {
	value uint32
	vf    ValidRuleElementFunc
}

// ValidRuleElementFunc функция контроля входных данных.
type ValidRuleElementFunc func(value uint32) bool

func MakeNewRulesManager(vf ValidRuleElementFunc) *RulesManager {
	return &RulesManager{
		vf:    vf,
		value: 0,
	}
}

func (rm *RulesManager) isValid(value uint32) bool {
	res := rm.vf(value)
	if res == false {
		return res
	}
	rm.value |= value
	return true
}

func (rm *RulesManager) isValidPos(pos int) bool {
	if pos < 0 || pos > 31 {
		return false
	}
	return true
}

func (rm *RulesManager) IsSet(pos int) bool {
	isValid := rm.isValidPos(pos)
	if !isValid {
		return false
	}

	return (rm.value & 1 << pos) > 0
}

func (rm *RulesManager) SetByte(pos int) {
	isValid := rm.isValidPos(pos)
	if !isValid {
		return
	}
	newValue := uint32(1 << pos)
	if rm.isValid(newValue) == false {
		return
	}
	rm.value |= newValue
}

func (rm *RulesManager) AddData(value uint32) {
	if rm.isValid(value) == false {
		return
	}
	rm.value |= value
}

func (rm *RulesManager) ToUint32() uint32 {
	return rm.value
}
