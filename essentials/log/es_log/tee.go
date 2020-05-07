package es_log

func NewTee() Tee {
	return &teeImpl{
		loggers: make([]Logger, 0),
	}
}

// Tee implementation is mutable.
type Tee interface {
	Logger

	// Add subscriber. And returns self instance.
	AddSubscriber(l Logger) Tee

	// Remove subscriber.
	RemoveSubscriber(l Logger) Tee
}

type teeImpl struct {
	loggers []Logger
}

func (z teeImpl) withFunc(f func(l Logger) Logger) Logger {
	ls := make([]Logger, 0)
	for _, l := range z.loggers {
		x := f(l)
		ls = append(ls, x)
	}
	return &teeImpl{
		loggers: ls,
	}
}

func (z teeImpl) each(f func(l Logger)) {
	for _, l := range z.loggers {
		f(l)
	}
}

func (z teeImpl) With(fields ...Field) Logger {
	return z.withFunc(func(l Logger) Logger {
		return l.With(fields...)
	})
}

func (z teeImpl) AddCallerSkip(n int) Logger {
	return z.withFunc(func(l Logger) Logger {
		return l.AddCallerSkip(n)
	})
}

func (z teeImpl) Debug(msg string, fields ...Field) {
	z.each(func(l Logger) {
		l.Debug(msg, fields...)
	})
}

func (z teeImpl) Info(msg string, fields ...Field) {
	z.each(func(l Logger) {
		l.Info(msg, fields...)
	})
}

func (z teeImpl) Warn(msg string, fields ...Field) {
	z.each(func(l Logger) {
		l.Warn(msg, fields...)
	})
}

func (z teeImpl) Error(msg string, fields ...Field) {
	z.each(func(l Logger) {
		l.Error(msg, fields...)
	})
}

func (z teeImpl) Sync() error {
	var lastErr error
	z.each(func(l Logger) {
		if err := l.Sync(); err != nil {
			lastErr = err
		}
	})
	return lastErr
}

func (z *teeImpl) AddSubscriber(l Logger) Tee {
	x := l.AddCallerSkip(3)
	z.loggers = append(z.loggers, x)
	return z
}

func (z *teeImpl) RemoveSubscriber(l Logger) Tee {
	// currently nothing happens
	return z
}
