package api

type API struct {
	// Add a remote host to the system
	AddHost func(string, Host)

	// Remove a remote host from the system
	RemoveHost func(string)

	// Add a platform to the system
	AddPlatform func(Platform)

	// Add a game to the system
	AddGame func(GameInstance)

	// A launcher tells the system how to start a game
	AddLauncher func(Launcher)

	// May change. At the moment, a resolver takes in a game name and
	// outputs a metadata struct.
	AddMetadataResolver func(MetadataResolver)

	QueueNotification func(Notification)
}
