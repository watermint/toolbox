package esl

import (
	zapuber "go.uber.org/zap" // change package alias for prevent text replace on migration
	"testing"
	"time"
)

func TestAny(t *testing.T) {
	z := Any("key", 123)
	y := z.zap()
	if !y.Equals(zapuber.Any("key", 123)) {
		t.Error(z)
	}
}
func TestBinary(t *testing.T) {
	content := []byte("hello")
	z := Binary("key", content)
	y := z.zap()
	if !y.Equals(zapuber.Binary("key", content)) {
		t.Error(z)
	}
}
func TestBool(t *testing.T) {
	z := Bool("key", true)
	y := z.zap()
	if !y.Equals(zapuber.Bool("key", true)) {
		t.Error(z)
	}
}
func TestByteString(t *testing.T) {
	content := []byte("hello")
	z := ByteString("key", content)
	y := z.zap()
	if !y.Equals(zapuber.ByteString("key", content)) {
		t.Error(z)
	}
}
func TestDuration(t *testing.T) {
	z := Duration("key", 123)
	y := z.zap()
	if !y.Equals(zapuber.Duration("key", 123)) {
		t.Error(z)
	}
}
func TestFloat32(t *testing.T) {
	z := Float32("key", 123)
	y := z.zap()
	if !y.Equals(zapuber.Float32("key", 123)) {
		t.Error(z)
	}
}
func TestFloat64(t *testing.T) {
	z := Float64("key", 123)
	y := z.zap()
	if !y.Equals(zapuber.Float64("key", 123)) {
		t.Error(z)
	}
}
func TestInt(t *testing.T) {
	z := Int("key", 123)
	y := z.zap()
	if !y.Equals(zapuber.Int("key", 123)) {
		t.Error(z)
	}
}
func TestInt16(t *testing.T) {
	z := Int16("key", 123)
	y := z.zap()
	if !y.Equals(zapuber.Int16("key", 123)) {
		t.Error(z)
	}
}
func TestInt32(t *testing.T) {
	z := Int32("key", 123)
	y := z.zap()
	if !y.Equals(zapuber.Int32("key", 123)) {
		t.Error(z)
	}
}
func TestInt64(t *testing.T) {
	z := Int64("key", 123)
	y := z.zap()
	if !y.Equals(zapuber.Int64("key", 123)) {
		t.Error(z)
	}
}
func TestInt8(t *testing.T) {
	z := Int8("key", 123)
	y := z.zap()
	if !y.Equals(zapuber.Int8("key", 123)) {
		t.Error(z)
	}
}
func TestString(t *testing.T) {
	z := String("key", "hello")
	y := z.zap()
	if !y.Equals(zapuber.String("key", "hello")) {
		t.Error(z)
	}
}
func TestStrings(t *testing.T) {
	content := []string{"Hello", "World"}
	z := Strings("key", content)
	y := z.zap()
	if !y.Equals(zapuber.Strings("key", content)) {
		t.Error(z)
	}
}
func TestTime(t *testing.T) {
	content := time.Now()
	contentFormatted := content.Format(time.RFC3339)
	z := Time("key", content)
	y := z.zap()
	if !y.Equals(zapuber.String("key", contentFormatted)) {
		t.Error(z)
	}
}
func TestUint(t *testing.T) {
	z := Uint("key", 123)
	y := z.zap()
	if !y.Equals(zapuber.Uint("key", 123)) {
		t.Error(z)
	}
}
func TestUint16(t *testing.T) {
	z := Uint16("key", 123)
	y := z.zap()
	if !y.Equals(zapuber.Uint16("key", 123)) {
		t.Error(z)
	}
}
func TestUint32(t *testing.T) {
	z := Uint32("key", 123)
	y := z.zap()
	if !y.Equals(zapuber.Uint32("key", 123)) {
		t.Error(z)
	}
}
func TestUint64(t *testing.T) {
	z := Uint64("key", 123)
	y := z.zap()
	if !y.Equals(zapuber.Uint64("key", 123)) {
		t.Error(z)
	}
}
func TestUint8(t *testing.T) {
	z := Uint8("key", 123)
	y := z.zap()
	if !y.Equals(zapuber.Uint8("key", 123)) {
		t.Error(z)
	}
}
