package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultConfig(t *testing.T) {
	testConfig := NewConfig()

	assert.Equal(t, "1337", testConfig.Port, "Default Port not set correctly")
}

func TestConfigReadsCorrectEnvVariables(t *testing.T) {
	os.Setenv("PORT", "42")
	testConfig := NewConfig()

	assert.Equal(t, "42", testConfig.Port, "Port is not read correctly")
}
