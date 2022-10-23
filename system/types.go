package system

type SystemConfig struct {
	LiveSystem bool `mapstructure:"LIVE_SYSTEM"`
}

type BasicSystemInformation struct {
	Value      string  `json:"value" validate:"required"`
	Percentage float64 `json:"percentage" validate:"required"`
}

type LiveInformation struct {
	CPU          CpuSystemInformation   `json:"cpu" validate:"required"`
	Ram          BasicSystemInformation `json:"ram" validate:"required"`
	Disk         BasicSystemInformation `json:"disk" validate:"required"`
	ServerUptime uint64                 `json:"server_uptime" validate:"required"`
}

type StaticInformation struct {
	CPU  CPU    `json:"cpu" validate:"required"`
	Ram  string `json:"ram" validate:"required"`
	Disk string `json:"disk" validate:"required"`
}

type System struct {
	Live   LiveInformation   `json:"live" validate:"required"`
	Static StaticInformation `json:"static" validate:"required"`
}

type CPU struct {
	Name         string `json:"name" validate:"required"`
	Threads      int    `json:"threads" validate:"required"`
	Architecture string `json:"architecture" validate:"required"`
}

type CpuSystemInformation struct {
	Percentage float64 `json:"percentage" validate:"required"`
}
