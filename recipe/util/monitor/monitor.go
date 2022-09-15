package monitor

type Event struct {
	Time string      `json:"time"`
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

const (
	EventMonitorInfo   = "monitor"
	EventCpuInfo       = "cpuInfo"
	EventCpuPercent    = "cpuPercent"
	EventCpuTime       = "cpuTime"
	EventDiskPartition = "diskPartition"
	EventDiskUsage     = "diskUsage"
	EventHostInfo      = "hostInfo"
	EventLoadAverage   = "loadAverage"
	EventMemoryStat    = "memStat"
	EventNetIO         = "netIo"
	EventNetProtocol   = "netProtocol"
)
