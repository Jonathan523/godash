package system

import (
	"github.com/shirou/gopsutil/v3/host"
)

func (s *System) uptime() {
	i, err := host.Info()
	if err != nil {
		return
	}
	s.Live.Uptime.Days = i.Uptime / 84600
	s.Live.Uptime.Hours = uint8((i.Uptime % 86400) / 3600)
	s.Live.Uptime.Minutes = uint8(((i.Uptime % 86400) % 3600) / 60)
	s.Live.Uptime.Seconds = uint8(((i.Uptime % 86400) % 3600) % 60)
	s.Live.Uptime.Percentage = float32(s.Live.Uptime.Hours) / 24 * 100
}
