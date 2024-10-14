package ui_test

import (
	"testing"

	"github.com/kodflow/vpn/src/backend/ui"
	"github.com/stretchr/testify/assert"
)

func TestAPP(t *testing.T) {
	uiApp := ui.App()
	assert.NotNil(t, uiApp, "app should be not nil")
	uiApp2 := ui.App()
	assert.NotNil(t, uiApp2, "app should be not nil")
	assert.Equal(t, &uiApp, &uiApp2, "app should be the same")
}
