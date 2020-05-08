package esl

import (
	zapuber "go.uber.org/zap"
	"time"
)

type zapField struct {
	zf zapuber.Field
}

func (z zapField) zap() zapuber.Field {
	return z.zf
}

func Any(key string, val interface{}) Field {
	return zapField{zf: zapuber.Any(key, val)}
}
func Binary(key string, val []byte) Field {
	return zapField{zf: zapuber.Binary(key, val)}
}
func Bool(key string, val bool) Field {
	return zapField{zf: zapuber.Bool(key, val)}
}
func ByteString(key string, val []byte) Field {
	return zapField{zf: zapuber.ByteString(key, val)}
}
func Duration(key string, val time.Duration) Field {
	return zapField{zf: zapuber.Duration(key, val)}
}
func Error(err error) Field {
	return zapField{zf: zapuber.Error(err)}
}
func Errors(key string, val []error) Field {
	return zapField{zf: zapuber.Errors(key, val)}
}
func Float32(key string, val float32) Field {
	return zapField{zf: zapuber.Float32(key, val)}
}
func Float64(key string, val float64) Field {
	return zapField{zf: zapuber.Float64(key, val)}
}
func Int(key string, val int) Field {
	return zapField{zf: zapuber.Int(key, val)}
}
func Int16(key string, val int16) Field {
	return zapField{zf: zapuber.Int16(key, val)}
}
func Int32(key string, val int32) Field {
	return zapField{zf: zapuber.Int32(key, val)}
}
func Int64(key string, val int64) Field {
	return zapField{zf: zapuber.Int64(key, val)}
}
func Int8(key string, val int8) Field {
	return zapField{zf: zapuber.Int8(key, val)}
}
func String(key string, val string) Field {
	return zapField{zf: zapuber.String(key, val)}
}
func Strings(key string, val []string) Field {
	return zapField{zf: zapuber.Strings(key, val)}
}
func Time(key string, val time.Time) Field {
	return zapField{zf: zapuber.Time(key, val)}
}
func Uint(key string, val uint) Field {
	return zapField{zf: zapuber.Uint(key, val)}
}
func Uint16(key string, val uint16) Field {
	return zapField{zf: zapuber.Uint16(key, val)}
}
func Uint32(key string, val uint32) Field {
	return zapField{zf: zapuber.Uint32(key, val)}
}
func Uint64(key string, val uint64) Field {
	return zapField{zf: zapuber.Uint64(key, val)}
}
func Uint8(key string, val uint8) Field {
	return zapField{zf: zapuber.Uint8(key, val)}
}
