package mo_warning

type Warning struct {
	WarningMsg  string `json:"warning_msg" path:"warning_msg"`
	WarningName string `json:"warning_name" path:"warning_name"`
}
