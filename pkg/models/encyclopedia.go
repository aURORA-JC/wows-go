package models

type EncyclopediaInfoResponse struct {
	WowsSuccessResponse
	Data struct {
		ShipsUpdatedAt    int64                                    `json:"ships_updated_at"`
		ShipTypes         map[string]string                        `json:"ship_types"`
		Languages         map[string]string                        `json:"languages"`
		ShipModifications map[string]string                        `json:"ship_modifications"`
		ShipModules       map[string]string                        `json:"ship_modules"`
		ShipTypeImages    map[string]EncyclopediaInfoShipTypeImage `json:"ship_type_images"`
		ShipNations       map[string]string                        `json:"ship_nations"`
		GameVersion       string                                   `json:"game_version"`
	} `json:"data"`
}

type EncyclopediaShipsResponse struct {
	WowsSuccessResponse
	Data map[string]EncyclopediaShipsData `json:"data"`
}

type EncyclopediaAchievementsResponse struct {
	WowsSuccessResponse
	Data struct {
		Battle map[string]EncyclopediaAchievementsData `json:"battle"`
	} `json:"data"`
}

type EncyclopediaShipProfileResponse struct {
	WowsSuccessResponse
	Data map[string]EncyclopediaShipsDefaultProfile `json:"data"`
}

type EncyclopediaModulesResponse struct {
	WowsSuccessResponse
	Data map[string]EncyclopediaModulesData `json:"data"`
}

type EncyclopediaAccountLevelsResponse struct {
	WowsSuccessResponse
	Data map[string]EncyclopediaAccountLevelsData `json:"data"`
}

type EncyclopediaCrewsResponse struct {
	WowsSuccessResponse
	Data map[string]EncyclopediaCrewsData `json:"data"`
}

type EncyclopediaCrewSkillsResponse struct {
	WowsSuccessResponse
	Data map[string]EncyclopediaCrewSkillsData `json:"data"`
}

type EncyclopediaCrewRanksResponse struct {
	WowsSuccessResponse
	Data map[string][]EncyclopediaCrewRanksData `json:"data"`
}

type EncyclopediaBattleTypesResponse struct {
	WowsSuccessResponse
	Data map[string]EncyclopediaBattleTypesData `json:"data"`
}

type EncyclopediaConsumablesResponse struct {
	WowsSuccessResponse
	Data map[string]EncyclopediaConsumablesData `json:"data"`
}

type EncyclopediaCollectionsResponse struct {
	WowsSuccessResponse
	Data map[string]EncyclopediaCollectionsData `json:"data"`
}

type EncyclopediaCollectionCardsResponse struct {
	WowsSuccessResponse
	Data map[string]EncyclopediaCollectionCardsData `json:"data"`
}

type EncyclopediaBattleArenasResponse struct {
	WowsSuccessResponse
	Data map[string]EncyclopediaBattleArenasData `json:"data"`
}

type EncyclopediaInfoShipTypeImage struct {
	ImagePremium string `json:"image_premium"`
	Image        string `json:"image"`
	ImageElite   string `json:"image_elite"`
}

