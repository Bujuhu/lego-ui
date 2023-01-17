package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/bujuhu/lego-ui/services"
)

//Graphical User Interface configuration and initalisation
type MainWindow struct {
	services services.SingletonServices
	data []string
}

func New(services services.SingletonServices, data []string) MainWindow {
	return MainWindow { services: services, data: data }
}

func (mw MainWindow) Show(a fyne.App) {
	win := a.NewWindow("LEGO-UI")
	win.Resize(fyne.NewSize(1000, 1000))

	//Invoke partial View here to get DNS provider specific form

	dns1Form := CertificateForm(mw.services)
	dns2Form := CertificateForm(mw.services)

	userConfigView := container.NewTabItem("Settings", UserConfigView(mw.services))
	addItem := container.NewTabItem("NEW", CertificateForm(mw.services))
	tabItem1 := container.NewTabItem("DNS1", dns1Form)
	tabItem2 := container.NewTabItem("DNS2", dns2Form)
	tabs := container.NewAppTabs(addItem, userConfigView, tabItem1, tabItem2)

	win.SetContent(tabs)
	win.ShowAndRun()
}
