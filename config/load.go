package config

import (
	"bytes"
	"fmt"
	"gopkg.in/yaml.v2"
	"io"
)

// LoadConfiguration
// Looks for a yaml file in the root of the application and loads the content into the Options variable this configuration
// type, this type is then returned to the caller, if the file is not found it will cause the application to close.
func (c *Configuration) LoadConfiguration(file io.Reader) (err error) {
	buf := new(bytes.Buffer)
	if _, err := buf.ReadFrom(file); err != nil {
		return fmt.Errorf("unable to read configuration file: %w", err)
	}

	if err := yaml.Unmarshal(buf.Bytes(), c); err != nil {
		return fmt.Errorf("unable to parse the yaml file: %v", err)
	}

	return nil
}
