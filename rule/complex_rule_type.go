package rule

const (
	/* младший байт категория числа и категория рода */

	// NumbersIsChanges изменяется по числам
	NumbersIsChanges uint32 = 0x01
	// NumbersIsSingle единственное число
	NumbersIsSingle uint32 = 0x02
	// NumbersIsPlural множественное число
	NumbersIsPlural uint32 = 0x04

	/* категория рода */

	// GenusIsChanges измеряется по родам
	GenusIsChanges uint32 = 0x08
	// GenusFemale женский
	GenusFemale uint32 = 0x10
	// GenusMale мужской
	GenusMale uint32 = 0x20
	// GenusMiddle средний
	GenusMiddle uint32 = 0x40
	// GenusCommon Общий
	GenusCommon uint32 = 0x80

	/* второй байт категория лица и времени */

	/* категория лица */

	// PersonIsChanges изменяется по лицам
	PersonIsChanges uint32 = 0x01
	// PersonFirst 1-ое лицо
	PersonFirst uint32 = 0x02
	// PersonSecond 2-ое лицо
	PersonSecond uint32 = 0x04
	// PersonThird 3-e лицо
	PersonThird uint32 = 0x08

	/* категория времени */

	// TimeIsChanges изменяется по категории времени
	TimeIsChanges uint32 = 0x10
	// TimeFuture будущее время
	TimeFuture uint32 = 0x20
	// TimePresent настоящее время
	TimePresent uint32 = 0x40
	// TimePast прошедшее время
	TimePast uint32 = 0x80

	/* Третий байт категория падежа */

	// CaseIsChanges изменяется по падежам
	CaseIsChanges uint32 = 0x01
	// CaseNom именительный падеж
	CaseNom uint32 = 0x02
	// CaseGen родительный падеж
	CaseGen uint32 = 0x04
	// CaseDat дательный падеж
	CaseDat uint32 = 0x08
	// CaseAcc винительный падеж
	CaseAcc uint32 = 0x10
	// CaseIns творительный падеж
	CaseIns uint32 = 0x20
	// CasePre предложный падеж
	CasePre uint32 = 0x40

	// NumbersTypeCategory категория числа
	NumbersTypeCategory uint32 = 0x01
	// GenusTypeCategory категория рода
	GenusTypeCategory uint32 = 0x02
	// PersonTypeCategory категория лица
	PersonTypeCategory uint32 = 0x04
	// TimeTypeCategory категория времени
	TimeTypeCategory uint32 = 0x08
	// CaseTypeCategory категория падежа
	CaseTypeCategory uint32 = 0x10
)

type ComplexRule uint32

// IsValidComplexRuleElement проверка на валидность входящих данных
func IsValidComplexRuleElement(value uint32) bool {
	allElements := NumbersIsChanges | NumbersIsSingle | NumbersIsPlural |
		GenusIsChanges | GenusFemale | GenusMale | GenusMiddle | GenusCommon |
		PersonIsChanges | PersonFirst | PersonSecond | PersonThird |
		CaseIsChanges | CaseNom | CaseGen | CaseDat | CaseAcc | CaseIns | CasePre
	return allElements&value > 0
}

