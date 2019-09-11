package app_log

import "go.uber.org/zap"

type Record interface {
	Level() Level
	Message() string
	Fields() []Field
	ZapFields() []zap.Field
}

type Level interface {
	Out(l *zap.Logger, r Record)
}

type Field interface {
	ZapField() zap.Field
}

type Log func() Record

type debugImpl struct {
}

func (z *debugImpl) Out(l *zap.Logger, r Record) {
	l.Debug(r.Message(), r.ZapFields()...)
}

type recordImpl struct {
	level   Level
	message string
	fields  []Field
}

func (z *recordImpl) ZapFields() []zap.Field {
	zf := make([]zap.Field, len(z.fields))
	for i, f := range z.fields {
		zf[i] = f.ZapField()
	}
	return zf
}

func (z *recordImpl) Level() Level {
	return z.level
}

func (z *recordImpl) Message() string {
	return z.message
}

func (z *recordImpl) Fields() []Field {
	return z.fields
}

func Debug(message string, fields ...Field) Log {
	return func() Record {
		return &recordImpl{
			level:   &debugImpl{},
			message: message,
			fields:  fields,
		}
	}
}
