package main

import (
	"net/http/httptest"
	"testing"
)

func TestHello(t *testing.T) {
	rec := httptest.NewRecorder()
	hello(rec, httptest.NewRequest("GET", "/hello", nil))

	if got, want := rec.Body.String(), "Hello, World!\n"; got != want {
		t.Errorf("body = %q, want %q", got, want)
	}
}
