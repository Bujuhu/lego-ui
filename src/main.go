package main

import (
	"fyne.io/fyne/v2/app"
	"github.com/bujuhu/lego-ui/view"
	"github.com/bujuhu/lego-ui/services/configService"
	"github.com/bujuhu/lego-ui/services/userLogService"
	"github.com/bujuhu/lego-ui/services"
	"net/http"
	"time"
)

var client *http.Client

var data = []string{"Cert1", "Cert2", "Cert3"}


//Entrypoint of the Application
func main() {
	client = &http.Client{Timeout: 10 * time.Second}
	a := app.NewWithID("at.bujuhu.legoui")
	var services = services.SingletonServices {
		UserConfig: config.New(a.Preferences()),
		UserLog: userLogService.New(),
		Preferences: a.Preferences(),
	}
	view.New(services, data).Show(a)
}
