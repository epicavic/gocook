package metrics

import (
	"net/http"

	"github.com/rcrowley/go-metrics"
)

// ReportHandler will emit the current metrics in json format
func ReportHandler(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)

	t := metrics.GetOrRegisterTimer("reporthandler.writemetrics", nil)
	t.Time(func() {
		metrics.WriteJSONOnce(metrics.DefaultRegistry, w)
	})
}
