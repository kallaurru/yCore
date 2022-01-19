package rule

const (
	/* младший байт занимаем под флаги изменчивости по категориям и доп флаги */

	// NumbersIsChanges изменяется по числам
	NumbersIsChanges uint32 = 0x01
	// GenusIsChanges измеряется по родам
	GenusIsChanges uint32 = 0x02
	// PersonIsChanges изменяется по лицам
	PersonIsChanges uint32 = 0x04
	// CaseIsChanges изменяется по падежам
	CaseIsChanges uint32 = 0x08
	// TimeIsChanges изменяется по категории времени
	TimeIsChanges uint32 = 0x10
	// WordHasHomonym у слова есть омонимы
	WordHasHomonym uint32 = 0x20
	// Reserved-1 0x40
	// Reserved-2 0x80

	/* второй байт категория лица и времени */

	/* категория числа */

	// NumbersIsSingle единственное число
	NumbersIsSingle uint32 = 0x01
	// NumbersIsPlural множественное число
	NumbersIsPlural uint32 = 0x02

	/* категория лица */

	// PersonFirst 1-ое лицо
	PersonFirst uint32 = 0x04
	// PersonSecond 2-ое лицо
	PersonSecond uint32 = 0x08
	// PersonThird 3-e лицо
	PersonThird uint32 = 0x10

	/* категория времени */

	// TimeFuture будущее время
	TimeFuture uint32 = 0x20
	// TimePresent настоящее время
	TimePresent uint32 = 0x40
	// TimePast прошедшее время
	TimePast uint32 = 0x80

	/* Третий байт категория падежа */

	/* категория рода */

	// GenusFemale женский
	GenusFemale uint32 = 0x01
	// GenusMale мужской
	GenusMale uint32 = 0x02
	// GenusMiddle средний
	GenusMiddle uint32 = 0x04
	// GenusCommon Общий
	GenusCommon uint32 = 0x08
	// Reserved 0x10, 0x20, 0x40, 0x80

	/* Четвертый байт категория падежа */

	/* категория падежа */

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

// MakeComplexRule желательно передавать данные полученные из констант
func MakeComplexRule(
	categoryFlags uint32,
	numberBytes uint32,
	genusBytes uint32,
	personBytes uint32,
	timeBytes uint32,
	caseBytes uint32,
) ComplexRule {

	var (
		cr ComplexRule = 0
	)
	cr |= ComplexRule(categoryFlags & getChangingFlagsMask())

	diff := getDiffBytes(NumbersTypeCategory)
	cr |= ComplexRule((numberBytes & getNumbersMask()) << diff)

	diff = getDiffBytes(PersonTypeCategory)
	cr |= ComplexRule((personBytes & getPersonsMask()) << diff)

	diff = getDiffBytes(TimeTypeCategory)
	cr |= ComplexRule((timeBytes & getTimeMask()) << diff)

	diff = getDiffBytes(GenusTypeCategory)
	cr |= ComplexRule(genusBytes & getGenusMask() << diff)

	diff = getDiffBytes(CaseTypeCategory)
	cr |= ComplexRule((caseBytes & getCaseMask()) << diff)

	return cr
}

func (cr ComplexRule) SetIsNumberChanging() ComplexRule {
	return cr | ComplexRule(NumbersIsChanges)
}
func (cr ComplexRule) SetIsGenusChanging() ComplexRule {
	return cr | ComplexRule(GenusIsChanges)
}
func (cr ComplexRule) SetIsPersonChanging() ComplexRule {
	return cr | ComplexRule(PersonIsChanges)
}
func (cr ComplexRule) SetIsTimeChanging() ComplexRule {
	return cr | ComplexRule(TimeIsChanges)
}
func (cr ComplexRule) SetIsCaseChanging() ComplexRule {
	return cr | ComplexRule(CaseIsChanges)
}
func (cr ComplexRule) SetHasHomonym() ComplexRule {
	return cr | ComplexRule(WordHasHomonym)
}

func (cr ComplexRule) SetNumbersIsSingle() ComplexRule {
	return cr | ComplexRule(NumbersIsSingle<<getDiffBytes(NumbersTypeCategory))
}
func (cr ComplexRule) SetNumbersIsPlural() ComplexRule {
	return cr | ComplexRule(NumbersIsPlural<<getDiffBytes(NumbersTypeCategory))
}
func (cr ComplexRule) SetPeronFirst() ComplexRule {
	return cr | ComplexRule(PersonFirst<<getDiffBytes(PersonTypeCategory))
}
func (cr ComplexRule) SetPersonSecond() ComplexRule {
	return cr | ComplexRule(PersonSecond<<getDiffBytes(PersonTypeCategory))
}
func (cr ComplexRule) SetPersonThird() ComplexRule {
	return cr | ComplexRule(PersonThird<<getDiffBytes(PersonTypeCategory))
}
func (cr ComplexRule) SetTimeFuture() ComplexRule {
	return cr | ComplexRule(TimeFuture<<getDiffBytes(TimeTypeCategory))
}
func (cr ComplexRule) SetTimePresent() ComplexRule {
	return cr | ComplexRule(TimePresent<<getDiffBytes(TimeTypeCategory))
}
func (cr ComplexRule) SetTimePast() ComplexRule {
	return cr | ComplexRule(TimePast<<getDiffBytes(TimeTypeCategory))
}

func (cr ComplexRule) SetGenusFemale() ComplexRule {
	return cr | ComplexRule(GenusFemale<<getDiffBytes(GenusTypeCategory))
}

func (cr ComplexRule) SetGenusMale() ComplexRule {
	return cr | ComplexRule(GenusMale<<getDiffBytes(GenusTypeCategory))
}

func (cr ComplexRule) SetGenusMiddle() ComplexRule {
	return cr | ComplexRule(GenusMiddle<<getDiffBytes(GenusTypeCategory))
}

func (cr ComplexRule) SetGenusCommon() ComplexRule {
	return cr | ComplexRule(GenusCommon<<getDiffBytes(GenusTypeCategory))
}

func (cr ComplexRule) SetCaseNom() ComplexRule {
	return cr | ComplexRule(CaseNom<<getDiffBytes(CaseTypeCategory))
}
func (cr ComplexRule) SetCaseGen() ComplexRule {
	return cr | ComplexRule(CaseGen<<getDiffBytes(CaseTypeCategory))
}
func (cr ComplexRule) SetCaseDat() ComplexRule {
	return cr | ComplexRule(CaseDat<<getDiffBytes(CaseTypeCategory))
}
func (cr ComplexRule) SetCaseAcc() ComplexRule {
	return cr | ComplexRule(CaseAcc<<getDiffBytes(CaseTypeCategory))
}
func (cr ComplexRule) SetCaseIns() ComplexRule {
	return cr | ComplexRule(CaseIns<<getDiffBytes(CaseTypeCategory))
}
func (cr ComplexRule) SetCasePre() ComplexRule {
	return cr | ComplexRule(CasePre<<getDiffBytes(CaseTypeCategory))
}
