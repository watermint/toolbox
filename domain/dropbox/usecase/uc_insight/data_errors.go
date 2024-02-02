package uc_insight

type ApiErrorRecord interface {
	// ToParam Convert to retry parameter
	ToParam() interface{}
}
