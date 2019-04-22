package api_recipe_flow

func EachRow(filePath string, validate func(cols []string) error, exec func(cols []string) error) error {
	return exec([]string{})
}

func IsErrorPrefix(prefix string, err error) bool {
	return false
}

type RowDataFile interface {
	EachRow(validate func(cols []string) error, exec func(cols []string) error) error
}
