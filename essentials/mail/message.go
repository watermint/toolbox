package mail

const (
	// RFC 2822 2.1: Line separator CRLF
	lineSeparator = "\r\n"

	// RFC 2822 2.1.1: Line Length Limits:
	// Each line of characters MUST be no more than 998 characters, and SHOULD be no more than 78 characters, excluding the CRLF.

	maxLineNumOfChars    = 78
	maxLineLogicalLength = 998

	// RFC 2822 2.2: Header Fields
	headerSeparator = ": "

	// RFC 2822 2.2.3: Long Header Fields
	headerContinue = "\t"

	// RFC 2822 3.6.2: Originator fields
	headerFrom    = "From"
	headerSender  = "Sender"
	headerReplyTo = "Reply-To"

	// RFC 2822 3.6.3; Destination address fields

	headerTo  = "To"
	headerCc  = "Cc"
	headerBcc = "Bcc"

	// RFC 2822 3.6.5: Informational fields

	headerSubject = "Subject"

	// RFC 2822: 3.1 Date and Time Specification

	headerDate       = "Date"
	headerDateFormat = "Mon, 02 Jan 2006 15:04:05 -0700"
)

type Header interface {
	Name() string
	Value() string
}
