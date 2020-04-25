package dbx_async

type Async interface {
	Param(p interface{}) Async
	Status(endpoint string) Async
	PollInterval(second int) Async
	Call() (res Response, err error)
}
