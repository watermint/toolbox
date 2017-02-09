package report

type ReportRow interface {
}

type ReportHeader struct {
	Headers []string
}

type ReportData struct {
	Data []string
}

type ReportEOF struct {
}
