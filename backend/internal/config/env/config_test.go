package env

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	config, err := LoadConfig(".", "test_secrets")
	assert.Nil(t, err)
	assert.Equal(t, "foo", config.TwitterAuth.AppID)
	assert.Equal(t, "bar", config.TwitterAuth.ApiKey)
	assert.Equal(t, "baz", config.TwitterAuth.ApiSecret)
	assert.Equal(t, "fizzbazz", config.TwitterAuth.ApiToken)
}

func TestCondigError(t *testing.T){
	config, err := LoadConfig("foo", "bar")
	assert.Error(t, err)
	assert.Nil(t, config)
}