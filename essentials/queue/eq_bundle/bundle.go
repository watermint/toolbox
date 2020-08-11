package eq_bundle

import (
	"encoding/json"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/queue/eq_pipe"
	"time"
)

// Bundle interface will not return error.
// If there any error happens, the impl. should recover from the error themself or
// raise panic() to tell critical issue to the caller.
type Bundle interface {
	// Enqueue data with batchId
	Enqueue(d Data)

	// Fetch data from the queue.
	Fetch() (d Data, found bool)

	// Mark data as completed
	Complete(d Data)

	// Queue storage size per batchId
	Size() (sizes map[string]int, total int)

	// Close this bundle.
	Close()

	// Preserve this bundle.
	Preserve() (session Session, err error)
}

type Session struct {
	Pipes      map[string]eq_pipe.SessionId `json:"pipes"`
	InProgress eq_pipe.SessionId            `json:"in_progress"`
}

func NewData(batchId string, d []byte) Data {
	return Data{
		BatchId: batchId,
		Time:    time.Now().Unix(),
		D:       d,
	}
}

type Data struct {
	BatchId string `json:"batch_id"`
	Time    int64  `json:"time"`
	D       []byte `json:"d"`
}

// Serialized to byte sequence
func (z Data) ToBytes() (d []byte) {
	d, err := json.Marshal(z)
	if err != nil {
		l := esl.Default()
		l.Error("Unable to marshal", esl.Error(err))
		panic(err)
	}
	return d
}

// Deserialize from byte sequence.
func FromBytes(b []byte) (d Data, err error) {
	err = json.Unmarshal(b, &d)
	return
}
