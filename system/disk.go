package system

import (
	"github.com/dariubs/percent"
	"github.com/dustin/go-humanize"
	"github.com/shirou/gopsutil/v3/disk"
	"math"
)

func staticDisk() string {
	d, err := disk.Usage("/")
	if err != nil {
		return ""
	}
	return humanize.IBytes(d.Total)
}

func (s *System) liveDisk() {
	d, err := disk.Usage("/")
	if err != nil {
		return
	}
	used := d.Used
	total := d.Total
	s.Live.Disk.Value = humanize.IBytes(d.Used)
	s.Live.Disk.Percentage = math.RoundToEven(percent.PercentOfFloat(float64(used), float64(total)))
}
