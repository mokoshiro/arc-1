package logger

import (
	"go.uber.org/zap"
)

var L *zap.Logger

func Init() {
	L, _ = zap.NewProduction()
}

func Destruction() {
	L.Sync()
}
