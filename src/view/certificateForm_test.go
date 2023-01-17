package view

import (
	"testing"
	"fyne.io/fyne/v2/app"
	"github.com/bujuhu/lego-ui/services"
	"github.com/bujuhu/lego-ui/services/configService"
	"github.com/bujuhu/lego-ui/services/userLogService"
)

func TestCertificateForm(t *testing.T) {
	a := app.NewWithID("LEGO-Test")
	w := a.NewWindow("LEGO-Test")
	var services = services.SingletonServices {
		UserConfig: config.New(a.Preferences()),
		Preferences: a.Preferences(),
		UserLog: userLogService.New(),
	}
	w.SetContent(CertificateForm(services))

}