package events_test

import (
	"testing"

	"github.com/kodflow/vpn/src/backend/ui/events"
	"github.com/stretchr/testify/assert"
	"github.com/wailsapp/wails/v3/pkg/application"
)

func TestEmptyTime(t *testing.T) {
	assert.Panics(t, func() {
		events.Time(nil)
	}, "should panic")
}

func TestTime(t *testing.T) {
	events.Time(application.New(application.Options{}))
}
