package monitor

type Event struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

const (
	EventCpuInfo       = "cpuInfo"
	EventCpuTime       = "cpuTime"
	EventHostInfo      = "hostInfo"
	EventDiskPartition = "diskPartition"
	EventDiskUsage     = "diskUsage"
	EventLoadAverage   = "loadAverage"
	EventMemoryStat    = "memStat"
)