type EncyclopediaShipsData struct {
	Description    string                             `json:"description"`
	PriceGold      int                                `json:"price_gold"`
	ShipIdStr      string                             `json:"ship_id_str"`
	HasDemoProfile bool                               `json:"has_demo_profile"`
	Images         *EncyclopediaShipsImages           `json:"images"`
	Modules        *EncyclopediaShipsModulesList      `json:"modules"`
	ModulesTree    map[string]EncyclopediaShipsModule `json:"modules_tree"`
	Nation         string                             `json:"nation"`
	IsPremium      bool                               `json:"is_premium"`
	ShipId         int64                              `json:"ship_id"`
	PriceCredit    int                                `json:"price_credit"`
	DefaultProfile *EncyclopediaShipsDefaultProfile   `json:"default_profile"`
	Upgrades       []int64                            `json:"upgrades"`
	Tier           int                                `json:"tier"`
	NextShips      map[string]int64                   `json:"next_ships"`
	ModSlots       int                                `json:"mod_slots"`
	Type           string                             `json:"type"`
	IsSpecial      bool                               `json:"is_special"`
	Name           string                             `json:"name"`
}
type EncyclopediaShipsImages struct {
	Small   string `json:"small"`
	Large   string `json:"large"`
	Medium  string `json:"medium"`
	Contour string `json:"contour"`
}
type EncyclopediaShipsModulesList struct {
	Engine        []int64 `json:"engine"`
	TorpedoBomber []int64 `json:"torpedo_bomber"`
	Fighter       []int64 `json:"fighter"`
	Hull          []int64 `json:"hull"`
	Artillery     []int64 `json:"artillery"`
	Torpedoes     []int64 `json:"torpedoes"`
	FireControl   []int64 `json:"fire_control"`
	Sonar         []int64 `json:"sonar"`
	FlightControl []int64 `json:"flight_control"`
	DiveBomber    []int64 `json:"dive_bomber"`
}
type EncyclopediaShipsModule struct {
	Name        string  `json:"name"`
	NextModules []int64 `json:"next_modules"`
	IsDefault   bool    `json:"is_default"`
	PriceXp     int     `json:"price_xp"`
	PriceCredit int     `json:"price_credit"`
	NextShips   []int64 `json:"next_ships"`
	ModuleId    int64   `json:"module_id"`
	Type        string  `json:"type"`
	ModuleIdStr string  `json:"module_id_str"`
}
type EncyclopediaShipsDefaultProfile struct {
	Engine              *EncyclopediaShipsEngine            `json:"engine"`
	TorpedoBomber       *EncyclopediaShipsTorpedoBomber     `json:"torpedo_bomber"`
	AntiAircraft        *EncyclopediaShipsAA                `json:"anti_aircraft"`
	DepthCharge         *EncyclopediaShipsDepthCharge       `json:"depth_charge"`
	Mobility            *EncyclopediaShipsMobility          `json:"mobility"`
	Hull                *EncyclopediaShipsHull              `json:"hull"`
	Atbas               *EncyclopediaShipsAtbas             `json:"atbas"`
	Artillery           *EncyclopediaShipsArtillery         `json:"artillery"`
	Torpedoes           *EncyclopediaShipsTorpedoes         `json:"torpedoes"`
	Fighters            *EncyclopediaShipsFighters          `json:"fighters"`
	SubmarineMobility   *EncyclopediaShipsSubmarineMobility `json:"submarine_mobility"`
	FireControl         *EncyclopediaShipsFireControl       `json:"fire_control"`
	SubmarineSonar      *EncyclopediaShipsSubmarineSonar    `json:"submarine_sonar"`
	BattleLevelRangeMax int                                 `json:"battle_level_range_max"`
	BattleLevelRangeMin int                                 `json:"battle_level_range_min"`
	FlightControl       *EncyclopediaShipsFlightControl     `json:"flight_control"`
	SubmarineBattery    *EncyclopediaShipsSubmarineBattery  `json:"submarine_battery"`
	Concealment         *EncyclopediaShipsConcealment       `json:"concealment"`
	Armour              *EncyclopediaShipsArmour            `json:"armour"`
	Weaponry            *EncyclopediaShipsWeaponry          `json:"weaponry"`
	DiveBomber          *EncyclopediaShipsDiveBomber        `json:"dive_bomber"`
}
type EncyclopediaShipsValue struct {
	Max int `json:"max"`
	Min int `json:"min"`
}
type EncyclopediaShipsEngine struct {
	EngineIdStr string  `json:"engine_id_str"`
	MaxSpeed    float64 `json:"max_speed"`
	EngineId    int     `json:"engine_id"`
}
type EncyclopediaShipsTorpedoBomber struct {
	TorpedoDistance    float64                 `json:"torpedo_distance"`
	PlaneLevel         int                     `json:"plane_level"`
	Squadrons          int                     `json:"squadrons"`
	Name               string                  `json:"name"`
	CruiseSpeed        int                     `json:"cruise_speed"`
	PrepareTime        interface{}             `json:"prepare_time"`
	TorpedoDamage      int                     `json:"torpedo_damage"`
	CountInSquadron    *EncyclopediaShipsValue `json:"count_in_squadron"`
	TorpedoMaxSpeed    int                     `json:"torpedo_max_speed"`
	TorpedoBomberIdStr string                  `json:"torpedo_bomber_id_str"`
	GunnerDamage       int                     `json:"gunner_damage"`
	MaxDamage          int                     `json:"max_damage"`
	MaxHealth          int                     `json:"max_health"`
	TorpedoBomberId    int64                   `json:"torpedo_bomber_id"`
	TorpedoName        string                  `json:"torpedo_name"`
}
type EncyclopediaShipsAA struct {
	Slots   map[string]EncyclopediaShipsAASlot `json:"slots"`
	Defense int                                `json:"defense"`
}
type EncyclopediaShipsAASlot struct {
	Distance  float64 `json:"distance"`
	AvgDamage int     `json:"avg_damage"`
	Caliber   int     `json:"caliber"`
	Name      string  `json:"name"`
	Guns      int     `json:"guns"`
}
type EncyclopediaShipsDepthCharge struct {
	NumBombsInPack int     `json:"num_bombs_in_pack"`
	BombMaxDamage  int     `json:"bomb_max_damage"`
	MaxPacks       int     `json:"max_packs"`
	ReloadTime     float64 `json:"reload_time"`
}
type EncyclopediaShipsMobility struct {
	RudderTime    float64 `json:"rudder_time"`
	Total         int     `json:"total"`
	MaxSpeed      float64 `json:"max_speed"`
	TurningRadius int     `json:"turning_radius"`
}
type EncyclopediaShipsHull struct {
	HullId              int64                       `json:"hull_id"`
	HullIdStr           string                      `json:"hull_id_str"`
	TorpedoesBarrels    int                         `json:"torpedoes_barrels"`
	AntiAircraftBarrels int                         `json:"anti_aircraft_barrels"`
	Range               *EncyclopediaShipsHullRange `json:"range"`
	Health              int                         `json:"health"`
	PlanesAmount        interface{}                 `json:"planes_amount"`
	ArtilleryBarrels    int                         `json:"artillery_barrels"`
	AtbaBarrels         int                         `json:"atba_barrels"`
}
type EncyclopediaShipsHullRange struct {
	Max int `json:"max"`
	Min int `json:"min"`
}
type EncyclopediaShipsAtbas struct {
	Distance float64                               `json:"distance"`
	Slots    map[string]EncyclopediaShipsAtbasSlot `json:"slots"`
}
type EncyclopediaShipsAtbasSlot struct {
	BurnProbability float64 `json:"burn_probability"`
	BulletSpeed     int     `json:"bullet_speed"`
	Name            string  `json:"name"`
	ShotDelay       float64 `json:"shot_delay"`
	Damage          int     `json:"damage"`
	BulletMass      float64 `json:"bullet_mass"`
	Type            string  `json:"type"`
	GunRate         float64 `json:"gun_rate"`
}
type EncyclopediaShipsArtillery struct {
	MaxDispersion  int                                        `json:"max_dispersion"`
	Shells         map[string]EncyclopediaShipsArtilleryShell `json:"shells"`
	ShotDelay      float64                                    `json:"shot_delay"`
	RotationTime   float64                                    `json:"rotation_time"`
	Distance       float64                                    `json:"distance"`
	ArtilleryId    int64                                      `json:"artillery_id"`
	ArtilleryIdStr string                                     `json:"artillery_id_str"`
	Slots          map[string]EncyclopediaShipsArtillerySlot  `json:"slots"`
	GunRate        float64                                    `json:"gun_rate"`
}
type EncyclopediaShipsArtilleryShell struct {
	BurnProbability float64 `json:"burn_probability"`
	BulletSpeed     int     `json:"bullet_speed"`
	Name            string  `json:"name"`
	Damage          int     `json:"damage"`
	BulletMass      int     `json:"bullet_mass"`
	Type            string  `json:"type"`
}
type EncyclopediaShipsArtillerySlot struct {
	Barrels int    `json:"barrels"`
	Name    string `json:"name"`
	Guns    int    `json:"guns"`
}
type EncyclopediaShipsTorpedoes struct {
	VisibilityDist float64                                   `json:"visibility_dist"`
	Distance       float64                                   `json:"distance"`
	TorpedoesId    int64                                     `json:"torpedoes_id"`
	TorpedoName    string                                    `json:"torpedo_name"`
	ReloadTime     int                                       `json:"reload_time"`
	TorpedoSpeed   int                                       `json:"torpedo_speed"`
	RotationTime   float64                                   `json:"rotation_time"`
	TorpedoesIdStr string                                    `json:"torpedoes_id_str"`
	Slots          map[string]EncyclopediaShipsTorpedoesSlot `json:"slots"`
	MaxDamage      int                                       `json:"max_damage"`
}
type EncyclopediaShipsTorpedoesSlot struct {
	Barrels int    `json:"barrels"`
	Caliber int    `json:"caliber"`
	Name    string `json:"name"`
	Guns    int    `json:"guns"`
}
type EncyclopediaShipsFighters struct {
	FightersId      int64                   `json:"fighters_id"`
	Squadrons       int                     `json:"squadrons"`
	Name            string                  `json:"name"`
	CruiseSpeed     int                     `json:"cruise_speed"`
	PrepareTime     int                     `json:"prepare_time"`
	GunnerDamage    int                     `json:"gunner_damage"`
	FightersIdStr   string                  `json:"fighters_id_str"`
	CountInSquadron *EncyclopediaShipsValue `json:"count_in_squadron"`
	MaxAmmo         int                     `json:"max_ammo"`
	PlaneLevel      int                     `json:"plane_level"`
	AvgDamage       int                     `json:"avg_damage"`
	MaxHealth       int                     `json:"max_health"`
}
type EncyclopediaShipsSubmarineMobility struct {
	MaxSpeedUnderWater int     `json:"max_speed_under_water"`
	BuoyancyRudderTime float64 `json:"buoyancy_rudder_time"`
	Total              int     `json:"total"`
	MaxBuoyancySpeed   float64 `json:"max_buoyancy_speed"`
}
type EncyclopediaShipsFireControl struct {
	FireControlId    int64   `json:"fire_control_id"`
	Distance         float64 `json:"distance"`
	DistanceIncrease int     `json:"distance_increase"`
	FireControlIdStr string  `json:"fire_control_id_str"`
}
type EncyclopediaShipsSubmarineSonar struct {
	SubmarineSonarIdStr string  `json:"submarine_sonar_id_str"`
	WaveDuration0       int     `json:"wave_duration_0"`
	WaveDuration1       int     `json:"wave_duration_1"`
	WaveShotDelay       float64 `json:"wave_shot_delay"`
	SubmarineSonarId    int64   `json:"submarine_sonar_id"`
	WaveMaxDist         float64 `json:"wave_max_dist"`
	WaveSpeedMax        int     `json:"wave_speed_max"`
	WaveWidth           int     `json:"wave_width"`
	Total               int     `json:"total"`
}
type EncyclopediaShipsFlightControl struct {
	BomberSquadrons    int    `json:"bomber_squadrons"`
	FighterSquadrons   int    `json:"fighter_squadrons"`
	FlightControlId    int64  `json:"flight_control_id"`
	FlightControlIdStr string `json:"flight_control_id_str"`
	TorpedoSquadrons   int    `json:"torpedo_squadrons"`
}
type EncyclopediaShipsSubmarineBattery struct {
	MaxCapacity     float64 `json:"max_capacity"`
	RegenRate       float64 `json:"regen_rate"`
	Total           int     `json:"total"`
	ConsumptionRate float64 `json:"consumption_rate"`
}
type EncyclopediaShipsConcealment struct {
	Total                     int     `json:"total"`
	DetectDistanceByPlane     float64 `json:"detect_distance_by_plane"`
	DetectDistanceBySubmarine float64 `json:"detect_distance_by_submarine"`
	DetectDistanceByShip      float64 `json:"detect_distance_by_ship"`
}
type EncyclopediaShipsArmour struct {
	Casemate    *EncyclopediaShipsValue `json:"casemate"`
	Deck        *EncyclopediaShipsValue `json:"deck"`
	FloodDamage int                     `json:"flood_damage"`
	FloodProb   int                     `json:"flood_prob"`
	Range       *EncyclopediaShipsValue `json:"range"`
	Health      int                     `json:"health"`
	Extremities *EncyclopediaShipsValue `json:"extremities"`
	Total       int                     `json:"total"`
	Citadel     *EncyclopediaShipsValue `json:"citadel"`
}
type EncyclopediaShipsWeaponry struct {
	Torpedoes    int `json:"torpedoes"`
	Aircraft     int `json:"aircraft"`
	Artillery    int `json:"artillery"`
	AntiAircraft int `json:"anti_aircraft"`
}
type EncyclopediaShipsDiveBomber struct {
	Squadrons           int                     `json:"squadrons"`
	Name                string                  `json:"name"`
	CruiseSpeed         int                     `json:"cruise_speed"`
	DiveBomberId        int64                   `json:"dive_bomber_id"`
	PrepareTime         int                     `json:"prepare_time"`
	GunnerDamage        int                     `json:"gunner_damage"`
	BombDamage          int                     `json:"bomb_damage"`
	CountInSquadron     *EncyclopediaShipsValue `json:"count_in_squadron"`
	BombName            string                  `json:"bomb_name"`
	BombBulletMass      int                     `json:"bomb_bullet_mass"`
	PlaneLevel          int                     `json:"plane_level"`
	BombBurnProbability float64                 `json:"bomb_burn_probability"`
	MaxDamage           int                     `json:"max_damage"`
	MaxHealth           int                     `json:"max_health"`
	DiveBomberIdStr     string                  `json:"dive_bomber_id_str"`
	Accuracy            *EncyclopediaShipsValue `json:"accuracy"`
}

