package api_recipe_flow

func OnRow(filePath string, validate func(cols []string) error, exec func(cols []string) error) error {
	return exec([]string{})
}

func IsErrorPrefix(prefix string, err error) bool {
	return false
}
