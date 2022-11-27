package server

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	hertzlogrus "github.com/hertz-contrib/logger/logrus"
	"github.com/sirupsen/logrus"
	"godash/logging"
)

func setupLogging() {
	logger := hertzlogrus.NewLogger()
	logger.Logger().SetFormatter(logging.LogrusFormatter)
	if logrus.GetLevel() != logrus.TraceLevel {
		logger.SetLevel(hlog.LevelError)
	}
	hlog.SetLogger(logger)
}
