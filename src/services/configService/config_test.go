package config

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGeneratingPrivateKeysShouldNotBeEqual(t *testing.T) {
	assert := assert.New(t)
	assert.NotEqual(NewPrivateKey(), NewPrivateKey(), "Generating Private Keys should result in different results every time")
}