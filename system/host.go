package system

import (
	"github.com/shirou/gopsutil/v3/host"
	"runtime"
)

func staticHost() Host {
	var h Host
	info, _ := host.Info()
	h.Architecture = runtime.GOARCH
	h.HostName = info.Hostname
	return h
}
