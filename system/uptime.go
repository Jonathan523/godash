package system

import (
	"github.com/shirou/gopsutil/v3/host"
)

func (s *System) uptime() {
	i, err := host.Info()
	if err != nil {
		return
	}
	s.CurrentSystem.Live.Uptime.Days = i.Uptime / 84600
	s.CurrentSystem.Live.Uptime.Hours = uint16((i.Uptime % 86400) / 3600)
	s.CurrentSystem.Live.Uptime.Minutes = uint16(((i.Uptime % 86400) % 3600) / 60)
	s.CurrentSystem.Live.Uptime.Seconds = uint16(((i.Uptime % 86400) % 3600) % 60)
	s.CurrentSystem.Live.Uptime.Percentage = float32((s.CurrentSystem.Live.Uptime.Minutes*100)+s.CurrentSystem.Live.Uptime.Seconds) / 60
}
