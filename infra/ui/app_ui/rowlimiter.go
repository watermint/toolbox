package app_ui

const (
	consoleNumRowsThreshold = 500
)

type RowLimiter interface {
	Row(f func())
	Flush()
}

func NewRowLimiter(sy Syntax, name string) RowLimiter {
	return &rowLimitImpl{
		sy:      sy,
		name:    name,
		numRows: 0,
	}
}

type rowLimitImpl struct {
	sy      Syntax
	name    string
	numRows int
}

func (z *rowLimitImpl) Row(f func()) {
	z.numRows++
	if z.numRows <= consoleNumRowsThreshold {
		f()
	}
	if z.numRows%consoleNumRowsThreshold == 0 {
		z.sy.Info(MConsole.Progress.
			With("Label", z.name).
			With("Progress", z.numRows))
	}
}

func (z *rowLimitImpl) Flush() {
	if z.numRows >= consoleNumRowsThreshold {
		z.sy.Info(MConsole.LargeReport.
			With("Label", z.name).
			With("Num", z.numRows))
	}
}
