package eq_progress

type Progress interface {
	OnEnqueue(mouldId, batchId string, completed, total int)
	OnComplete(mouldId, batchId string, completed, total int)
}
