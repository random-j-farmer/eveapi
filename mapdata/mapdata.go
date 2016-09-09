package mapdata

import "github.com/random-j-farmer/eveapi/types"

// SolarSystemByID gets the system info
func SolarSystemByID(id uint64) types.SolarSystem {
	return systemNameByID[id]
}

// RegionBySolarSystem gives region information by solar system
func RegionBySolarSystem(solarSystemID uint64) (regionID uint64, regionName string) {
	sysdata := SolarSystemByID(solarSystemID)
	regionID = sysdata.RegionID
	regionName = regionNameByID[regionID]
	// named return
	return
}

// GatesBySolarSystemID gives a solar systems gates
func GatesBySolarSystemID(solarSystemID uint64) []types.Gate {
	return gatesBySolarSystemID[solarSystemID]
}
