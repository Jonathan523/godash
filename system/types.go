package system

type SystemConfig struct {
	LiveSystem bool `mapstructure:"LIVE_SYSTEM"`
}

type LiveStorageInformation struct {
	Value      string  `json:"value"`
	Percentage float64 `json:"percentage"`
}

type LiveInformation struct {
	CPU    float64                `json:"cpu"`
	Ram    LiveStorageInformation `json:"ram"`
	Disk   LiveStorageInformation `json:"disk"`
	Uptime Uptime                 `json:"uptime"`
}

type Uptime struct {
	Days           uint64  `json:"days"`
	Hours          uint8   `json:"hours"`
	Minutes        uint8   `json:"minutes"`
	Seconds        uint8   `json:"seconds"`
	SecondsPercent float32 `json:"seconds_percent"`
}

type CPU struct {
	Name    string `json:"name"`
	Threads string `json:"threads"`
}

type Host struct {
	Architecture string `json:"architecture"`
	HostName     string `json:"host_name"`
}

type Ram struct {
	Total string `json:"total"`
	Swap  string `json:"swap"`
}

type Disk struct {
	Total      string `json:"total"`
	Partitions string `json:"partitions"`
}

type StaticInformation struct {
	CPU  CPU  `json:"cpu"`
	Ram  Ram  `json:"ram"`
	Disk Disk `json:"disk"`
	Host Host `json:"host"`
}

type System struct {
	Live   LiveInformation   `json:"live"`
	Static StaticInformation `json:"static"`
}
