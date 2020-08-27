package eq_progress

import "github.com/watermint/toolbox/essentials/queue/eq_stat"

type Progress interface {
	OnEnqueue(mouldId, batchId string, stat eq_stat.Stat)
	OnComplete(mouldId, batchId string, stat eq_stat.Stat)
}
