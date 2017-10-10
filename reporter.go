package reporter

import (
	"context"

	"github.com/guptasu/remoteAdapter/config"
	remoteRpt "github.com/guptasu/remoteRpt"

	"istio.io/mixer/pkg/adapter"
	"istio.io/mixer/template/logentry"
)

type (
	builder  struct{}
	instance struct {
		logger adapter.Logger
	}
)
var (
	_ logentry.HandlerBuilder = &builder{}
	_ remoteRpt.HandlerBuilder = &builder{}

	_ logentry.Handler = &instance{}
	_ remoteRpt.Handler = &instance{}
)

func (h *instance) HandleLogEntry(context.Context, []*logentry.Instance) error {
	h.logger.Warningf("HandleLogEntry handler in action!")
	return nil
}
func (h *instance) HandleMyCustomReport(context.Context, []*remoteRpt.Instance) error {
	h.logger.Warningf("HandleMyCustomReport handler in action!")
	return nil
}

func (h *instance) Close() error { return nil }

// GetInfo returns the BuilderInfo associated with this adapter implementation.
func GetInfo() adapter.Info {
	return adapter.Info{
		Name:        "mixer-noop-reporter",
		Impl:        "github.com/guptasu/remote-adapter",
		Description: "Does nothing (useful for testing)",
		SupportedTemplates: []string{
			logentry.TemplateName,
			remoteRpt.TemplateName,
		},
		DefaultConfig: &config.Params{},
		NewBuilder:    func() adapter.HandlerBuilder { return &builder{} },
	}
}

func (b *builder) SetLogEntryTypes(types map[string]*logentry.Type) {}
func (b *builder) SetMyCustomReportTypes(types map[string]*remoteRpt.Type) {}

func (b *builder) SetAdapterConfig(cfg adapter.Config)              {}
func (*builder) Validate() (ce *adapter.ConfigErrors)               { return }

func (b *builder) Build(context context.Context, env adapter.Env) (adapter.Handler, error) {
	return &instance{env.Logger()}, nil
}
