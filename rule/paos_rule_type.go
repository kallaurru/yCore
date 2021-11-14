package rule

const (
	// NounPaOS существительное
	NounPaOS uint32 = 0x01
	// AdjectivePaOS прилагательное
	AdjectivePaOS uint32 = 0x02
	// VerbPaOS глагол
	VerbPaOS uint32 = 0x04
	// ParticiplePaOS причастие
	ParticiplePaOS uint32 = 0x08
	// GerundPaOS деепричастие
	GerundPaOS uint32 = 0x10
	// AdverbPaOS наречие
	AdverbPaOS uint32 = 0x20
	// PretextPaOS предлог
	PretextPaOS uint32 = 0x40
	// PronounPaOS местоимение
	PronounPaOS uint32 = 0x80
	// UnionPaOS союз
	UnionPaOS uint32 = 0x100
)

//IsValidPaosElement проверка на валидность входящих данных
func IsValidPaosElement(value uint32) bool {
	allPaOS := NounPaOS | AdjectivePaOS | VerbPaOS | ParticiplePaOS | GerundPaOS | AdverbPaOS | PretextPaOS | PronounPaOS | UnionPaOS
	return allPaOS&value > 0
}
