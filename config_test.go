package config

import "testing"
import "github.com/stretchr/testify/assert"

var c Config

func init() {
	c = NewConfig("./xbuilder.yml")
}

func TestNewConfig(t *testing.T) {
	assert.Equal(t, nil, c.Error, "error")
	assert.Equal(t, true, c.valid, "valid")
	assert.Equal(t, "isempty", c.Alias.EmptyValue, "EmptyValue")
}
