package app_budget

import (
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/log/esl_rotate"
)

const (
	logChunkSize       = 100 * 1048576 // 100MiB
	unlimitedQuota     = esl_rotate.UnlimitedQuota
	unlimitedNumBackup = esl_rotate.UnlimitedBackups
	standardQuota      = 5 * 1073741824 // 5GiB * 2 = 10GiB
	standardNumBackup  = standardQuota * 2 / logChunkSize
	lowQuota           = 500 * 1048576 // 500MiB * 2 = 1GiB
	lowNumBackup       = lowQuota * 2 / logChunkSize

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
		return logChunkSize, lowQuota, lowNumBackup
	case BudgetNormal:
		return logChunkSize, standardQuota, standardNumBackup
	case BudgetUnlimited:
		return logChunkSize, unlimitedQuota, unlimitedNumBackup
	default:
		l := esl.Default()
		l.Error("Unsupported budget type, fallback to BudgetNormal", esl.String("budget", string(budget)))
		return StorageBudget(BudgetNormal)
	}
}
