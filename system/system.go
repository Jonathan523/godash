package system

import (
	"go.uber.org/zap"
	"time"
)

func NewSystemService(logging *zap.SugaredLogger) *System {
	s := System{log: logging}
	s.Initialize()
	return &s
}

func (s *System) UpdateLiveInformation() {
	for {
		s.liveCpu()
		s.liveRam()
		s.liveDisk()
		s.uptime()
		time.Sleep(1 * time.Second)
	}
}

func (s *System) Initialize() {
	s.Static.Host = staticHost()
	s.Static.CPU = staticCpu()
	s.Static.Ram = staticRam()
	s.Static.Disk = staticDisk()
	go s.UpdateLiveInformation()
	s.log.Debugw("system updated", "cpu", s.Static.CPU.Name, "arch", s.Static.Host.Architecture)
}