type EncyclopediaAchievementsData struct {
	BattleTypes    []string `json:"battle_types"`
	Multiple       int      `json:"multiple"`
	Name           string   `json:"name"`
	CountPerBattle string   `json:"count_per_battle"`
	MaxProgress    int      `json:"max_progress"`
	IsProgress     int      `json:"is_progress"`
	Image          string   `json:"image"`
	MaxShipLevel   int      `json:"max_ship_level"`
	MinShipLevel   int      `json:"min_ship_level"`
	AchievementId  string   `json:"achievement_id"`
	ImageInactive  string   `json:"image_inactive"`
	Hidden         int      `json:"hidden"`
	Reward         bool     `json:"reward"`
	Type           string   `json:"type"`
	SubType        string   `json:"sub_type"`
	Description    string   `json:"description"`
}

type EncyclopediaModulesData struct {
	Profile     EncyclopediaShipsHull `json:"profile"`
	Name        string                `json:"name"`
	Image       string                `json:"image"`
	Tag         string                `json:"tag"`
	ModuleIdStr string                `json:"module_id_str"`
	ModuleId    int64                 `json:"module_id"`
	Type        string                `json:"type"`
	PriceCredit int                   `json:"price_credit"`
}

type EncyclopediaCrewsData struct {
	BaseTrainingLevel      int                 `json:"base_training_level"`
	FirstNames             []string            `json:"first_names"`
	GoldTrainingHirePrice  int                 `json:"gold_training_hire_price"`
	GoldRetrainingPrice    int                 `json:"gold_retraining_price"`
	SubnationIndex         int                 `json:"subnation_index"`
	Icons                  []map[string]string `json:"icons"`
	BaseTrainingHirePrice  int                 `json:"base_training_hire_price"`
	MoneyTrainingHirePrice int                 `json:"money_training_hire_price"`
	Nation                 string              `json:"nation"`
	LastNames              []string            `json:"last_names"`
	MoneyRetrainingPrice   int                 `json:"money_retraining_price"`
	MoneyTrainingLevel     int                 `json:"money_training_level"`
	IsRetrainable          bool                `json:"is_retrainable"`
	GoldTrainingLevel      int                 `json:"gold_training_level"`
}

