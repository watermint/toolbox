package trace

import "github.com/cihub/seelog"

// Value container for sensitive information like credentials, api token, etc.
type SensitiveValue struct {
	Value interface{}
}

// Value container for privacy information like email addr, name, etc.
type PrivacyValue struct {
	Value interface{}
}

type TraceReceiver struct {
}

// Receive and parse message
// message shall be JSON formatted. If the message is not valid JSON, receiver just omit message as plain text.
//
// ## Acceptable message formats:
//
// * `{"p":"plain text"}`
// * `{"s":"sprintf like format", "a":["arg1", {"v":"arg2", "s":true}, {"v":"arg3", "s":false}, "arg4"]}`
//  - if the argument has map form; attribute `"v"` is value. `"s"`: is sensitive (true/false) flag.
func (t *TraceReceiver) ReceiveMessage(message string, level seelog.LogLevel, context seelog.LogContextInterface) error {
	return nil
}

func (t *TraceReceiver) AfterParse(args seelog.CustomReceiverInitArgs) error {
	return nil
}

func (t *TraceReceiver) Flush() {
}

func (t *TraceReceiver) Close() error {
	return nil
}
