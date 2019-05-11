package gridlock

import (
	"github.com/denisbrodbeck/machineid"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"

	"github.com/Stumblinbear/gridlock/api"
)

func GetSystemInfo() (api.SystemInfo, error) {
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

	return api.SystemInfo{
		ID:       id,
		Hostname: hostinfo.Hostname,

		Platform: api.PlatformInfo{
			ID:      hostinfo.OS,
			Name:    hostinfo.Platform,
			Version: hostinfo.PlatformVersion,
		},

		CPU: api.CPUInfo{
			Name: cpuinfo[0].ModelName,
		},
		RAM: api.RAMInfo{
			Total: meminfo.Total,
			Free:  meminfo.Free,
			Used:  meminfo.Used,
		},
	}, nil
}
