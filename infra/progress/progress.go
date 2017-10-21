package progress

import (
	"github.com/cihub/seelog"
	"github.com/gosuri/uiprogress"
	"github.com/watermint/toolbox/infra"
)

// Wrapper for gosuri/uiprogress
type ProgressUI struct {
	UI    *uiprogress.Progress
	Bar   *uiprogress.Bar
	Infra *infra.InfraOpts
}

func (p *ProgressUI) Start(cnt int) {
	seelog.Flush()
	p.UI = uiprogress.New()
	p.UI.Start()
	p.Bar = p.UI.AddBar(cnt)
	p.Bar.PrependElapsed()
	p.Bar.AppendCompleted()
}

func (p *ProgressUI) End() {
	p.UI.Stop()
}

func (p *ProgressUI) Incr() bool {
	return p.Bar.Incr()
}
