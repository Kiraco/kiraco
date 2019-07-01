package utils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadEnv(t *testing.T) {
	LoadEnv()
	assert.Equal(t, ":8080", os.Getenv("PORT"))
}
