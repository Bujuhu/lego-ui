package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"github.com/bujuhu/lego-ui/dnsViewModels/netcup"
	"github.com/bujuhu/lego-ui/services"
)

func CertificateForm(services services.SingletonServices, tabItem *container.TabItem) *fyne.Container {
	
	overhead := container.New(layout.NewVBoxLayout(), Title("Get Your TLS Certificates"))
	inputForm := netcup.PartialView(services,tabItem)

	return container.New(layout.NewGridLayout(2), container.New(layout.NewVBoxLayout(), overhead, inputForm), LogBox(services))
}