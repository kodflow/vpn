package view_test

import (
	"testing"

	"github.com/kodflow/vpn/src/backend/ui/view"
	"github.com/stretchr/testify/assert"
	"github.com/wailsapp/wails/v3/pkg/application"
)

func TestEmptyEntryPoint(t *testing.T) {
	assert.Panics(t, func() {
		view.Entrypoint(nil)
	}, "should panic")
}

func TestEmptyApp(t *testing.T) {
	assert.NotNil(t, view.Entrypoint(application.New(application.Options{})))
}
