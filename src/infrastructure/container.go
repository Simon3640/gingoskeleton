package infrastructure

import (
	settings "gingoskeleton/src/application/shared/settings"
	config "gingoskeleton/src/infrastructure/config"
)

func init() {
	settings.AppSettingsInstance.Initialize(config.ConfigInstance.ToMap())
}
