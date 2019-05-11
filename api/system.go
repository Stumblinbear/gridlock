package api

type SystemInfo struct {
	ID string `json:"id"`

	Hostname string `json:"hostname"`
	Address  string `json:"address"`

	Platform PlatformInfo `json:"platform"`

	CPU CPUInfo `json:"cpu"`
	RAM RAMInfo `json:"ram"`
}

type PlatformInfo struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Version string `json:"version"`
}

type CPUInfo struct {
	Name string `json:"name"`
}

type RAMInfo struct {
	Total uint64 `json:"total"`
	Free  uint64 `json:"free"`
	Used  uint64 `json:"used"`
}
