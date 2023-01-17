package netcup

import (
	"testing"
	"fyne.io/fyne/v2/app"
	"github.com/bujuhu/lego-ui/services"
	"github.com/bujuhu/lego-ui/services/configService"
	"github.com/bujuhu/lego-ui/services/userLogService"
)

func TestPartialView(t *testing.T) {
	a := app.NewWithID("LEGO-Test")
	w := a.NewWindow("LEGO-Test")
	var services = services.SingletonServices {
		UserConfig: config.New(a.Preferences()),
		UserLog: userLogService.New(),
	}
	w.SetContent(PartialView(services))
}