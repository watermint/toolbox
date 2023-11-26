package eimage

import "github.com/watermint/toolbox/essentials/islet/eidiom/eoutcome"

const (
	exportOutcomeSuccess = iota
	exportOutcomeUnsupportedFormat
	exportOutcomeEncodeFailure
	exportOutcomeWriteFailure
)

func NewExportOutcomeSuccess() ExportOutcome {
	return &imgExportOutcomeImpl{
		OutcomeBase: eoutcome.NewOutcomeBaseOk(),
		reason:      exportOutcomeSuccess,
	}
}
func NewExportOutcomeUnsupportedFormat(given ImageFormat) ExportOutcome {
	return &imgExportOutcomeImpl{
		OutcomeBase: eoutcome.NewOutcomeBaseWithErrMessage("unsupported format %d", given),
		reason:      exportOutcomeUnsupportedFormat,
	}
}
func NewExportOutcomeEncodeFailure(err error) ExportOutcome {
	return &imgExportOutcomeImpl{
		OutcomeBase: eoutcome.NewOutcomeBaseError(err),
		reason:      exportOutcomeEncodeFailure,
	}
}
func NewExportOutcomeWriteFailure(err error) ExportOutcome {
	return &imgExportOutcomeImpl{
		OutcomeBase: eoutcome.NewOutcomeBaseError(err),
		reason:      exportOutcomeWriteFailure,
	}
}

type imgExportOutcomeImpl struct {
	eoutcome.OutcomeBase
	reason int
}

func (z imgExportOutcomeImpl) IsUnsupportedFormat() bool {
	return z.reason == exportOutcomeUnsupportedFormat
}

func (z imgExportOutcomeImpl) IsEncodeFailure() bool {
	return z.reason == exportOutcomeEncodeFailure
}

func (z imgExportOutcomeImpl) IsWriteFailure() bool {
	return z.reason == exportOutcomeWriteFailure
}
