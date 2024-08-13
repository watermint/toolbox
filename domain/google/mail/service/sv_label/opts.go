package sv_label

type Opts struct {
	LabelListVisibility   string
	MessageListVisibility string
	ColorBackground       string
	ColorText             string
	Name                  string
}
type Opt func(o Opts) Opts

func (z Opts) Apply(opts ...Opt) Opts {
	switch len(opts) {
	case 0:
		return z
	case 1:
		return opts[0](z)
	default:
		return opts[0](z).Apply(opts[1:]...)
	}
}

func (z Opts) Param() LabelParam {
	p := LabelParam{
		Name:                  z.Name,
		LabelListVisibility:   z.LabelListVisibility,
		MessageListVisibility: z.MessageListVisibility,
	}
	if z.ColorText != "" || z.ColorBackground != "" {
		p.Color = &LabelColorParam{
			BackgroundColor: z.ColorBackground,
			TextColor:       z.ColorText,
		}
	}
	return p
}

func Name(v string) Opt {
	return func(o Opts) Opts {
		o.Name = v
		return o
	}
}

func LabelListVisibility(v string) Opt {
	return func(o Opts) Opts {
		o.LabelListVisibility = v
		return o
	}
}
func MessageListVisibility(v string) Opt {
	return func(o Opts) Opts {
		o.MessageListVisibility = v
		return o
	}
}
func ColorBackground(c string) Opt {
	return func(o Opts) Opts {
		o.ColorBackground = c
		return o
	}
}
func ColorText(c string) Opt {
	return func(o Opts) Opts {
		o.ColorText = c
		return o
	}
}
