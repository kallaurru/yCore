package rule

/** Специфические для каждой части речи свойства */

const (
	//  СУЩЕСТВИТЕЛЬНОЕ */

	// IsAlive  on - одушевленное, off не одушевленное
	IsAlive uint32 = 0x10
	// IsSelf on - имя собственное, off
	IsSelf uint32 = 0x20

	//
)

func IsValidPaosRuleElement(paos, value uint32) bool {
	return false
}
