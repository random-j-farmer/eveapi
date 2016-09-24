package mapdata

import "github.com/random-j-farmer/eveapi/types"
import "strings"

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

// RegionNameByID returns the name of the region
func RegionNameByID(regionID uint64) string {
	return regionNameByID[regionID]
}

// GatesBySolarSystemID gives a solar systems gates
func GatesBySolarSystemID(solarSystemID uint64) []types.Gate {
	return gatesBySolarSystemID[solarSystemID]
}

// NameAndID is a name/id pair for searching
type NameAndID struct {
	Name string
	ID   uint64
}

// FindRegions finds regions by prefix
func FindRegions(prefix string) []NameAndID {
	prefix = strings.ToLower(prefix)
	var result []NameAndID
	for id, name := range regionNameByID {
		if strings.HasPrefix(strings.ToLower(name), prefix) {
			result = append(result, NameAndID{Name: name, ID: id})
		}
	}

	return result
}

// FindSolarSystems finds systems by prefix
func FindSolarSystems(prefix string) []NameAndID {
	prefix = strings.ToLower(prefix)
	var result []NameAndID
	for id, system := range systemNameByID {
		if strings.HasPrefix(strings.ToLower(system.SolarSystemName), prefix) {
			result = append(result, NameAndID{Name: system.SolarSystemName, ID: id})
		}
	}

	return result
}
