package app_report

type Report interface {
	Row(row interface{})
	Flush()
	Close()
}

type State func() StateContent

type StateContent struct {
	Kind   string
	Reason string
}

func Success() State {
	return func() StateContent {
		return StateContent{
			Kind: "success",
		}
	}
}
func Failure(reason string) State {
	return func() StateContent {
		return StateContent{
			Kind:   "failure",
			Reason: reason,
		}
	}
}
func Skip(reason string) State {
	return func() StateContent {
		return StateContent{
			Kind:   "skip",
			Reason: reason,
		}
	}
}

func TransactionHeader(input interface{}, result interface{}) TransactionRow {
	return TransactionRow{
		Input:  input,
		Result: result,
	}
}

func Transaction(state State, input interface{}, result interface{}) TransactionRow {
	s := state()
	return TransactionRow{
		Status: s.Kind,
		Reason: s.Reason,
		Input:  input,
		Result: result,
	}
}

type TransactionRow struct {
	Status string
	Reason string
	Input  interface{}
	Result interface{}
}

const (
	reportPath = "reports"
)
