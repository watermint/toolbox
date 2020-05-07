package lang

// Returns language and priority.
// Returns an array [TargetLanguage, DefaultLanguage] if a target language is default language.
// Otherwise returns [TargetLanguage].
func Priority(target Lang) []Lang {
	if target.IsDefault() {
		return []Lang{target}
	}
	return []Lang{target, Default}
}
