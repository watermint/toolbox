package ea_indicator

import (
	"github.com/vbauerster/mpb/v5"
	"sync"
)

type StatusListener interface {
	Update()
}

type Status interface {
	UpdateTitle(title string)
	UpdateStatus(status string)
	CurrentTitle() string
	CurrentStatus() string
	SetListener(listener StatusListener)
}

func newStatus() Status {
	return &statusImpl{}
}

type statusImpl struct {
	listener StatusListener
	title    string
	status   string
}

func (z *statusImpl) SetListener(listener StatusListener) {
	z.listener = listener
}

func (z *statusImpl) UpdateTitle(title string) {
	z.title = title
}

func (z *statusImpl) UpdateStatus(status string) {
	z.status = status
}

func (z *statusImpl) CurrentTitle() string {
	return z.title
}

func (z *statusImpl) CurrentStatus() string {
	return z.status
}

// Status bar
type StatusBar interface {
	Status
	StatusListener
	Done()
}

func newStatusBar(name string, status Status, bar *mpb.Bar) StatusBar {
	return &statusBarImpl{
		name:   name,
		status: status,
		bar:    bar,
	}
}

type statusBarImpl struct {
	status   Status
	name     string
	progress int
	mutex    sync.Mutex
	bar      *mpb.Bar
}

func (z *statusBarImpl) SetListener(listener StatusListener) {
	// nop
}

func (z *statusBarImpl) Update() {
	z.updateBar()
}

func (z *statusBarImpl) CurrentTitle() string {
	return z.status.CurrentTitle()
}

func (z *statusBarImpl) CurrentStatus() string {
	return z.status.CurrentStatus()
}

func (z *statusBarImpl) updateBar() {
	z.mutex.Lock()
	defer z.mutex.Unlock()

	z.progress = z.progress ^ 1
	z.bar.SetCurrent(int64(z.progress))
}

func (z *statusBarImpl) UpdateTitle(title string) {
	z.status.UpdateTitle(title)
	z.updateBar()
}

func (z *statusBarImpl) UpdateStatus(status string) {
	z.status.UpdateStatus(status)
	z.updateBar()
}

func (z *statusBarImpl) Done() {
	z.bar.SetTotal(2, true)
}

func NewNopStatus() StatusBar {
	return &nopStatusImpl{}
}

type nopStatusImpl struct {
}

func (z nopStatusImpl) SetListener(listener StatusListener) {
}

func (z nopStatusImpl) Update() {
}

func (z nopStatusImpl) CurrentTitle() string {
	return ""
}

func (z nopStatusImpl) CurrentStatus() string {
	return ""
}

func (z nopStatusImpl) UpdateTitle(title string) {
}

func (z nopStatusImpl) UpdateStatus(status string) {
}

func (z nopStatusImpl) Done() {
}
