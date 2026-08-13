package health_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/llevintza/openswim/apps/api/internal/health"
	"github.com/llevintza/openswim/apps/api/internal/metrics"
)

func TestReservedOpsStubs(t *testing.T) {
	mux := http.NewServeMux()
	health.RegisterRoutes(mux)
	metrics.RegisterRoutes(mux)

	for _, path := range []string{"/livez", "/readyz", "/metrics"} {
		req := httptest.NewRequest(http.MethodGet, path, nil)
		rec := httptest.NewRecorder()
		mux.ServeHTTP(rec, req)
		if rec.Code != http.StatusNotImplemented {
			t.Fatalf("%s: status = %d, want 501", path, rec.Code)
		}
		var body map[string]string
		if err := json.NewDecoder(rec.Body).Decode(&body); err != nil {
			t.Fatalf("%s: decode: %v", path, err)
		}
		if body["status"] != "not_implemented" {
			t.Fatalf("%s: status field = %q", path, body["status"])
		}
	}
}
