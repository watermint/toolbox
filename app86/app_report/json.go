package app_report

type Json struct {
}

func (z *Json) Row(row interface{}) {
	panic("implement me")
}

func (z *Json) Transaction(kind State, row interface{}, result interface{}) {
	panic("implement me")
}

func (z *Json) Flush() {
	panic("implement me")
}

func (z *Json) Close() {
	panic("implement me")
}
