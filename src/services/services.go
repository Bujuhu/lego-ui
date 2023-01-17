package services

import(
	"fyne.io/fyne/v2"
	"github.com/bujuhu/lego-ui/services/userLogService"
	"github.com/bujuhu/lego-ui/services/configService"
)

//The SingletonServices Struct is used, to simply pass long living, single instance services, everywhere where they are needed
type SingletonServices struct {
	UserLog userLogService.UserLogService
	UserConfig config.Config
	Preferences fyne.Preferences
}