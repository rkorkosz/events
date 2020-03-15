package event

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPathVersionExtractor(t *testing.T) {
	version := "v1"
	u := fmt.Sprintf("https://example.com/%s", version)
	req := httptest.NewRequest("GET", u, nil)
	out := PathVersionExtractor(req)
	if out != version {
		t.Errorf("Wrong version: got %s want %s", out, version)
	}
}

func TestPathTypeExtractor(t *testing.T) {
	typ := "user"
	u := fmt.Sprintf("https://example.com/v1/%s", typ)
	req := httptest.NewRequest("GET", u, nil)
	out := PathTypeExtractor(req)
	if out != typ {
		t.Errorf("Wrong type: got %s want %s", out, typ)
	}
}

func TestMiddleware(t *testing.T) {
	store := NewInMemoryStore()
	fn := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Test")
	})
	handler := Middleware(store, nil)(fn)
	req := httptest.NewRequest("GET", "http://example.com/v1/user", nil)
	handler.ServeHTTP(httptest.NewRecorder(), req)
	if len(store.db) != 1 {
		t.Errorf("Middleware error: %v", store.db)
	}
}
