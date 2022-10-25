package system

import (
	"github.com/dariubs/percent"
	"github.com/dustin/go-humanize"
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
	result.Total = humanize.IBytes(d.Total)
	result.Partitions = strconv.Itoa(len(p)) + " partitions"
	return result
}

func (s *System) liveDisk() {
	d, err := disk.Usage("/")
	if err != nil {
		return
	}
	s.Live.Disk.Value = humanize.IBytes(d.Used)
	s.Live.Disk.Percentage = math.RoundToEven(percent.PercentOfFloat(float64(d.Used), float64(d.Total)))
}
