package app_report

type Report interface {
	Row(row interface{})
	Transaction(state State, input interface{}, result interface{})
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

type Transaction struct {
	State  StateContent
	Input  interface{}
	Result interface{}
}

const (
	reportPath = "reports"
)
