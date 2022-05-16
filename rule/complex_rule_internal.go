package rule

func getDiffBytes(typeCategory uint32) int {
	switch typeCategory {
	case NumbersTypeCategory:
	case PersonTypeCategory:
	case TimeTypeCategory:
		return 8
	case GenusTypeCategory:
		return 16
	case CaseTypeCategory:
		return 24
	}
	return 0
}

func getNumbersMask() uint32 {
	return NumbersIsSingle | NumbersIsPlural
}

func getPersonsMask() uint32 {
	return PersonFirst | PersonSecond | PersonThird
}

func getTimeMask() uint32 {
	return TimePast | TimePresent | TimeFuture
}

func getGenusMask() uint32 {
	return GenusCommon | GenusFemale | GenusMale | GenusMiddle
}

func getCaseMask() uint32 {
	return CaseNom | CaseGen | CaseDat | CaseAcc | CaseIns | CasePre
}

func getChangingFlagsMask() uint32 {
	return 0xff
}
