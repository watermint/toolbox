package empsc

import "github.com/watermint/toolbox/essentials/islet/eidiom"

// Producer of Multi producer, single consumer
type Producer interface {
	// Send a value to the consumer. This func will block if a channel is full.
	Send(v interface{})
}

type Consumer func(v interface{}) (err error)

type Channel interface {
	eidiom.Closer

	// Producer Create a new producer
	Producer() Producer
}
