package cmd

import (
	"account_service/config"
)

func Execute() {
	config.InitConfig()
	config.InitTimeZone()
	config.InitDatabase()

}
