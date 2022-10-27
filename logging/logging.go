package logging

import (
	"github.com/sirupsen/logrus"
	"godash/config"
)

func NewGlobalLogger() {
	var conf PackageConfig
	logrus.SetFormatter(&logrus.TextFormatter{TimestampFormat: "2006/01/02 15:04:05", FullTimestamp: true})
	config.ParseViperConfig(&conf, config.AddViperConfig("logging"))
	conf.setConfigLogLevel()
}

func (conf *PackageConfig) setConfigLogLevel() {
	logLevel, err := logrus.ParseLevel(conf.LogLevel)
	if err != nil {
		logrus.SetLevel(logrus.FatalLevel)
	} else {
		logrus.SetLevel(logLevel)
	}
}
