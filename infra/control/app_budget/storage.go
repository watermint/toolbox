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
	standardChunkSize  = 200 * 1024    // 200KiB
	standardQuota      = 500 * 1048576 // 500MiB * 2 = 1000MiB
	standardNumBackup  = esl_rotate.UnlimitedBackups
	lowChunkSize       = 100 * 1024   // 100KiB
	lowQuota           = 50 * 1048576 // 5MiB * 2 =  100MiB
	lowNumBackup       = esl_rotate.UnlimitedBackups

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
