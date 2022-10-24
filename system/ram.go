package system

import (
	"github.com/dariubs/percent"
	"github.com/dustin/go-humanize"
	"github.com/shirou/gopsutil/v3/mem"
	"math"
	"strings"
)

func staticRam() Ram {
	var result = Ram{}
	r, err := mem.VirtualMemory()
	if err != nil {
		return result
	}
	result.Total = humanize.IBytes(r.Total)
	if r.SwapTotal > 0 {
		result.Swap = humanize.IBytes(r.SwapTotal) + " swap"
	} else {
		result.Swap = "No swap"
	}
	return result
}

func (s *System) liveRam() {
	r, err := mem.VirtualMemory()
	if err != nil {
		return
	}
	used := r.Used
	s.Live.Ram.Value = humanize.IBytes(r.Used)
	if strings.HasSuffix(s.Live.Ram.Value, strings.Split(s.Static.Ram.Total, " ")[1]) {
		s.Live.Ram.Value = strings.Split(s.Live.Ram.Value, " ")[0]
	}
	s.Live.Ram.Percentage = math.RoundToEven(percent.PercentOfFloat(float64(used), float64(r.Total)))
}
