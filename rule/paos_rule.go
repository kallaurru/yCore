package rule

/** Специфические для каждой части речи свойства */

const (
	//  СУЩЕСТВИТЕЛЬНОЕ

	// IsAlive  on - одушевленное, off не одушевленное
	IsAlive uint32 = 0x10
	// IsSelf on - имя собственное, off
	IsSelf uint32 = 0x20

	// ПРИЛАГАТЕЛЬНОЕ

	// IsFull - полная форма прилагательного
	IsFull uint32 = 0x01
	// IsShort - краткая форма прилагательного
	IsShort uint32 = 0x02
	// IsQuality - качественное прилагательное
	IsQuality uint32 = 0x04
	// IsPredicative - предикативное прилагательное
	IsPredicative uint32 = 0x100
	// IsRelativeAdj -  относительное прилагательное
	IsRelativeAdj uint32 = 0x200

	// МЕСТОИМЕНИЯ

	// IsPointer - указательное местоимение
	IsPointer uint32 = 0x10
	// IsPersonal - персональное местоимение
	IsPersonal uint32 = 0x20
	// IsRelative -  относительное местоимение
	IsRelative uint32 = 0x40

	// ГЛАГОЛ

	// IsPerfectForm - совершенный вид глагола
	IsPerfectForm uint32 = 0x01
	// IsImperfectForm - несовершенный вид глагола
	IsImperfectForm uint32 = 0x02

	// IsTransitiveVerb - on - переходный глагол, off - непереходный глагол
	IsTransitiveVerb uint32 = 0x10
	// IsNonReflexive  - не возвратный глагол
	IsNonReflexive uint32 = 0x20
	// IsReflexive - возвратный глагол
	IsReflexive uint32 = 0x40
	// Imperative - побудительное наклонение глагола
	Imperative uint32 = 0x80
	// Imperative2 - повелитeльное наклонение глагола
	Imperative2 uint32 = 0x100
)
