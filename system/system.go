package system

import (
	"github.com/sirupsen/logrus"
	"godash/config"
	"godash/hub"
	"time"
)

var Config = SystemConfig{}
var Sys = System{}

func NewSystemService() {
	config.ParseViperConfig(&Config, config.AddViperConfig("system"))
	if Config.LiveSystem {
		Sys.Initialize()
	}
}

func (s *System) UpdateLiveInformation() {
	for {
		s.liveCpu()
		s.liveRam()
		s.liveDisk()
		s.uptime()
		hub.LiveInformationCh <- hub.Message{WsType: hub.System, Message: s.Live}
		time.Sleep(1 * time.Second)
	}
}

func (s *System) Initialize() {
	s.Static.Host = staticHost()
	s.Static.CPU = staticCpu()
	s.Static.Ram = staticRam()
	s.Static.Disk = staticDisk()
	go s.UpdateLiveInformation()
	logrus.WithFields(logrus.Fields{"cpu": s.Static.CPU.Name, "arch": s.Static.Host.Architecture}).Debug("system updated")
}
