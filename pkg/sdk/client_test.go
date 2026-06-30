package sdk

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {
	t.Run("a new client is created", func(t *testing.T) {
		client := Connect("https://api.vulncheck.com", "vulncheck_token")
		if client == nil {
			t.Error("client is nil")
		}

		assert.NotNil(t, client)
	})
}

func TestSetAuthHeader(t *testing.T) {
	t.Run("auth header is set", func(t *testing.T) {
		client := Connect("https://api.vulncheck.com", "vulncheck_token")
		req := httptest.NewRequest("GET", "/index", nil)
		client.SetAuthHeader(req)

		assert.Equal(t, "Bearer vulncheck_token", req.Header.Get("Authorization"))
	})
}

func TestGetUrl(t *testing.T) {
	t.Run("url is set", func(t *testing.T) {
		client := Connect("https://api.vulncheck.com", "vulncheck_token")
		assert.Equal(t, "https://api.vulncheck.com", client.GetUrl())
	})
}

// TestRequestRespectsContextCancellation locks in Phase 2's contract:
// a Client with a cancelled context never lets the request reach the wire.
func TestRequestRespectsContextCancellation(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Stall long enough that ctx cancellation wins the race.
		time.Sleep(2 * time.Second)
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	ctx, cancel := context.WithCancel(context.Background())
	cancel() // cancel up-front

	client := Connect(srv.URL, "tok").WithContext(ctx)
	_, err := client.Request(http.MethodGet, "/")
	if err == nil {
		t.Fatal("expected an error from a cancelled context, got nil")
	}
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("expected context.Canceled, got %v", err)
	}
}

func TestRequestDefaultsToBackgroundContext(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	// No WithContext() — should default to context.Background and succeed.
	client := Connect(srv.URL, "tok")
	resp, err := client.Request(http.MethodGet, "/")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
}
