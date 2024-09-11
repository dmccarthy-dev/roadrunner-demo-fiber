package roadrunner_demo_fiber

import (
	"context"
	"github.com/dmccarthy-dev/roadrunner-demo-fiber/internal/app/fiberapp"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/roadrunner-server/endure/v2/dep"
	"github.com/roadrunner-server/errors"
	"go.uber.org/zap"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	"net/http"
	"strconv"
)

const (
	PluginName     = "roadrunner_demo_fiber"
	RootPluginName = "http"
)

type Plugin struct {
	logger *zap.Logger
}

type Configurer interface {
	// UnmarshalKey takes a single key and unmarshal it into a Struct.
	UnmarshalKey(name string, out any) error
	// Has checks if config section exists.
	Has(name string) bool
}

type Logger interface { // <-- logger plugin implements
	NamedLogger(name string) *zap.Logger
}

func (p *Plugin) Init(cfg Configurer, log Logger) error {
	p.logger = log.NamedLogger("roadrunner_demo_fiber")
	p.logger.Info("start initializing demo plugin")

	const op = errors.Op("roadrunner_demo_fiber_init")

	if !cfg.Has(RootPluginName) {
		return errors.E(op, errors.Disabled)
	}

	p.logger.Info("end initializing demo plugin")
	return nil
}

func (p *Plugin) Middleware(next http.Handler) http.Handler {

	fiberApp := fiberapp.Boot()

	fiberApp.Use(func(c *fiber.Ctx) error {

		h := adaptor.HTTPHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			span, ok := tracer.SpanFromContext(c.UserContext())

			if ok {
				r.Header.Set("x-datadog-parent-id", strconv.FormatUint(span.Context().SpanID(), 10))
				r.Header.Set("x-datadog-trace-id", strconv.FormatUint(span.Context().TraceID(), 10))
			}

			next.ServeHTTP(w, r)
		}))
		return h(c)
	})

	return adaptor.FiberApp(fiberApp)
}

func (p *Plugin) Stop(_ context.Context) error {
	p.logger.Info("Stopping plugin")
	return nil
}

func (p *Plugin) Name() string {
	return PluginName
}

func (p *Plugin) Collects() []*dep.In {
	return nil
}
