package goodreads

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadConfig(t *testing.T) {
	cfg := new(Config)
	testConfigFile := filepath.Join(".", "testdata", "config.yml")

	cfg, err := ReadConfig(testConfigFile)
	assert.NoError(t, err)

	assert.Equal(t, "testingKey", cfg.DeveloperKey)
	assert.Equal(t, "testingSecret", cfg.DeveloperSecret)
	assert.Equal(t, "testingAccessToken", cfg.AccessToken)
	assert.Equal(t, "testingAccessSecret", cfg.AccessSecret)
}
