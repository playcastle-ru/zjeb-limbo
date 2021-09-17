package main

import (
	"fmt"
	t "github.com/TyphoonMC/TyphoonCore"
)

func main() {
	core := t.Init()

	loadConfig(core)
	
	core.SetBrand(*config.Brand)
	core.SetGamemode(config.Spawn.GamemodeParsed)

	if spawn != nil {
		fmt.Println("Using schematic world")
		spawn.Dimension = config.Spawn.DimensionParsed
		if config.Spawn.Location != nil {
			spawn.Spawn = *config.Spawn.Location
		}
		core.SetMap(spawn)
	}

	core.On(func(e *t.PlayerJoinEvent) {
		if config.JoinMessage != nil {
			e.Player.SendRawMessage(string(config.JoinMessage))
		}
		if &bossbarCreate != nil {
			e.Player.WritePacket(&bossbarCreate)
		}
		if &playerListHF != nil {
			e.Player.WritePacket(&playerListHF)
		}
	})
	
	core.On(func(e *t.PlayerChatEvent) {
		e.Player.WritePacket(&t.PacketPlayerPositionLook{
			config.Spawn.Location.X,
			config.Spawn.Location.Y,
			config.Spawn.Location.Z,
			0,
			0,
			0xFF,
			0,
		})
	})
	
	core.Start()
}

type ChunkSave struct {
	X int `json:"x"`
	Y int `json:"y"`
	Bitmask int `json:"bitmask"`
	Data string `json:"data"`
}