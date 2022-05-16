package rule

const (
	// AdvAbbrev - аббревиатура
	AdvAbbrev uint32 = 0x01
	// AdvSpecWords - email, url, цифровые коды, почтовые индексы
	AdvSpecWords uint32 = 0x02
	// AdvSlang - сленговое слово
	AdvSlang uint32 = 0x04
	// AdvObscene - нецензурное слово
	AdvObscene uint32 = 0x08
)

func IsValidAdvElement(value uint32) bool {
	allAdv := AdvAbbrev | AdvSpecWords | AdvSlang | AdvObscene

	return (allAdv & value) > 0
}
