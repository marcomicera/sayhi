package sayhi

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Tests HTTP routing assuming that all the tested endpoints do no return a 404 on GET
func TestNewRouter(t *testing.T) {

	tests := []struct {
		name     string
		endpoint string
	}{
		{
			name:     "Salutation routing test",
			endpoint: "helloworld",
		},
		{
			name:     "Version routing test",
			endpoint: "versionz",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// Retrieve right endpoint
			router := NewRouter()
			srv := httptest.NewServer(router)
			defer srv.Close()
			res, err := http.Get(fmt.Sprintf("%s/%s", srv.URL, tt.endpoint))
			if err != nil {
				t.Fatalf("Could not send GET request: %v", err)
			}
			defer res.Body.Close()

			// Response code should not be 404
			if res.StatusCode == http.StatusNotFound {
				t.Errorf("Endpoint %s not found (404)", tt.endpoint)
			}
		})
	}
}
