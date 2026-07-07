package main

var playerIDMap = make(map[uint64]uint8)

var nextPlayerID uint8 = 0

func GetPlayerID(steamID uint64) uint8 {

	if id, ok := playerIDMap[steamID]; ok {
		return id
	}

	id := nextPlayerID

	playerIDMap[steamID] = id

	nextPlayerID++

	return id
}
