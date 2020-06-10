package main

type Location int

const (
	// 东
	LocationEast Location = 1
	// 南
	LocationSouth Location = 2
	// 西
	LocationWest Location = 3
	// 北
	LocationNorth Location = 4
)

func (receiver Location) String() string {
	str := ""
	switch receiver {
	case LocationEast:
		str = "东"
	case LocationSouth:
		str = "南"
	case LocationWest:
		str = "西"
	case LocationNorth:
		str = "北"
	default:
		str = "未知"
	}

	return str
}
