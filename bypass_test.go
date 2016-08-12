package bypass

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func testHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Test"))
}

func skipHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Skip"))
}

func newBypassHandler() http.Handler {
	handler := http.HandlerFunc(testHandler)
	b := NewBypass()
	b.AddPath("/skip", skipHandler)
	return b.Handle(handler)
}

func TestDoesBypass(t *testing.T) {
	h := newBypassHandler()
	req, _ := http.NewRequest("GET", "http://localhost/skip", nil)
	recorder := httptest.NewRecorder()
	h.ServeHTTP(recorder, req)

	body := recorder.Body.String()
	if body != "Skip" {
		t.Fatalf("got %s but wanted Skip...", body)
	}
}

func TestRegularRoute(t *testing.T) {
	h := newBypassHandler()
	req, _ := http.NewRequest("GET", "http://localhost/test", nil)
	recorder := httptest.NewRecorder()
	h.ServeHTTP(recorder, req)

	body := recorder.Body.String()
	if body != "Test" {
		t.Fatalf("got %s but wanted Test...", body)
	}
}
