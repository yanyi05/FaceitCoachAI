package main

type Player struct {
	Name      string `json:"name"`
	SteamID64 uint64 `json:"steamId"`
	Team      string `json:"team"`
	TeamGroup string `json:"teamGroup"`
	PlayerID  uint8  `json:"playerId"`
	IsBot     bool   `json:"isBot"`
}

type PlayerStats struct {
	Name      string `json:"name"`
	SteamID64 uint64 `json:"steamId"`

	Kills     int `json:"kills"`
	Deaths    int `json:"deaths"`
	Assists   int `json:"assists"`
	Headshots int `json:"headshots"`

	ADR             float64 `json:"adr"`
	KAST            float64 `json:"kast"`
	Rating          float64 `json:"rating"`
	Accuracy        float64 `json:"accuracy"`
	HeadshotPercent float64 `json:"headshotPercent"`

	Damage      int `json:"damage"`
	DamageTaken int `json:"damageTaken"`

	ShotsFired int `json:"shotsFired"`
	ShotsHit   int `json:"shotsHit"`

	EntryKills    int     `json:"entryKills"`
	EntryDeaths   int     `json:"entryDeaths"`
	EntryAttempts int     `json:"entryAttempts"`
	EntrySuccess  float64 `json:"entrySuccess"`

	AverageTTD float64 `json:"averageTTD"`

	TotalTTDTicks int `json:"totalTTDTicks"`

	TTDCount   int `json:"ttdCount"`
	TradeKills int `json:"tradeKills"`

	TradeDeaths int `json:"tradeDeaths"`

	TradeSuccess float64 `json:"tradeSuccess"`

	OpeningKills  int `json:"openingKills"`
	OpeningDeaths int `json:"openingDeaths"`
}

type PlayerState struct {
	// 时间
	Tick  int
	Round int

	// 玩家
	PlayerID  uint8
	SteamID64 uint64
	Team      string

	// 位置
	X int16
	Y int16
	Z int16

	// 视角
	ViewYaw   float32
	ViewPitch float32

	// 状态
	Alive bool

	HP    uint8
	Armor uint8

	Money int

	// 武器
	Weapon string

	AmmoInMagazine int
	AmmoReserve    int

	ZoomLevel int

	// 动作
	IsScoped      bool
	IsWalking     bool
	IsStanding    bool
	IsDucking     bool
	IsAirborne    bool
	IsReloading   bool
	IsBlinded     bool
	FlashDuration float32

	// 装备
	HasHelmet    bool
	HasDefuseKit bool

	// 特殊行为
	IsPlanting bool
	IsDefusing bool

	// 地图
	LastPlace string

	// 下一阶段计算
	Velocity  float32
	VelocityX float32
	VelocityY float32
	VelocityZ float32
}
