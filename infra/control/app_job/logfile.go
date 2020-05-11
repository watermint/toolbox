package app_job

import "io"

type LogFile interface {
	// Log file type
	Type() LogFileType

	// Name of the file (not including path)
	Name() string

	// Path to the log.
	Path() string

	// True when the file is compressed.
	IsCompressed() bool

	// Read & copy to the stream
	CopyTo(writer io.Writer) error
}

type LogFileType string

const (
	LogFileTypeToolbox LogFileType = "toolbox"
	LogFileTypeCapture LogFileType = "capture"
	LogFileTypeSummary LogFileType = "summary"
	LogFileTypeRecipe  LogFileType = "recipe"
	LogFileTypeResult  LogFileType = "result"
)

var (
	LogFileTypes = []string{
		string(LogFileTypeToolbox),
		string(LogFileTypeCapture),
		string(LogFileTypeSummary),
		string(LogFileTypeRecipe),
		string(LogFileTypeResult),
	}
)
