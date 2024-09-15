package config

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

// TestLoadConfiguration
// Ensure that the configuration loader parses the Yaml file correctly
func TestLoadConfiguration(t *testing.T) {

	config := &Configuration{}

	file := strings.NewReader(
		`
DatabaseConfiguration:
  Host: Testing
  Port: 5432
  Username: user
  Password: password
  DriverName: sqlite3
  DatabaseName: ./foo.db

ServerConfiguration:
  Port: 8080

LoggingConfiguration:
  WriterType: console
  Level: DEBUG

Options:
  Testing: testing
`,
	)

	assert.NoError(t, config.LoadConfiguration(file))
	assert.Equal(t, "console", config.LoggingConfiguration.WriterType)
	assert.Equal(t, "DEBUG", config.LoggingConfiguration.Level)
}
