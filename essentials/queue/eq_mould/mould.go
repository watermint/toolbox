package eq_mould

import (
	"encoding/json"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/queue/eq_bundle"
	"reflect"
)

type Mould interface {
	// Enqueue data
	Pour(p interface{})

	// Process the data
	Process(d eq_bundle.Data)

	// With batchId
	Batch(batchId string) Mould
}

func New(s eq_bundle.Bundle, f interface{}, ctx ...interface{}) Mould {
	l := esl.Default()

	if s == nil {
		l.Debug("No storage")
		panic("No storage for the queue")
	}

	handlerType := reflect.TypeOf(f)
	handlerValue := reflect.ValueOf(f)
	if handlerType.Kind() != reflect.Func {
		l.Debug("f is not a func")
		panic("f is not a func")
	}
	if handlerType.NumIn() != 1+len(ctx) {
		l.Debug("f must have one + num ctx arguments")
		panic("f must have one + num ctx arguments")
	}
	paramType := handlerType.In(0)
	paramTypeOrig := paramType
	paramIsPtr := false
	if paramType.Kind() == reflect.Ptr {
		paramIsPtr = true
		paramType = paramType.Elem()
	}
	paramTypeKind := paramType.Kind()
	switch paramTypeKind {
	case reflect.Bool, reflect.Int, reflect.String, reflect.Struct:
		l.Debug("first in param have serializable type")
	default:
		l.Debug("first in param does not have serializable type", esl.Any("kind", paramTypeKind))
		panic("f param type is not serializable")
	}

	for i, c := range ctx {
		ctxType := reflect.TypeOf(c)
		argCtxType := handlerType.In(i + 1)
		if !ctxType.ConvertibleTo(argCtxType) {
			l.Debug("invalid param", esl.Int("index", i+1), esl.Any("expected", argCtxType), esl.Any("actual", ctxType))
			panic("invalid param type")
		}
	}

	var hasErrorOut bool
	switch handlerType.NumOut() {
	case 0:
		hasErrorOut = false
	case 1:
		if handlerType.Out(0).AssignableTo(reflect.TypeOf((*error)(nil)).Elem()) {
			hasErrorOut = true
		} else {
			l.Debug("f return type must be no return or error", esl.Int("numOut", handlerType.NumOut()))
			panic("f return type must be no return or error")
		}
	default:
		l.Debug("f has two or more returns", esl.Int("numOut", handlerType.NumOut()))
		panic("f has two or more returns")
	}

	return &mouldImpl{
		ctx:           ctx,
		storage:       s,
		handler:       f,
		handlerType:   handlerType,
		handlerValue:  handlerValue,
		paramType:     paramType,
		paramIsPtr:    paramIsPtr,
		paramTypeKind: paramTypeKind,
		paramTypeOrig: paramTypeOrig,
		hasErrorOut:   hasErrorOut,
	}
}

type mouldImpl struct {
	ctx     []interface{}
	storage eq_bundle.Bundle

	batchId string

	handler       interface{}
	handlerType   reflect.Type
	handlerValue  reflect.Value
	paramType     reflect.Type
	paramTypeOrig reflect.Type
	paramTypeKind reflect.Kind
	paramIsPtr    bool
	hasErrorOut   bool
}

func (z mouldImpl) Batch(batchId string) Mould {
	z.batchId = batchId
	return &z
}

// p is the execution parameter.
// The value must be serializable into JSON format.
func (z mouldImpl) Pour(p interface{}) {
	l := esl.Default()

	// validate param type
	if z.paramIsPtr {
		pt := reflect.TypeOf(p)
		if pt.Kind() != reflect.Ptr {
			l.Debug("Type incompatible")
			panic("param type incompatible")
		}
		if pt.Elem() != z.paramType {
			l.Debug("Type incompatible")
			panic("param type incompatible")
		}
	} else {
		pt := reflect.TypeOf(p)
		if pt != z.paramType {
			l.Debug("Type incompatible")
			panic("param type incompatible")
		}
	}

	msg, err := json.Marshal(p)
	if err != nil {
		l.Debug("Unable to marshal", esl.Error(err))
		panic(err)
	}

	d := eq_bundle.NewData(z.batchId, msg)
	l.Debug("Enqueue", esl.Any("Data", d))
	z.storage.Enqueue(d)
}

func (z mouldImpl) Process(d eq_bundle.Data) {
	l := esl.Default()
	p := reflect.New(z.paramType).Interface()

	if err := json.Unmarshal(d.D, p); err != nil {
		l.Debug("Unable to unmarshal", esl.Error(err))
		panic(err)
	}

	l.Debug("param after unmarshal", esl.Any("p", p), esl.String("batchId", d.BatchId))

	v := reflect.ValueOf(p)
	if !z.paramIsPtr {
		v = v.Elem()
	}

	params := make([]reflect.Value, 0)
	params = append(params, v.Convert(z.paramTypeOrig))
	for _, ctx := range z.ctx {
		params = append(params, reflect.ValueOf(ctx))
	}

	l.Debug("Call processor", esl.Int("NumParams", len(params)))
	out := z.handlerValue.Call(params)
	if z.hasErrorOut {
		// Do not verify len(out), and type of the value. That is verified on creation.
		outErr := out[0].Interface().(error)
		l.Debug("Error form the processor", esl.Error(outErr))
	}

	l.Debug("Mark as completed", esl.Any("Data", d))
	z.storage.Complete(d)
	l.Debug("Completed", esl.Any("Data", d))
}