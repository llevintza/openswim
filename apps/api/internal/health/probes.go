package health

import (
	"encoding/json"
	"net/http"
)

// RegisterRoutes mounts reserved container probe stubs.
// Real semantics are defined in docs/adr/0002-api-observability.md (E12-F5-T3).
func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /livez", stub("E12-F5-T3"))
	mux.HandleFunc("GET /readyz", stub("E12-F5-T3"))
}

func stub(ticket string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotImplemented)
		_ = json.NewEncoder(w).Encode(map[string]string{
			"status": "not_implemented",
			"ticket": ticket,
			"adr":    "docs/adr/0002-api-observability.md",
		})
	}
}
