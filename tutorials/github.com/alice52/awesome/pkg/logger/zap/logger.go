package zap

import (
	"net/http"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type App struct {
	logger  *zap.Logger
	slogger *zap.SugaredLogger
	clogger *zap.SugaredLogger
}

var (
	app = &App{}
)

func init() {
	app.logger, _ = zap.NewProduction()
	defer app.logger.Sync()

	app.slogger = app.logger.Sugar()
	defer app.slogger.Sync()

	// custom logger: err 单独成文件
	core1 := zapcore.NewCore(getEncoder(), getErrorLogWriter(), zapcore.ErrorLevel)
	core2 := zapcore.NewCore(getEncoder(), getLogWriter(), zapcore.ErrorLevel)
	app.clogger = zap.New(zapcore.NewTee(core1, core2)).Sugar()
	defer app.clogger.Sync()
}

func HttpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		app.slogger.Errorw(
			"Failed",
			"url", url,
			"err", zap.Error(err))

		app.clogger.Errorw(
			"Failed",
			"url", url,
			"err", err)
	} else {
		app.slogger.Infow("Success",
			"url", url,
			"status", resp.Status,
			"body", resp.Body,
		)
		defer resp.Body.Close()
	}
}
