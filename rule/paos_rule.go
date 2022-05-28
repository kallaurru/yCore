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

	// МЕСТОИМЕНИЯ

	// IsPointer - указательное местоимение
	IsPointer uint32 = 0x10

	// ГЛАГОЛ

	// IsReflexiveChange - изменяется по категории переходности
	IsReflexiveChange uint32 = 0x01
	// IsTransitiveVerb - on - переходный глагол, off - непереходный глагол
	IsTransitiveVerb uint32 = 0x10
	// IsNonReflexive  - не возвратный глагол
	IsNonReflexive uint32 = 0x20
	// IsReflexive - возвратный глагол
	IsReflexive uint32 = 0x40
	// Imperative - побудительное наклонение глагола
	Imperative uint32 = 0x80
)
