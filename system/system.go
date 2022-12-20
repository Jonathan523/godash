package system

import (
	"go.uber.org/zap"
	"godash/hub"
	"time"
)

func NewSystemService(enabled bool, logging *zap.SugaredLogger, hub *hub.Hub) *System {
	var s Config
	if enabled {
		s = Config{log: logging, hub: hub}
		s.Initialize()
	}
	return &s.System
}

func (c *Config) UpdateLiveInformation() {
	for {
		c.liveCpu()
		c.liveRam()
		c.liveDisk()
		c.uptime()
		c.hub.LiveInformationCh <- hub.Message{WsType: hub.System, Message: c.System.Live}
		time.Sleep(1 * time.Second)
	}
}

func (c *Config) Initialize() {
	c.System.Static.Host = staticHost()
	c.System.Static.CPU = staticCpu()
	c.System.Static.Ram = staticRam()
	c.System.Static.Disk = staticDisk()
	go c.UpdateLiveInformation()
	c.log.Debugw("system updated", "cpu", c.System.Static.CPU.Name, "arch", c.System.Static.Host.Architecture)
}
