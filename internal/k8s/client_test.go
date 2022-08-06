package k8s_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fernandoocampo/goku/internal/k8s"
	"k8s.io/client-go/rest"
)

func TestNewClient(t *testing.T) {
	// Given
	server := httptest.NewServer(&k8sdummy{})
	defer server.Close()
	configData := rest.Config{
		Host: server.URL,
	}
	// When
	client, err := k8s.NewClient(&configData)
	// Then
	if err != nil {
		t.Errorf("unexpected error: %s", err)
		t.FailNow()
	}
	if client == nil {
		t.Error("client cannot be nil")
	}
}

type k8sdummy struct {
}

func (k *k8sdummy) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
}
