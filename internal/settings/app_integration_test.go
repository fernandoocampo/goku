package settings_test

import (
	"flag"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/fernandoocampo/goku/internal/filesystems"
	"github.com/fernandoocampo/goku/internal/settings"
)

var integration = flag.Bool("integration", false, "integration test")

func TestSetUpApplicationE2E(t *testing.T) {
	flag.Parse()
	if !*integration {
		t.Skip("this is an integration test, please provide the integration flag")
	}
	// Given
	expectedSetUpData := settings.Configuration{
		DefaultNamespace: "default",
	}
	setupData := settings.Configuration{
		DefaultNamespace: "default",
	}
	newSetting := settings.New(&filesystems.OSFileSystem{})
	// When
	err := newSetting.SetUp(setupData)
	if err != nil {
		t.Fatalf("unexpected error setting up the application: %s", err)
	}
	setupObtained, err := newSetting.LoadConfiguration()
	if err != nil {
		t.Fatalf("unexpected error reading application setup: %s", err)
	}
	// Then
	assert.Equal(t, expectedSetUpData, setupObtained)
}
