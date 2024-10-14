package service_test

import (
	"testing"

	"github.com/kodflow/vpn/src/backend/ui/service"
	"github.com/stretchr/testify/assert"
)

func TestHello_World(t *testing.T) {
	greet := service.Hello{}
	assert.Equal(t, "Hello world!", greet.World("world"), "should be equal")
}
