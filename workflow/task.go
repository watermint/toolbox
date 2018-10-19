package workflow

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"github.com/cihub/seelog"
	"github.com/syndtr/goleveldb/leveldb/iterator"
)

type TaskIterator struct {
	iter iterator.Iterator
}

func (t *TaskIterator) Next() bool {
	return t.iter.Next()
}
func (t *TaskIterator) Prev() bool {
	return t.iter.Prev()
}
func (t *TaskIterator) First() bool {
	return t.iter.First()
}
func (t *TaskIterator) Last() bool {
	return t.iter.Last()
}
func (t *TaskIterator) Release() {
	t.iter.Release()
}
func (t *TaskIterator) Error() error {
	return t.iter.Error()
}
func (t *TaskIterator) Key() (state, prefix, taskId string) {
	return parseKey(t.iter.Key())
}
func (t *TaskIterator) Value() []byte {
	return t.iter.Value()
}
func (t *TaskIterator) Task() (state string, task *Task) {
	return NewTaskFromKeyAndValue(t.iter.Key(), t.Value())
}

type TaskValueContainer struct {
	DeferUntil int64
	Context    []byte
}

func (t *TaskValueContainer) Value() []byte {
	buf := bytes.NewBuffer(nil)
	err := gob.NewEncoder(buf).Encode(t)
	if err != nil {
		seelog.Errorf("Unable to decode value: error[%s]", err)
		panic("Unable to encode value")
	}
	return buf.Bytes()
}

func NewTaskValueContainer(task *Task) *TaskValueContainer {
	return &TaskValueContainer{
		DeferUntil: task.DeferUntil,
		Context:    task.Context,
	}
}
func NewTaskFromKeyAndValue(key []byte, value []byte) (state string, task *Task) {
	state, taskPrefix, taskId := parseKey(key)
	buf := bytes.NewBuffer(value)
	tvc := TaskValueContainer{}
	err := gob.NewDecoder(buf).Decode(&tvc)
	if err != nil {
		seelog.Errorf("Unable to decode value: error[%s]", err)
		panic("Unable to decode value")
	}
	task = &Task{
		TaskPrefix: taskPrefix,
		TaskId:     taskId,
		Context:    tvc.Context,
		DeferUntil: tvc.DeferUntil,
	}
	return
}

func UnmarshalContext(t *Task, c interface{}) {
	err := json.Unmarshal(t.Context, c)
	if err != nil {
		seelog.Errorf("Unable to unmarshal context: TaskPrefix[%s]", t.TaskPrefix)
		panic("Unable to unmarshal context")
	}
}
func MarshalTask(prefix, taskId string, context interface{}) *Task {
	c, err := json.Marshal(context)
	if err != nil {
		seelog.Errorf("Unable to unmarshal context: TaskPrefix[%s]", prefix)
		panic("Unable to unmarshal context")
	}
	return &Task{
		TaskPrefix: prefix,
		TaskId:     taskId,
		Context:    c,
	}
}
