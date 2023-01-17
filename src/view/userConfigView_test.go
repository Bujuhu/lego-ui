package view

import (
	"testing"
	"fyne.io/fyne/v2/app"
	"github.com/bujuhu/lego-ui/services"
	"github.com/bujuhu/lego-ui/services/configService"
	"github.com/bujuhu/lego-ui/services/userLogService"
)


func TestUserConfigView(t *testing.T) {
	a := app.NewWithID("LEGO-Test")
	w := a.NewWindow("LEGO-Test")
	var services = services.SingletonServices {
		UserConfig: config.New(a.Preferences()),
		UserLog: userLogService.New(),
		Preferences: a.Preferences(),
	}
	w.SetContent(UserConfigView(services))
}