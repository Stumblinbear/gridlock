package sys

import (
	"github.com/denisbrodbeck/machineid"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

type SystemInfo struct {
	ID string

	Hostname string

	Platform PlatformInfo

	CPU CPUInfo
	RAM RAMInfo
}

type PlatformInfo struct {
	ID      string
	Name    string
	Version string
}

type CPUInfo struct {
	Name string
}

type RAMInfo struct {
	Total uint64
	Free  uint64
	Used  uint64
}

func GetSystemInfo() (*SystemInfo, error) {
	defer func() {
		err := recover()
		if err != nil {
			panic(err)
		}
	}()

	id, err := machineid.ProtectedID("gridlock")
	if err != nil {
		panic("Unable to create unique ID")
	}

	hostinfo, err := host.Info()
	if err != nil {
		panic(err)
	}

	cpuinfo, err := cpu.Info()
	if err != nil {
		panic(err)
	}

	meminfo, err := mem.VirtualMemory()
	if err != nil {
		panic(err)
	}

	return &SystemInfo{
		ID:       id,
		Hostname: hostinfo.Hostname,

		Platform: PlatformInfo{
			ID:      hostinfo.OS,
			Name:    hostinfo.Platform,
			Version: hostinfo.PlatformVersion,
		},

		CPU: CPUInfo{
			Name: cpuinfo[0].ModelName,
		},
		RAM: RAMInfo{
			Total: meminfo.Total,
			Free:  meminfo.Free,
			Used:  meminfo.Used,
		},
	}, nil
}
