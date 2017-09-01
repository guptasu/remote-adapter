package reporter

import (
	"context"

	"github.com/douglas-reid/mixer-noop-reporter/config"

	_ "istio.io/api/mixer/v1"
	"istio.io/mixer/pkg/adapter"
	"istio.io/mixer/pkg/handler"
	"istio.io/mixer/template/logentry"
)

type (
	builder  struct{}
	instance struct {
		logger adapter.Logger
	}
)

var _ logentry.HandlerBuilder = builder{}
var _ logentry.Handler = instance{}

func (builder) Build(c adapter.Config, env adapter.Env) (adapter.Handler, error) {
	return instance{env.Logger()}, nil
}

func (builder) ConfigureLogEntryHandler(map[string]*logentry.Type) error {
	return nil
}

func (h instance) HandleLogEntry(context.Context, []*logentry.Instance) error {
	h.logger.Warningf("douglas-reid handler in action!")
	return nil
}

func (instance) Close() error { return nil }

// GetInfo returns the BuilderInfo associated with this adapter implementation.
func GetInfo() handler.Info {
	return handler.Info{
		Name:        "mixer-noop-reporter",
		Impl:        "github.com/douglas-reid/mixer-noop-reporter",
		Description: "Does nothing (useful for testing)",
		SupportedTemplates: []string{
			logentry.TemplateName,
		},
		CreateHandlerBuilder: func() adapter.HandlerBuilder { return builder{} },
		DefaultConfig:        &config.Params{},
		ValidateConfig:       func(adapter.Config) *adapter.ConfigErrors { return nil },
	}
}
