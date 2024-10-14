package ui_test

import (
	"testing"

	"github.com/kodflow/vpn/src/backend/ui"
	"github.com/stretchr/testify/assert"
)

func TestVPN(t *testing.T) {
	uiApp := ui.VPN()
	assert.NotNil(t, uiApp, "app should be not nil")
	uiApp2 := ui.VPN()
	assert.NotNil(t, uiApp2, "app should be not nil")
	assert.Equal(t, &uiApp, &uiApp2, "app should be the same")
}

func TestWindow(t *testing.T) {
	testApp := ui.VPN()
	assert.NotNil(t, testApp, "app should be not nil")
	webviewWindow := ui.Window(testApp)
	assert.NotNil(t, webviewWindow, "webviewWindow should be not nil")
}

func TestWindowAppNil(t *testing.T) {
	assert.Panics(t, func() {
		ui.Window(nil)
	}, "should panic")
}
