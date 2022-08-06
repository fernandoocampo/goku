package k8s_test

import (
	"flag"
	"testing"

	"github.com/fernandoocampo/goku/internal/k8s"
	"github.com/stretchr/testify/assert"
)

var integration = flag.Bool("integration", false, "integration test")

func TestNewK8sClient(t *testing.T) {
	flag.Parse()
	if !*integration {
		t.Skip("this is an integration test, please provide the integration flag")
	}
	// Given
	configdata, err := k8s.LoadDefaultKubeConfig()
	if err != nil {
		t.Fatalf("unexpected error loading default kube config: %s", err)
	}
	// When
	client, err := k8s.NewClient(configdata)
	// Then
	assert.NoError(t, err)
	assert.NotNil(t, client)
}
