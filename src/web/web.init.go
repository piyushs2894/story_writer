package web

import (
	"net/http"
	"time"

	"story_writer/src/common/config"
	"story_writer/src/constant"
	"story_writer/src/manager"
)

type Web struct {
	cfg           *config.Config
	client        *http.Client
	managerModule manager.Manager
}

func New(mod manager.Manager) *Web {
	// Default configuration options
	cfg := config.GetConfig()

	return InitWeb(cfg, mod)
}

func InitWeb(cfg *config.Config, mod manager.Manager) *Web {
	w := Web{
		client:        NewHttpClient(cfg),
		cfg:           cfg,
		managerModule: mod,
	}

	return &w
}

func NewHttpClient(cfg *config.Config) *http.Client {
	client := &http.Client{
		Timeout: time.Duration(constant.HTTPClientTimeout * time.Second),
	}
	return client
}