// MakeComplexRule желательно передавать данные полученные из констант
func MakeComplexRule(
	numberBytes uint32,
	genusBytes uint32,
	personBytes uint32,
	timeBytes uint32,
	caseBytes uint32,
) ComplexRule {

	var cr ComplexRule = 0
	cr |= ComplexRule(numberBytes & NumbersIsChanges)
	cr |= ComplexRule(numberBytes & NumbersIsSingle)
	cr |= ComplexRule(numberBytes & NumbersIsPlural)

	cr |= ComplexRule(genusBytes & GenusIsChanges)
	cr |= ComplexRule(genusBytes & GenusFemale)
	cr |= ComplexRule(genusBytes & GenusMale)
	cr |= ComplexRule(genusBytes & GenusMiddle)
	cr |= ComplexRule(genusBytes & GenusCommon)

	cr |= ComplexRule((personBytes & PersonIsChanges) << 8)
	cr |= ComplexRule((personBytes & PersonFirst) << 8)
	cr |= ComplexRule((personBytes & PersonSecond) << 8)
	cr |= ComplexRule((personBytes & PersonThird) << 8)

	cr |= ComplexRule((timeBytes & TimeIsChanges) << 8)
	cr |= ComplexRule((timeBytes & TimeFuture) << 8)
	cr |= ComplexRule((timeBytes & TimePresent) << 8)
	cr |= ComplexRule((timeBytes & TimePast) << 8)

	cr |= ComplexRule((caseBytes & CaseIsChanges) << 1)
	cr |= ComplexRule((caseBytes & CaseNom) << 16)
	cr |= ComplexRule((caseBytes & CaseGen) << 16)
	cr |= ComplexRule((caseBytes & CaseDat) << 16)
	cr |= ComplexRule((caseBytes & CaseAcc) << 16)
	cr |= ComplexRule((caseBytes & CaseIns) << 16)
	cr |= ComplexRule((caseBytes & CasePre) << 16)

	return cr
}

func (cr ComplexRule) SetIsNumberChanging() ComplexRule {
	return cr | ComplexRule(NumbersIsChanges)
}

func (cr ComplexRule) SetNumbersIsSingle() ComplexRule {
	return cr | ComplexRule(NumbersIsSingle)
}

func (cr ComplexRule) SetNumbersIsPlural() ComplexRule {
	return cr | ComplexRule(NumbersIsPlural)
}

func (cr ComplexRule) SetIsGenusChanging() ComplexRule {
	return cr | ComplexRule(GenusIsChanges)
}

func (cr ComplexRule) SetGenusFemale() ComplexRule {
	return cr | ComplexRule(GenusFemale)
}

func (cr ComplexRule) SetGenusMale() ComplexRule {
	return cr | ComplexRule(GenusMale)
}

func (cr ComplexRule) SetGenusMiddle() ComplexRule {
	return cr | ComplexRule(GenusMiddle)
}

func (cr ComplexRule) SetGenusCommon() ComplexRule {
	return cr | ComplexRule(GenusCommon)
}

// второй байт. Делаем сдвиг на 8

func (cr ComplexRule) SetIsPersonChanging() ComplexRule {
	return cr | ComplexRule(PersonIsChanges<<8)
}

func (cr ComplexRule) SetPeronFirst() ComplexRule {
	return cr | ComplexRule(PersonFirst<<8)
}

func (cr ComplexRule) SetPersonSecond() ComplexRule {
	return cr | ComplexRule(PersonSecond<<8)
}

func (cr ComplexRule) SetPersonThird() ComplexRule {
	return cr | ComplexRule(PersonThird<<8)
}

func (cr ComplexRule) SetIsTimeChanging() ComplexRule {
	return cr | ComplexRule(TimeIsChanges<<8)
}

func (cr ComplexRule) SetTimeFuture() ComplexRule {
	return cr | ComplexRule(TimeFuture<<8)
}

func (cr ComplexRule) SetTimePresent() ComplexRule {
	return cr | ComplexRule(TimePresent<<8)
}

func (cr ComplexRule) SetTimePast() ComplexRule {
	return cr | ComplexRule(TimePast<<8)
}

/** Третий байт сдвигаем на 16 */

func (cr ComplexRule) SetIsCaseChanging() ComplexRule {
	return cr | ComplexRule(CaseIsChanges<<16)
}

func (cr ComplexRule) SetCaseNom() ComplexRule {
	return cr | ComplexRule(CaseNom<<16)
}

func (cr ComplexRule) SetCaseGen() ComplexRule {
	return cr | ComplexRule(CaseGen<<16)
}

func (cr ComplexRule) SetCaseDat() ComplexRule {
	return cr | ComplexRule(CaseDat<<16)
}

func (cr ComplexRule) SetCaseAcc() ComplexRule {
	return cr | ComplexRule(CaseAcc<<16)
}

func (cr ComplexRule) SetCaseIns() ComplexRule {
	return cr | ComplexRule(CaseIns<<16)
}

func (cr ComplexRule) SetCasePre() ComplexRule {
	return cr | ComplexRule(CasePre<<16)
}
