package hardware

import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

var (
	virtualMemory, _ = mem.VirtualMemory()
	cpuInfoStats, _  = cpu.Info()
)

type memInfo struct {
	Total       uint64  `json:"total"`
	Available   uint64  `json:"available"`
	Used        uint64  `json:"used"`
	UsedPercent float64 `json:"usedPercent"`
	Free        uint64  `json:"free"`
}

type cpuInfo struct {
	CPU        int32  `json:"cpu"`
	VendorID   string `json:"vendorId"`
	Family     string `json:"family"`
	Model      string `json:"model"`
	Stepping   int32  `json:"stepping"`
	PhysicalID string `json:"physicalId"`
	CoreID     string `json:"coreId"`
	Cores      int32  `json:"cores"`
}

func MemInfo() *memInfo {
	var memInfo = memInfo{Total: virtualMemory.Total, Free: virtualMemory.Free}
	return &memInfo
}

func CpuInfo() *cpuInfo {
	var first = cpuInfoStats[0]
	var cpuInfo = cpuInfo{CPU: first.CPU, VendorID: first.VendorID, Family: first.Family, Model: first.Model, Stepping: first.Stepping, PhysicalID: first.PhysicalID, Cores: first.Cores, CoreID: first.CoreID}
	return &cpuInfo
}
