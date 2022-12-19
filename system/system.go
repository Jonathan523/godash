package system

import (
	"go.uber.org/zap"
	"godash/hub"
	"time"
)

func NewSystemService(enabled bool, logging *zap.SugaredLogger, hub *hub.Hub) *System {
	var s System
	if enabled {
		s = System{log: logging, hub: hub}
		s.Initialize()
	}
	return &s
}

func (s *System) UpdateLiveInformation() {
	for {
		s.liveCpu()
		s.liveRam()
		s.liveDisk()
		s.uptime()
		s.hub.LiveInformationCh <- hub.Message{WsType: hub.System, Message: s.CurrentSystem.Live}
		time.Sleep(1 * time.Second)
	}
}

func (s *System) Initialize() {
	s.CurrentSystem.Static.Host = staticHost()
	s.CurrentSystem.Static.CPU = staticCpu()
	s.CurrentSystem.Static.Ram = staticRam()
	s.CurrentSystem.Static.Disk = staticDisk()
	go s.UpdateLiveInformation()
	s.log.Debugw("system updated", "cpu", s.CurrentSystem.Static.CPU.Name, "arch", s.CurrentSystem.Static.Host.Architecture)
}
