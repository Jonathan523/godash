package system

type SystemConfig struct {
	LiveSystem bool `mapstructure:"LIVE_SYSTEM"`
}

type LiveStorageInformation struct {
	Value      string  `json:"value"`
	Percentage float64 `json:"percentage"`
}

type LiveInformation struct {
	CPU          float64                `json:"cpu"`
	Ram          LiveStorageInformation `json:"ram"`
	Disk         LiveStorageInformation `json:"disk"`
	ServerUptime uint64                 `json:"server_uptime"`
}

type CPU struct {
	Name         string `json:"name"`
	Threads      string `json:"threads"`
	Architecture string `json:"architecture"`
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
}

type System struct {
	Live   LiveInformation   `json:"live"`
	Static StaticInformation `json:"static"`
}
