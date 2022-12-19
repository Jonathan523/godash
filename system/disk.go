package system

import (
	"github.com/dariubs/percent"
	"github.com/shirou/gopsutil/v3/disk"
	"math"
	"strconv"
)

func staticDisk() Disk {
	var result = Disk{}
	d, err := disk.Usage("/")
	if err != nil {
		return result
	}
	p, err := disk.Partitions(false)
	result.Total = readableSize(d.Total)
	result.Partitions = strconv.Itoa(len(p)) + " partitions"
	return result
}

func (s *System) liveDisk() {
	d, err := disk.Usage("/")
	if err != nil {
		return
	}
	s.CurrentSystem.Live.Disk.Value = readableSize(d.Used)
	s.CurrentSystem.Live.Disk.Percentage = math.RoundToEven(percent.PercentOfFloat(float64(d.Used), float64(d.Total)))
}
