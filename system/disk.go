package system

import (
	"github.com/dariubs/percent"
	"github.com/dustin/go-humanize"
	"github.com/shirou/gopsutil/v3/disk"
	"math"
	"strconv"
	"strings"
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
	used := d.Used
	s.Live.Disk.Value = humanize.IBytes(d.Used)
	if strings.HasSuffix(s.Live.Disk.Value, strings.Split(s.Static.Disk.Total, " ")[1]) {
		s.Live.Disk.Value = strings.Split(s.Live.Disk.Value, " ")[0]
	}
	s.Live.Disk.Percentage = math.RoundToEven(percent.PercentOfFloat(float64(used), float64(d.Total)))
}