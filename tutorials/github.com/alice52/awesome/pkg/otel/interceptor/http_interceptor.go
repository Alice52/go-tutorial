package interceptor

import (
	"github.com/alice52/awesome/pkg/otel/handler"
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func NewHTTPHandler() http.Handler {
	mux := http.NewServeMux()

	// handleFunc 是 mux.HandleFunc 的替代品, 它使用 http.route 模式丰富了 handler 的 HTTP 测量
	handleFunc := func(pattern string, handlerFunc func(http.ResponseWriter, *http.Request)) {
		// 为 HTTP 测量配置 "http.route"
		mux.Handle(pattern, otelhttp.WithRouteTag(pattern, http.HandlerFunc(handlerFunc)))
	}

	// Register handlers.
	handleFunc("/roll", handler.Roll)

	// 为整个服务器添加 HTTP 测量。
	return otelhttp.NewHandler(mux, "/")
}
