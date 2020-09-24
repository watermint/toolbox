package em_file_random

import "time"

type Opts struct {
	fileSizeRangeMin      int64
	fileSizeRangeMax      int64
	fileDateRangeMin      time.Time
	fileDateRangeMax      time.Time
	numDescendantRangeMin int
	numDescendantLambda   int
	numDescendantRangeMax int
	numFiles              int
	depthRangeMax         int
	seed                  int64
}

func (z Opts) Apply(opt []Opt) Opts {
	switch len(opt) {
	case 0:
		return z
	case 1:
		return opt[0](z)
	default:
		return opt[0](z).Apply(opt[1:])
	}
}

type Opt func(o Opts) Opts

func Default() Opts {
	return Opts{
		fileSizeRangeMax:      2048,
		fileSizeRangeMin:      0,
		fileDateRangeMin:      time.Now().Add(-2 * 365 * 24 * time.Hour),
		fileDateRangeMax:      time.Now(),
		numDescendantRangeMin: 0,
		numDescendantRangeMax: 1 << 15,
		numDescendantLambda:   8,
		numFiles:              1_000,
		depthRangeMax:         8,
		seed:                  time.Now().UnixNano(),
	}
}

func FileSize(min, max int64) Opt {
	return func(o Opts) Opts {
		o.fileSizeRangeMin = min
		o.fileSizeRangeMax = max
		return o
	}
}

func FileDate(min, max time.Time) Opt {
	return func(o Opts) Opts {
		o.fileDateRangeMin = min
		o.fileDateRangeMax = max
		return o
	}
}

func Depth(max int) Opt {
	return func(o Opts) Opts {
		o.depthRangeMax = max
		return o
	}
}

func NumDescendant(min, max int) Opt {
	return func(o Opts) Opts {
		o.numDescendantRangeMin = min
		o.numDescendantRangeMax = max
		return o
	}
}

func NumFiles(files int) Opt {
	return func(o Opts) Opts {
		o.numFiles = files
		return o
	}
}

func Seed(seed int64) Opt {
	return func(o Opts) Opts {
		o.seed = seed
		return o
	}
}
