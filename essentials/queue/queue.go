package queue

import (
	"encoding/json"
	"github.com/watermint/toolbox/essentials/log/esl"
	"reflect"
)

type FlowError interface {
	error

	// True when the error is obviously retryable.
	// Otherwise it's not retriable, or difficult to determine.
	Retryable() bool
}

type Queue interface {
	Enqueue(p interface{})
	Dequeue()
	Batch(batchId string) Queue
}

func NewQueue(f interface{}, ctx ...interface{}) Queue {
	l := esl.Default()

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
		if !ctxType.AssignableTo(argCtxType) {
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

	return &queueImpl{
		ctx:           ctx,
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

type queueImpl struct {
	ctx  []interface{}
	last []byte

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

func (z *queueImpl) Batch(batchId string) Queue {
	return &queueImpl{
		ctx:           z.ctx,
		batchId:       batchId,
		handler:       z.handler,
		handlerType:   z.handlerType,
		handlerValue:  z.handlerValue,
		paramType:     z.paramType,
		paramTypeOrig: z.paramTypeOrig,
		paramTypeKind: z.paramTypeKind,
		paramIsPtr:    z.paramIsPtr,
		hasErrorOut:   z.hasErrorOut,
	}
}

// p is the execution parameter.
// The value must be serializable into JSON format.
func (z *queueImpl) Enqueue(p interface{}) {
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

	z.last = msg
}

func (z *queueImpl) Dequeue() {
	l := esl.Default()
	if z.last == nil {
		l.Info("Empty queue")
		return
	}

	p := reflect.New(z.paramType).Interface()

	if err := json.Unmarshal(z.last, p); err != nil {
		l.Debug("Unable to unmarshal", esl.Error(err))
		panic(err)
	}

	l.Debug("param after unmarshal", esl.Any("p", p))

	v := reflect.ValueOf(p)
	if !z.paramIsPtr {
		v = v.Elem()
	}

	params := make([]reflect.Value, 0)
	params = append(params, v.Convert(z.paramTypeOrig))
	for _, ctx := range z.ctx {
		params = append(params, reflect.ValueOf(ctx))
	}

	out := z.handlerValue.Call(params)
	if z.hasErrorOut {
		if x := len(out); x != 1 {
			l.Debug("invalid out size", esl.Int("len(out)", x))
			panic("invalid out size")
		}
	}
}
