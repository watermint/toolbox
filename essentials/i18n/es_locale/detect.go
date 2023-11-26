package es_locale

// CurrentLocale returns detected language.
// Fallback to the default language if no default language detected.
// The method will not return nil.
func CurrentLocale() Locale {
	tag, err := currentLocaleString()

	if err != nil {
		// fallback to Default language
		return Default
	}

	lc, out := Parse(tag)
	if out.IsError() {
		// fallback to Default language
		return Default
	}

	return lc
}
