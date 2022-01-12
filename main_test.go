package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFibbonaciRoute(t *testing.T) {

	ts := httptest.NewServer(setupServer())
	defer ts.Close()

	tests := []struct {
		name        string
		n           string
		want_status int
	}{
		{"send_invalid_type", "this_is_not_a_number", 400},
		{"send_negative_value", "-1", 400},
		{"valid_small_number", "5", 200},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := http.Get(fmt.Sprintf("%s/fibbonaci?n=%s", ts.URL, tt.n))

			if err != nil {
				t.Fatalf("Expected no error, got %v", err)
			}

			if resp.StatusCode != tt.want_status {
				t.Fatalf("Expected status code %d, got %v", tt.want_status, resp.StatusCode)
			}

		})
	}
}
