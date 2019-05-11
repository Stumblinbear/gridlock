package api

type API struct {
	// Add a remote host to the system
	AddHost func(string, Host)

	// Remove a remote host from the system
	RemoveHost func(string)

	// A launcher tells the system how to start a remote game.
	AddLauncher func(string, Launcher)

	// A platform
	AddPlatform func(string, Launcher)

	// May change. At the moment, a "fetcher" takes in a game name and
	// outputs a metadata struct.
	AddMetadataFetcher func(string, func(string) Metadata)

	QueueNotification func(Notification)
}
