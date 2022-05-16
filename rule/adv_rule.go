package rule

// UncensoredWord uncensored word, slang, abbreviation
const (
	// UncensoredWord признак не цензурного слова
	UncensoredWord uint32 = 0x01

	// SlangWord признак сленгового слова. Требует связи слова со словом оригиналом
	SlangWord uint32 = 0x02

	// AbbreviationWord признак того, что слово аббревиатура
	AbbreviationWord uint32 = 0x04

	// SpecialWord признак спец слова. Email, Тэг, Сумма, Телефон и т.д.
	SpecialWord uint32 = 0x08

	// ModelWord слово с буквами и цифрами и возможно дефисом
	ModelWord uint32 = 0x10
)

// IsValidAdvRuleElement проверка на валидность входящих данных
func IsValidAdvRuleElement(value uint32) bool {
	allElements := UncensoredWord | SlangWord | AbbreviationWord | SpecialWord | ModelWord
	return allElements&value > 0
}
