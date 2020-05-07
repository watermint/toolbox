package dbx_async

type Async interface {
	Call(opts ...AsyncOpt) Response
}

type AsyncOpts struct {
	PollInterval   int
	StatusEndpoint string
}

type AsyncOpt func(o AsyncOpts) AsyncOpts

func PollInterval(seconds int) AsyncOpt {
	return func(o AsyncOpts) AsyncOpts {
		o.PollInterval = seconds
		return o
	}
}
func Status(endpoint string) AsyncOpt {
	return func(o AsyncOpts) AsyncOpts {
		o.StatusEndpoint = endpoint
		return o
	}
}

func Combined(opts []AsyncOpt) AsyncOpts {
	ao := AsyncOpts{}
	for _, o := range opts {
		ao = o(ao)
	}
	return ao
}
