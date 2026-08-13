package metrics

import (
	"encoding/json"
	"net/http"
)

// RegisterRoutes mounts the reserved Prometheus metrics path stub.
// Real exposition is defined in docs/adr/0002-api-observability.md (E12-F5-T2).
// Do not add prometheus/client_golang until that task lands.
func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /metrics", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotImplemented)
		_ = json.NewEncoder(w).Encode(map[string]string{
			"status": "not_implemented",
			"ticket": "E12-F5-T2",
			"adr":    "docs/adr/0002-api-observability.md",
		})
	})
}
