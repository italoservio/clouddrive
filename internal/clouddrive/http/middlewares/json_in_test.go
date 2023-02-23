package middlewares_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/italoservio/clouddrive/internal/clouddrive/http/middlewares"
)

func TestShouldThrowErrorWhenHasAnEmptyContentType(t *testing.T) {
	handler := func(wri http.ResponseWriter, req *http.Request) *middlewares.HttpResponse {
		return middlewares.JsonIn()(wri, req)
	}

	req := httptest.NewRequest("GET", "http://example.com", nil)
	req.Header = map[string][]string{}
	wri := httptest.NewRecorder()
	res := handler(wri, req)

	if res == nil {
		t.Fatalf("Should throw error when has no content-type in request header")
	}
}

func TestShouldThrowBadRequestErrorWhenHasAnEmptyContentType(t *testing.T) {
	handler := func(wri http.ResponseWriter, req *http.Request) *middlewares.HttpResponse {
		return middlewares.JsonIn()(wri, req)
	}

	req := httptest.NewRequest("GET", "http://example.com", nil)
	req.Header = map[string][]string{}
	wri := httptest.NewRecorder()
	res := handler(wri, req)

	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("Should throw 404 BAD REQUEST when has no content-type in request header")
	}
}

func TestShouldNotThrowErrorWhenHasAJsonContentType(t *testing.T) {
	handler := func(wri http.ResponseWriter, req *http.Request) *middlewares.HttpResponse {
		return middlewares.JsonIn()(wri, req)
	}

	req := httptest.NewRequest("GET", "http://example.com", nil)
	req.Header = map[string][]string{
		http.CanonicalHeaderKey("content-type"): {"application/json"},
	}
	wri := httptest.NewRecorder()
	res := handler(wri, req)

	if res != nil {
		t.Fatalf("Should not throw error when has a json content-type in request header")
	}
}
