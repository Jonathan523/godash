package server

import (
	hertz "github.com/cloudwego/hertz/pkg/app/server"
	"godash/hub"
)

type Server struct {
	Router  *hertz.Hertz
	Hub     *hub.Hub
	Port    int
	PageUrl string `mapstructure:"PAGE_URL"`
	Title   string
}
