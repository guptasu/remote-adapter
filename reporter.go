package reporter

import (
	"context"

	"github.com/douglas-reid/mixer-noop-reporter/config"

	// _ "istio.io/api/mixer/v1"
	"istio.io/mixer/pkg/adapter"
	"istio.io/mixer/template/logentry"
)

type (
	builder  struct{}
	instance struct {
		logger adapter.Logger
	}
)

func (h instance) HandleLogEntry(context.Context, []*logentry.Instance) error {
	h.logger.Warningf("douglas-reid handler in action!")
	return nil
}

func (instance) Close() error { return nil }

// GetInfo returns the BuilderInfo associated with this adapter implementation.
func GetInfo() adapter.Info {
	return adapter.Info{
		Name:        "mixer-noop-reporter",
		Impl:        "github.com/douglas-reid/mixer-noop-reporter",
		Description: "Does nothing (useful for testing)",
		SupportedTemplates: []string{
			logentry.TemplateName,
		},
		DefaultConfig: &config.Params{},
		NewBuilder:    func() adapter.HandlerBuilder { return &builder{} },
	}
}

func (b *builder) SetLogEntryTypes(types map[string]*logentry.Type) {}
func (b *builder) SetAdapterConfig(cfg adapter.Config)              {}
func (*builder) Validate() (ce *adapter.ConfigErrors)               { return }

func (b *builder) Build(context context.Context, env adapter.Env) (adapter.Handler, error) {
	return instance{env.Logger()}, nil
}
