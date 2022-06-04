package entities

import (
	"kafka-pub-sub/config"

	"go.uber.org/zap"
)

type ServiceOptions struct {
	Cfg config.Config
	Log *zap.Logger
}
