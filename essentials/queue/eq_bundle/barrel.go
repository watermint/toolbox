package eq_bundle

import (
	"encoding/json"
	"github.com/watermint/toolbox/essentials/log/esl"
	"time"
)

func NewBarrel(mouldId, batchId string, d []byte) Barrel {
	return Barrel{
		MouldId: mouldId,
		BatchId: batchId,
		Time:    time.Now().Unix(),
		D:       d,
	}
}

// Deserialize from byte sequence.
func FromBytes(b []byte) (barrel Barrel, err error) {
	err = json.Unmarshal(b, &barrel)
	return
}

type Barrel struct {
	MouldId string `json:"mould_id"`
	BatchId string `json:"batch_id"`
	Time    int64  `json:"time"`
	D       []byte `json:"d"`
}

func (z Barrel) BarrelBatch() string {
	b := make(map[string]string)
	b["mouldId"] = z.MouldId
	b["batchId"] = z.BatchId
	bid, err := json.Marshal(b)
	if err != nil {
		l := esl.Default()
		l.Error("Unable to marshal", esl.Error(err))
		panic(err)
	}
	return string(bid)
}

func (z Barrel) ToBytes() (d []byte) {
	d, err := json.Marshal(z)
	if err != nil {
		l := esl.Default()
		l.Error("Unable to marshal", esl.Error(err))
		panic(err)
	}
	return d
}
