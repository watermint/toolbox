package elocale

type localeImpl struct {
	data LocaleData
}

func (z localeImpl) String() string {
	return z.data.Tag
}

func (z localeImpl) Language() string {
	return z.data.Lang
}

func (z localeImpl) LanguageTwoLetter() string {
	if len(z.data.Lang) == 2 {
		return z.data.Lang
	}
	if two, ok := iso631ThreeToTwoLetter[z.data.Lang]; ok {
		return two
	}
	return ""
}

func (z localeImpl) LanguageExtended() string {
	return z.data.LangExtended
}

func (z localeImpl) Extension() string {
	return z.data.Extension
}

func (z localeImpl) Script() string {
	return z.data.Script
}

func (z localeImpl) Variant() string {
	return z.data.Variant
}

func (z localeImpl) Region() string {
	return z.data.Region
}

func (z localeImpl) Data() LocaleData {
	return z.data
}
