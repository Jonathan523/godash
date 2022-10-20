package logging

import (
	"github.com/sirupsen/logrus"
	"launchpad/config"
)

func NewGlobalLogger() {
	var conf Config
	logrus.SetFormatter(&logrus.TextFormatter{TimestampFormat: "2006/01/02 15:04:05", FullTimestamp: true})
	config.ParseViperConfig(&conf, config.AddViperConfig("logging"))
	conf.setConfigLogLevel()
}

func (conf *Config) setConfigLogLevel() {
	logLevel, err := logrus.ParseLevel(conf.LogLevel)
	if err != nil {
		logrus.SetLevel(logrus.FatalLevel)
	} else {
		logrus.SetLevel(logLevel)
	}
}
