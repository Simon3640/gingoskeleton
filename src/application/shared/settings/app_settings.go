package settings

type AppSettings struct {
	// Application
	AppName    string `env:"APP_NAME" envDefault:"Gingoskeleton"`
	AppEnv     string `env:"APP_ENV" envDefault:"development"`
	AppPort    string `env:"APP_PORT" envDefault:"8080"`
	AppVersion string `env:"APP_VERSION" envDefault:"0.0.1"`
	EnableLog  bool   `env:"ENABLE_LOG" envDefault:"true"`
	DebugLog   bool   `env:"DEBUG_LOG" envDefault:"true"`
}

func NewAppSettings() *AppSettings {
	return &AppSettings{
		AppName:    "Gingoskeleton",
		AppEnv:     "development",
		AppPort:    "8080",
		AppVersion: "0.0.1",
		EnableLog:  true,
		DebugLog:   true,
	}
}

func (as *AppSettings) Initialize(values map[string]string) {
	as.AppName = values["APP_NAME"]
	as.AppEnv = values["APP_ENV"]
	as.AppPort = values["APP_PORT"]
	as.AppVersion = values["APP_VERSION"]
	as.EnableLog = values["ENABLE_LOG"] == "true"
	as.DebugLog = values["DEBUG_LOG"] == "true"
}

func (as *AppSettings) IsDevelopment() bool {
	return as.AppEnv == "development"
}

var AppSettingsInstance *AppSettings

func init() {
	AppSettingsInstance = NewAppSettings()
}
