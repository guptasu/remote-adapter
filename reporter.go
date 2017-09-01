package reporter

import (
	"context"

	"github.com/douglas-reid/mixer-noop-reporter/config"
	"github.com/golang/protobuf/proto"

	_ "istio.io/api/mixer/v1"
	"istio.io/mixer/pkg/adapter"
	"istio.io/mixer/template/logentry"
)

type (
	builder struct{}
	handler struct {
		logger adapter.Logger
	}
)

var _ logentry.HandlerBuilder = builder{}
var _ logentry.Handler = handler{}

func (builder) Build(c adapter.Config, env adapter.Env) (adapter.Handler, error) {
	return handler{env.Logger()}, nil
}

func (builder) ConfigureLogEntryHandler(map[string]*logentry.Type) error {
	return nil
}

func (h handler) HandleLogEntry(context.Context, []*logentry.Instance) error {
	h.logger.Warningf("douglas-reid handler in action!")
	return nil
}

func (handler) Close() error { return nil }

// GetBuilderInfo returns the BuilderInfo associated with this adapter implementation.
func GetBuilderInfo() adapter.BuilderInfo {
	return adapter.BuilderInfo{
		Name:        "github.com/douglas-reid/mixer-noop-reporter",
		Description: "Does nothing (useful for testing)",
		SupportedTemplates: []string{
			logentry.TemplateName,
		},
		CreateHandlerBuilder: func() adapter.HandlerBuilder { return builder{} },
		DefaultConfig:        &config.Params{},
		ValidateConfig:       func(msg proto.Message) *adapter.ConfigErrors { return nil },
	}
}
