package main

import (
	"encoding/json"
	t "github.com/TyphoonMC/TyphoonCore"
	"github.com/TyphoonMC/go.uuid"
)

type SpawnConfig struct {
	Schematic *string `json:"schematic"`
	Location *t.Location `json:"location"`
	Dimension string `json:"dimension"`
	DimensionParsed t.Dimension
	Gamemode string `json:"gamemode"`
	GamemodeParsed t.Gamemode
}

type Config struct {
	JoinMessage      json.RawMessage `json:"join_message"`
	BossBar          json.RawMessage `json:"boss_bar"`
	PlayerListHeader json.RawMessage `json:"playerlist_header"`
	PlayerListFooter json.RawMessage `json:"playerlist_footer"`
	Spawn *SpawnConfig `json:"spawn"`
	Brand *string `json:"brand"`
}

var (
	config        Config
	bossbarCreate t.PacketBossBar
	playerListHF  t.PacketPlayerListHeaderFooter
	spawn *t.Map
)

func loadConfig(core *t.Core) {
	core.GetConfig(&config)

	if config.BossBar != nil {
		bossbarCreate = t.PacketBossBar{
			UUID:     uuid.Must(uuid.NewV4()),
			Action:   t.BOSSBAR_ADD,
			Title:    string(config.BossBar),
			Health:   1.0,
			Color:    t.BOSSBAR_COLOR_RED,
			Division: t.BOSSBAR_NODIVISION,
			Flags:    0,
		}
	}

	playerListHF = t.PacketPlayerListHeaderFooter{}
	if config.PlayerListHeader != nil {
		msg := string(config.PlayerListHeader)
		playerListHF.Header = &msg
	}
	if config.PlayerListFooter != nil {
		msg := string(config.PlayerListFooter)
		playerListHF.Footer = &msg
	}
	if config.Spawn != nil && config.Spawn.Schematic != nil {
		m, err := t.LoadSchematic(*config.Spawn.Schematic)
		if err != nil {
			panic(err)
		}
		spawn = m
	}
	if config.Spawn.Dimension == "END" {
		config.Spawn.DimensionParsed = t.END
	} else if config.Spawn.Dimension == "NETHER" {
		config.Spawn.DimensionParsed = t.NETHER
	} else if config.Spawn.Dimension == "OVERWORLD" {
		config.Spawn.DimensionParsed = t.OVERWORLD
	} else { panic("Unknown dimension " + config.Spawn.Dimension) }
	
	if config.Spawn.Gamemode == "CREATIVE" {
		config.Spawn.GamemodeParsed = t.CREATIVE
	} else if config.Spawn.Gamemode == "SURVIVAL" {
		config.Spawn.GamemodeParsed = t.SURVIVAL
	} else if config.Spawn.Gamemode == "ADVENTURE" {
		config.Spawn.GamemodeParsed = t.ADVENTURE
	} else if config.Spawn.Gamemode == "SPECTATOR" {
		config.Spawn.GamemodeParsed = t.SPECTATOR
	} else { panic("Unknown gamemode " + config.Spawn.Gamemode) }
}
