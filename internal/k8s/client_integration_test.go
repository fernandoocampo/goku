package k8s_test

import (
	"flag"
	"testing"

	"github.com/fernandoocampo/goku/internal/k8s"
	"github.com/stretchr/testify/assert"
)

var integration = flag.Bool("integration", false, "integration test")

func TestNewClient(t *testing.T) {
	if !*integration {
		t.Skip("this is an integration test, please provide the integration flag")
	}
	client, err := k8s.NewClient()
	assert.NoError(t, err)
	assert.NotNil(t, client)
}
