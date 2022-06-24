package app_budget

import (
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/log/esl_rotate"
	"math"
)

const (
	unlimitedChunkSize = math.MaxInt64 // no rotate
	unlimitedQuota     = esl_rotate.UnlimitedQuota
	unlimitedNumBackup = esl_rotate.UnlimitedBackups
	standardChunkSize  = 100 * 1048576  // 100MiB
	standardQuota      = 5 * 1073741824 // 5GiB * 2 = 10GiB
	standardNumBackup  = standardQuota * 2 / standardChunkSize
	lowChunkSize       = 10 * 1048576  // 10MiB
	lowQuota           = 500 * 1048576 // 500MiB * 2 = 1GiB
	lowNumBackup       = lowQuota * 2 / lowChunkSize

	BudgetLow       Budget = "low"
	BudgetNormal    Budget = "normal"
	BudgetUnlimited Budget = "unlimited"
)

type Budget string

var (
	StorageBudgets = []string{
		string(BudgetLow), string(BudgetNormal), string(BudgetUnlimited),
	}
	DefaultBudget = BudgetNormal
)

func StorageBudget(budget Budget) (chunkSize, quota int64, numBackup int) {
	switch budget {
	case BudgetLow:
		return lowChunkSize, lowQuota, lowNumBackup
	case BudgetNormal:
		return standardChunkSize, standardQuota, standardNumBackup
	case BudgetUnlimited:
		return unlimitedChunkSize, unlimitedQuota, unlimitedNumBackup
	default:
		l := esl.Default()
		l.Error("Unsupported budget type, fallback to BudgetNormal", esl.String("budget", string(budget)))
		return StorageBudget(BudgetNormal)
	}
}
