package main

import (
	"context"
	"errors"
	"github.com/alice52/awesome/pkg/otel/interceptor"
	"github.com/alice52/awesome/pkg/otel/otels"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() (err error) {

	// 平滑处理 SIGINT (CTRL+C) .
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// 设置 OpenTelemetry.
	if otelShutdown, err := otels.SetupOTel(ctx); err != nil {
		return err
	} else {
		// 妥善处理停机，确保无泄漏
		defer func() {
			err = errors.Join(err, otelShutdown(context.Background()))
		}()
	}

	// 启动 HTTP server.
	srv := &http.Server{
		Addr:         ":8080",
		BaseContext:  func(_ net.Listener) context.Context { return ctx },
		ReadTimeout:  time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      interceptor.NewHTTPHandler(),
	}
	srvErr := make(chan error, 1)
	go func() {
		srvErr <- srv.ListenAndServe()
	}()

	// 等待中断.
	select {
	case err = <-srvErr:
		// 启动 HTTP 服务器时出错.
		return
	case <-ctx.Done():
		// 等待第一个 CTRL+C, 尽快停止接收信号通知.
		stop()
	}

	// 调用 Shutdown 时，ListenAndServe 会立即返回 ErrServerClosed
	err = srv.Shutdown(context.Background())
	return
}
