package components_test

import (
	"testing"

	"github.com/kodflow/vpn/src/backend/components"
	"github.com/stretchr/testify/assert"
)

func TestHello_World(t *testing.T) {
	greet := components.Hello{}
	assert.Equal(t, "Hello world!", greet.World("world"), "should be equal")
}
