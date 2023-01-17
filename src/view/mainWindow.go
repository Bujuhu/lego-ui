package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/bujuhu/lego-ui/services"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
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

	userConfigView := container.NewTabItem("Settings", UserConfigView(mw.services))
	tabs := container.NewAppTabs(userConfigView)

	newCertificateButton := widget.NewButton("Create New Certificate", func() {
		createItem := container.NewTabItem("New Certificate", container.New(layout.NewMaxLayout()))
		createItem.Content = CertificateForm(mw.services, createItem)
		tabs.Append(createItem)
		tabs.Select(createItem)
	})

	win.SetContent(container.New(layout.NewVBoxLayout(), tabs, newCertificateButton))
	win.ShowAndRun()
}
