package system

import (
	"github.com/dariubs/percent"
	"github.com/dustin/go-humanize"
	"github.com/shirou/gopsutil/v3/mem"
	"math"
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
	total := r.Total
	s.Live.Ram.Value = humanize.IBytes(r.Used)
	s.Live.Ram.Percentage = math.RoundToEven(percent.PercentOfFloat(float64(used), float64(total)))
}