type EncyclopediaCrewSkillsData struct {
	TypeName      string                               `json:"type_name"`
	TypeId        int                                  `json:"type_id"`
	Name          string                               `json:"name"`
	Customization *EncyclopediaCrewSkillsCustomization `json:"customization"`
	Icon          string                               `json:"icon"`
}
type EncyclopediaCrewSkillsCustomization struct {
	Tier   int                                       `json:"tier"`
	Column int                                       `json:"column"`
	Perks  []EncyclopediaCrewSkillsCustomizationPerk `json:"perks"`
}
type EncyclopediaCrewSkillsCustomizationPerk struct {
	PerkId      int    `json:"perk_id"`
	Description string `json:"description"`
}

type EncyclopediaCrewRanksData struct {
	Experience int               `json:"experience"`
	Name       string            `json:"name"`
	Rank       int               `json:"rank"`
	Names      map[string]string `json:"names"`
}

type EncyclopediaAccountLevelsData struct {
	Tier   int    `json:"tier"`
	Image  string `json:"image"`
	Points int    `json:"points"`
}

type EncyclopediaBattleTypesData struct {
	Image       string `json:"image"`
	Tag         string `json:"tag"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type EncyclopediaConsumablesData struct {
	Profile      *EncyclopediaConsumablesProfile `json:"profile"`
	Name         string                          `json:"name"`
	PriceGold    int                             `json:"price_gold"`
	Image        string                          `json:"image"`
	ConsumableId int64                           `json:"consumable_id"`
	PriceCredit  int                             `json:"price_credit"`
	Type         string                          `json:"type"`
	Description  string                          `json:"description"`
}
type EncyclopediaConsumablesProfile struct {
	Description string  `json:"description"`
	Value       float64 `json:"value"`
}

type EncyclopediaCollectionsData struct {
	Name         string `json:"name"`
	Image        string `json:"image"`
	Tag          string `json:"tag"`
	CardCost     int    `json:"card_cost"`
	CollectionId int64  `json:"collection_id"`
	Description  string `json:"description"`
}

type EncyclopediaCollectionCardsData struct {
	Description  string                            `json:"description"`
	CardId       int64                             `json:"card_id"`
	Tag          string                            `json:"tag"`
	Images       *EncyclopediaCollectionCardsImage `json:"images"`
	CollectionId int64                             `json:"collection_id"`
	Name         string                            `json:"name"`
}
type EncyclopediaCollectionCardsImage struct {
	Small  string `json:"small"`
	Large  string `json:"large"`
	Medium string `json:"medium"`
}

type EncyclopediaBattleArenasData struct {
	Description   string `json:"description"`
	Icon          string `json:"icon"`
	BattleArenaId int    `json:"battle_arena_id"`
	Name          string `json:"name"`
}
