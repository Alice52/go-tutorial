package zaps

import (
	"context"
	"errors"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
)

func main() {
	ctx := context.Background()
	// Wrap zap logger to extend Zap with API that accepts a context.Context.
	log := otelzap.New(zap.NewExample())

	// And then pass ctx to propagate the span.
	log.Ctx(ctx).Error("hello from zap",
		zap.Error(errors.New("hello world")),
		zap.String("foo", "bar"))

	// Alternatively.
	log.ErrorContext(ctx, "hello from zap",
		zap.Error(errors.New("hello world")),
		zap.String("foo", "bar"))
}
