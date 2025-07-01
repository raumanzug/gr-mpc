package interfaces

type IApplicationContext interface {
	Backend() IBackend
	Lifecycle() ILifecycle
	Preferences() IPreferences
	ServerTabs() IServerTabs
	UI() IUI
}
