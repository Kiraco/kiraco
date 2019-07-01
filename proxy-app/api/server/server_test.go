package server

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/kataras/iris"
)

func TestSetup(t *testing.T) {
	app := SetUp()
	assert.Equal(t, reflect.TypeOf(iris.New()), reflect.TypeOf(app))
}
