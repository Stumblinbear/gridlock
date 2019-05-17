package api

// A host is the machine that physically contains a game.
// This can be the local machine or a remote one.
type Host struct {
	// If the host is a remote machine
	Remote bool `json:"remote"`

	System SystemInfo `json:"system"`

	// The key of this map is the platform ID that launches this game.
	//
	// For example: steam, uplay, ps4, nintendo64, etc
	Platforms map[string]Platform `json:"platforms"`

	// Map of launchers supported by this host
	Launchers map[string]Launcher `json:"launchers"`
}

// A platform is the type of system that the game is run on.
// A library is defined as a list of games that run on a single platform.
// Therefore, of you have a steam library and a uplay library, then you
// will have two platforms.
//
// For example: PC, Steam, Xbox, PlayStation 4, Nintendo 64, etc
type Platform struct {
	// URL to an icon for the platform
	Icon        string `json:"icon"`
	Name        string `json:"name"`
	Description string `json:"description"`

	// List of games available on this platform
	Games map[string]Game `json:"games"`

	// Takes in a game instance and uses it to start the game immediately.
	// This one doesn't understand the concept of
	SpawnGame func(Game) error `json:"-"`
}

type Game struct {
	// Used to categorize the game into the system, but is discarded afterwards
	Platform string `json:"-"`
	Name     string `json:"-"`

	// If the game is installed
	Installed bool `json:"installed"`

	// The install directory of the game (some platforms may not need this)
	Directory string `json:"directory"`

	// The last time the user played the game
	LastPlayed int64 `json:"lastPlayed"`

	// The total time the user played the game
	TimePlayed int `json:"timePlayed"`

	// The amount of times the user played the game
	PlayCount int `json:"playCount"`
}

// A launcher is responsible for the entire process of setting up a game.
// For example, if using a remote streaming service, then the launcher should
// forward the request to the remote machine's platform, then start up the
// streaming system.
type Launcher struct {
	// URL to an icon for the launcher
	Icon        string `json:"icon"`
	Name        string `json:"name"`
	Description string `json:"description"`

	// If the launcher is required on both systems
	Require bool `json:"require"`

	// If the game can be started using this launcher
	CanStart func(GameQuery) bool `json:"-"`

	// Takes in a game instance and uses it to start the game.
	StartGame func(GameQuery) error `json:"-"`
}

// All the information needed to start a game.
type GameQuery struct {
	// The host that the game is starting on
	HostId string `json:"hostId"`

	// The launcher that's starting the game
	LauncherId string `json:"launcherId"`

	// The platform that launches the game
	PlatformId string `json:"platformId"`

	// The ID of the game that's being started
	GameID string `json:"gameId"`
}

// Contains the metadata information for games registered to the system.
type Metadata struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Version     string `json:"version"`

	Rating byte   `json:"rating"`
	Region string `json:"region"`

	Genre     string   `json:"genre"`
	Developer string   `json:"developer"`
	Year      int8     `json:"year"`
	Tags      []string `json:"tags"`
	Links     []string `json:"links"`

	News []string `json:"news"`

	Score MetadataScore `json:"score"`
	Art   MetadataArt   `json:"art"`
}

// How each category rated the game
type MetadataScore struct {
	User      int8 `json:"user"`
	Critic    int8 `json:"critic"`
	Community int8 `json:"community"`
}

// Art assets should be in the form of URLs.
type MetadataArt struct {
	Icon       string `json:"icon"`
	Cover      string `json:"cover"`
	Background string `json:"background"`
}

type MetadataResolver struct {
	Name string

	Resolve func(string) Metadata
}
