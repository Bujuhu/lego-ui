package netcup

import (
	"net/http"
	"time"
	"github.com/go-acme/lego/v4/challenge/dns01"


	"github.com/bujuhu/lego-ui/services"
	"github.com/bujuhu/lego-ui/services/configService"
	"github.com/bujuhu/lego-ui/services/certificateService"
	"fyne.io/fyne/v2/layout"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/go-acme/lego/v4/providers/dns/netcup"
)

// Creates a View which provides form fields to satisfy all requirements for Netcup DNS
func PartialView(services services.SingletonServices, tabItem *container.TabItem) *fyne.Container {

	domainName := canvas.NewText("Domain Name", color.White)
	domainName.Alignment = fyne.TextAlignLeading

	inputDomainName := widget.NewEntry()
	inputDomainName.SetPlaceHolder("Enter text...")
	inputDomainName.OnChanged = func(value string) {
		tabItem.Text = value
	}

	customerNumber := canvas.NewText("Customer Number", color.White)
	customerNumber.Alignment = fyne.TextAlignLeading

	inputCustomerNumber := widget.NewEntry()
	inputCustomerNumber.SetPlaceHolder("Enter text...")
	inputCustomerNumber.Text = services.Preferences.StringWithFallback("netcup.customerNumber", "")
	inputCustomerNumber.OnChanged = func(value string) {
		services.Preferences.SetString("netcup.customerNumber", value)
	}

	apiKey := canvas.NewText("API Key", color.White)
	apiKey.Alignment = fyne.TextAlignLeading

	inputAPIKey := widget.NewEntry()
	inputAPIKey.SetPlaceHolder("Enter text...")
	inputAPIKey.Text = services.Preferences.StringWithFallback("netcup.apiKey", "")
	inputAPIKey.OnChanged = func(value string) {
		services.Preferences.SetString("netcup.apiKey", value)
	}

	apiPass := canvas.NewText("Api Password", color.White)
	apiPass.Alignment = fyne.TextAlignLeading

	inputAPIPass := widget.NewPasswordEntry()
	inputAPIPass.SetPlaceHolder("Enter text...")
	inputAPIPass.Text = services.Preferences.StringWithFallback("netcup.apiPass", "")
	inputAPIPass.OnChanged = func(value string) {
		services.Preferences.SetString("netcup.apiPass", value)
	}

		// Set provider configuration and invoke certificate generation, if Button is pressed
		getCertButton := widget.NewButton("GetCert", func() {
		var providerConfig = ProviderConfigFromUserConfig(services.UserConfig)
		providerConfig.Customer = inputCustomerNumber.Text
		providerConfig.Key = inputAPIKey.Text
		providerConfig.Password = inputAPIPass.Text

			certificateService.Main(services, providerConfig, inputDomainName.Text)
	})

	return container.New(layout.NewVBoxLayout(), domainName, inputDomainName, customerNumber, inputCustomerNumber, apiKey, inputAPIKey, apiPass, inputAPIPass, getCertButton)
}

func ProviderConfigFromUserConfig(userConfig config.Config ) *netcup.Config {
	return &netcup.Config{
		TTL: dns01.DefaultTTL,
		PropagationTimeout: userConfig.PropagationTimeout,
		PollingInterval: userConfig.PollingIntervall,
		HTTPClient: &http.Client{
			Timeout: 10*time.Second,
		},
	}
}