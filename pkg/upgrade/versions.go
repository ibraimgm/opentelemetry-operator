package upgrade

import (
	"github.com/open-telemetry/opentelemetry-operator/pkg/apis/opentelemetry/v1alpha1"
	"github.com/open-telemetry/opentelemetry-operator/pkg/client"
)

type version struct {
	v       string
	upgrade func(cl *client.Clientset, otelcol *v1alpha1.OpenTelemetryCollector) (*v1alpha1.OpenTelemetryCollector, error)
	next    *version
}

var (
	v0_0_1 = version{v: "0.0.1", upgrade: noop, next: &v0_0_2}
	v0_0_2 = version{v: "0.0.2", upgrade: upgrade0_0_2}

	latest = &v0_0_2

	versions = map[string]version{
		v0_0_1.v: v0_0_1,
		v0_0_2.v: v0_0_2,
	}
)

func noop(cl *client.Clientset, otelcol *v1alpha1.OpenTelemetryCollector) (*v1alpha1.OpenTelemetryCollector, error) {
	return otelcol, nil
}
