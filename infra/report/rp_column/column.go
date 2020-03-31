package rp_column

type Column interface {
	Header() []string
	Values(r interface{}) []interface{}
	ValueStrings(r interface{}) []string
}
