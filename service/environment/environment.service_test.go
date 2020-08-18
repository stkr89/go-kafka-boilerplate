package environment

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestEnvironmentService_GetEnvVars_shouldReturnEnvironmentVariables(t *testing.T) {
	// given
	_ = os.Setenv("test", "test")

	// when
	envVars := NewEnvironmentService().GetEnvVars()

	// then
	assert.Equal(t, "test", envVars.Test)
}

func TestEnvironmentService_NewEnvironmentService_shouldReturnNewInstance(t *testing.T) {
	// when
	environmentService := NewEnvironmentService()

	// then
	assert.NotNil(t, environmentService)
}
