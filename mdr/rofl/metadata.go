package rofl

/// This file is the Go representation of the JSON schema.
// To make it please use the scripts in the scripts/ folder and go to https://app.quicktype.io/ to generate it from
// the JSON schema created with the scripts.

// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    metadata, err := UnmarshalMetadata(bytes)
//    bytes, err = metadata.Marshal()

import "encoding/json"

func UnmarshalMetadata(data []byte) (Metadata, error) {
	var r Metadata
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Metadata) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Metadata struct {
	Schema     string             `json:"$schema"`
	Title      string             `json:"title"`
	Type       string             `json:"type"`
	Properties MetadataProperties `json:"properties"`
	Required   []string           `json:"required"`
}

type MetadataProperties struct {
	GameLength      GameLength `json:"gameLength"`
	LastGameChunkID GameLength `json:"lastGameChunkId"`
	LastKeyFrameID  GameLength `json:"lastKeyFrameId"`
	StatsJSON       StatsJSON  `json:"statsJson"`
}

type GameLength struct {
	Type Type `json:"type"`
}

type StatsJSON struct {
	Type  string `json:"type"`
	Items Items  `json:"items"`
}

type Items struct {
	Type       string          `json:"type"`
	Properties ItemsProperties `json:"properties"`
	Required   []string        `json:"required"`
}

type ItemsProperties struct {
	TotalUnitsHealed                               GameLength    `json:"TOTAL_UNITS_HEALED"`
	HoLEliteKaiSaKillerInstinctKills               GameLength    `json:"HoL_Elite_KaiSaKillerInstinctKills"`
	S3A1Takedowns                                  GameLength    `json:"S3A1_Takedowns"`
	HoLKillsWhileLowHealth                         GameLength    `json:"HoL_KillsWhileLowHealth"`
	Spell2Cast                                     GameLength    `json:"SPELL2_CAST"`
	MissionsCrepeTakedownsWithInhibBuff            GameLength    `json:"Missions_Crepe_TakedownsWithInhibBuff"`
	Ping                                           GameLength    `json:"PING"`
	NeutralMinionsKilledEnemyJungle                GameLength    `json:"NEUTRAL_MINIONS_KILLED_ENEMY_JUNGLE"`
	PushPings                                      GameLength    `json:"PUSH_PINGS"`
	MagicDamageDealtToChampions                    GameLength    `json:"MAGIC_DAMAGE_DEALT_TO_CHAMPIONS"`
	The2026_S1A1SRRoleQuestComplete                GameLength    `json:"2026_S1A1_SR_RoleQuestComplete"`
	NodeNeutralizeAssist                           GameLength    `json:"NODE_NEUTRALIZE_ASSIST"`
	Perk0Var2                                      GameLength    `json:"PERK0_VAR2"`
	ObjectivesStolen                               GameLength    `json:"OBJECTIVES_STOLEN"`
	GoldSpent                                      GameLength    `json:"GOLD_SPENT"`
	PlayerScore11                                  GameLength    `json:"PLAYER_SCORE_11"`
	Perk4Var2                                      GameLength    `json:"PERK4_VAR2"`
	S3A1EventDoombotsTakenDownBefore5              GameLength    `json:"S3A1_Event_DoombotsTakenDownBefore5"`
	TripleKills                                    GameLength    `json:"TRIPLE_KILLS"`
	MissionsBXPEarnedPerGame                       GameLength    `json:"Missions_BXP_EarnedPerGame"`
	AtakhanKills                                   GameLength    `json:"ATAKHAN_KILLS"`
	PlayerAugment6                                 GameLength    `json:"PLAYER_AUGMENT_6"`
	AllInPings                                     GameLength    `json:"ALL_IN_PINGS"`
	LargestAttackDamage                            GameLength    `json:"LARGEST_ATTACK_DAMAGE"`
	MissionsTakedownEpicMonsters                   GameLength    `json:"Missions_TakedownEpicMonsters"`
	TotalHeal                                      GameLength    `json:"TOTAL_HEAL"`
	The2026_S1A1SRFaerieWards                      GameLength    `json:"2026_S1A1_SR_FaerieWards"`
	The2026_S1A1SkinsZiggs                         GameLength    `json:"2026_S1A1_Skins_Ziggs"`
	EventS1A2ArenaNoxianChampions                  GameLength    `json:"Event_S1_A2_Arena_NoxianChampions"`
	DemonsHandMissionPointsC                       GameLength    `json:"DemonsHand_MissionPointsC"`
	GameEndedInSurrender                           GameLength    `json:"GAME_ENDED_IN_SURRENDER"`
	TotalDamageDealtToTurrets                      GameLength    `json:"TOTAL_DAMAGE_DEALT_TO_TURRETS"`
	LongestTimeSpentLiving                         GameLength    `json:"LONGEST_TIME_SPENT_LIVING"`
	MissionsTakedownsAfterTeleporting              GameLength    `json:"Missions_TakedownsAfterTeleporting"`
	VisionClearedPings                             GameLength    `json:"VISION_CLEARED_PINGS"`
	WasLeaver                                      GameLength    `json:"WAS_LEAVER"`
	The2026_S1A1SkinsJayce                         GameLength    `json:"2026_S1A1_Skins_Jayce"`
	The2026_S1A1SkinsYuumi                         GameLength    `json:"2026_S1A1_Skins_Yuumi"`
	HoLEliteAsheCrystalArrowTakedowns              GameLength    `json:"HoL_Elite_AsheCrystalArrowTakedowns"`
	RiotIDTagLine                                  RiotIDTagLine `json:"RIOT_ID_TAG_LINE"`
	SummonSpell2Cast                               GameLength    `json:"SUMMON_SPELL2_CAST"`
	PentaKills                                     GameLength    `json:"PENTA_KILLS"`
	WardPlacedDetector                             GameLength    `json:"WARD_PLACED_DETECTOR"`
	TimeCcingOthers                                GameLength    `json:"TIME_CCING_OTHERS"`
	TimeOfFromLastDisconnect                       GameLength    `json:"TIME_OF_FROM_LAST_DISCONNECT"`
	HoLEliteEzrealTrueshotBarrageMultiHit          GameLength    `json:"HoL_Elite_EzrealTrueshotBarrageMultiHit"`
	S3A2PrismaticAug                               GameLength    `json:"S3A2_PrismaticAug"`
	DoubleKills                                    GameLength    `json:"DOUBLE_KILLS"`
	MissionsTakedownsWithHelpFromMonsters          GameLength    `json:"Missions_TakedownsWithHelpFromMonsters"`
	WeeklyMissionS2SpiritPetals                    GameLength    `json:"WeeklyMission_S2_SpiritPetals"`
	EventS1A2ArenaReviveAllies                     GameLength    `json:"Event_S1_A2_Arena_ReviveAllies"`
	LargestCriticalStrike                          GameLength    `json:"LARGEST_CRITICAL_STRIKE"`
	The2026_S1A1SkinsAshe                          GameLength    `json:"2026_S1A1_Skins_Ashe"`
	Perk0                                          GameLength    `json:"PERK0"`
	TrueDamageDealtPlayer                          GameLength    `json:"TRUE_DAMAGE_DEALT_PLAYER"`
	MagicDamageTaken                               GameLength    `json:"MAGIC_DAMAGE_TAKEN"`
	ChampionMissionStat1                           GameLength    `json:"CHAMPION_MISSION_STAT_1"`
	Win                                            GameLength    `json:"WIN"`
	PlayerPosition                                 GameLength    `json:"PLAYER_POSITION"`
	MissionsLegendaryItems                         GameLength    `json:"Missions_LegendaryItems"`
	HoLEnemyTakedownUnderTower                     GameLength    `json:"HoL_EnemyTakedownUnderTower"`
	PlayerScore6                                   GameLength    `json:"PLAYER_SCORE_6"`
	The2026_S1A1SkinsSeraphine                     GameLength    `json:"2026_S1A1_Skins_Seraphine"`
	DemonsHandMissionPointsE                       GameLength    `json:"DemonsHand_MissionPointsE"`
	MissionsPorosFed                               GameLength    `json:"Missions_PorosFed"`
	MissionsTakedownsBefore15Min                   GameLength    `json:"Missions_TakedownsBefore15Min"`
	Perk2                                          GameLength    `json:"PERK2"`
	TurretTakedowns                                GameLength    `json:"TURRET_TAKEDOWNS"`
	MissionsChampionsKilled                        GameLength    `json:"Missions_ChampionsKilled"`
	EventS2A2Exalted                               GameLength    `json:"Event_S2A2_Exalted"`
	MissionsTakedownEpicMonstersSingleGame         GameLength    `json:"Missions_TakedownEpicMonstersSingleGame"`
	Perk1Var3                                      GameLength    `json:"PERK1_VAR3"`
	Perk5                                          GameLength    `json:"PERK5"`
	HoLEliteVayneTumbleDodge                       GameLength    `json:"HoL_Elite_VayneTumbleDodge"`
	OnMyWayPings                                   GameLength    `json:"ON_MY_WAY_PINGS"`
	MissionsTimeSpentActivelyPlaying               GameLength    `json:"Missions_TimeSpentActivelyPlaying"`
	Perk3Var1                                      GameLength    `json:"PERK3_VAR1"`
	LargestMultiKill                               GameLength    `json:"LARGEST_MULTI_KILL"`
	Perk2Var3                                      GameLength    `json:"PERK2_VAR3"`
	Exp                                            GameLength    `json:"EXP"`
	Spell1Cast                                     GameLength    `json:"SPELL1_CAST"`
	PhysicalDamageDealtToChampions                 GameLength    `json:"PHYSICAL_DAMAGE_DEALT_TO_CHAMPIONS"`
	HoldPings                                      GameLength    `json:"HOLD_PINGS"`
	Item4                                          GameLength    `json:"ITEM4"`
	PlayersIMuted                                  GameLength    `json:"PLAYERS_I_MUTED"`
	MissionsChampionTakedownsWithIgnite            GameLength    `json:"Missions_ChampionTakedownsWithIgnite"`
	EventS1A2Mordekaiser                           GameLength    `json:"Event_S1_A2_Mordekaiser"`
	TrueDamageDealtToChampions                     GameLength    `json:"TRUE_DAMAGE_DEALT_TO_CHAMPIONS"`
	PhysicalDamageTaken                            GameLength    `json:"PHYSICAL_DAMAGE_TAKEN"`
	VictoryPointTotal                              GameLength    `json:"VICTORY_POINT_TOTAL"`
	TotalDamageDealt                               GameLength    `json:"TOTAL_DAMAGE_DEALT"`
	MissionsChampionTakedownsWhileGhosted          GameLength    `json:"Missions_ChampionTakedownsWhileGhosted"`
	The2026_S1A1SkinsCaitlyn                       GameLength    `json:"2026_S1A1_Skins_Caitlyn"`
	HqTakedowns                                    GameLength    `json:"HQ_TAKEDOWNS"`
	Perk3Var3                                      GameLength    `json:"PERK3_VAR3"`
	The2026_S1A1SkinsYasuo                         GameLength    `json:"2026_S1A1_Skins_Yasuo"`
	TeamPosition                                   GameLength    `json:"TEAM_POSITION"`
	EventARAMDocks                                 GameLength    `json:"Event_ARAM_Docks"`
	Perk5Var3                                      GameLength    `json:"PERK5_VAR3"`
	HqKilled                                       GameLength    `json:"HQ_KILLED"`
	DangerPings                                    GameLength    `json:"DANGER_PINGS"`
	PlayerScore8                                   GameLength    `json:"PLAYER_SCORE_8"`
	WardKilled                                     GameLength    `json:"WARD_KILLED"`
	EventS1A2EsportsTakedownEpicMonstersSingleGame GameLength    `json:"Event_S1_A2_Esports_TakedownEpicMonstersSingleGame"`
	ConsumablesPurchased                           GameLength    `json:"CONSUMABLES_PURCHASED"`
	PlayerScore0                                   GameLength    `json:"PLAYER_SCORE_0"`
	DragonKills                                    GameLength    `json:"DRAGON_KILLS"`
	MissionsTwoChampsKilledWithSameAbility         GameLength    `json:"Missions_TwoChampsKilledWithSameAbility"`
	VisionWardsBoughtInGame                        GameLength    `json:"VISION_WARDS_BOUGHT_IN_GAME"`
	PlayerScore1                                   GameLength    `json:"PLAYER_SCORE_1"`
	Item2                                          GameLength    `json:"ITEM2"`
	EventS1A1AprilFoolsDragon                      GameLength    `json:"Event_S1_A1_AprilFools_Dragon"`
	NodeNeutralize                                 GameLength    `json:"NODE_NEUTRALIZE"`
	MissionsChampionsHitWithAbilitiesEarlyGame     GameLength    `json:"Missions_ChampionsHitWithAbilitiesEarlyGame"`
	HoLFightsSurvivedWhileLowHealth                GameLength    `json:"HoL_FightsSurvivedWhileLowHealth"`
	Item3                                          GameLength    `json:"ITEM3"`
	TeamEarlySurrendered                           GameLength    `json:"TEAM_EARLY_SURRENDERED"`
	SummonerSpell1                                 GameLength    `json:"SUMMONER_SPELL_1"`
	BaronKills                                     GameLength    `json:"BARON_KILLS"`
	SightWardsBoughtInGame                         GameLength    `json:"SIGHT_WARDS_BOUGHT_IN_GAME"`
	VisionScore                                    GameLength    `json:"VISION_SCORE"`
	MissionsSorceryRune                            GameLength    `json:"Missions_SorceryRune"`
	The2026_S1A1SkinsBriar                         GameLength    `json:"2026_S1A1_Skins_Briar"`
	MissionsCrepeSnowballLanded                    GameLength    `json:"Missions_Crepe_SnowballLanded"`
	DemonsHandMissionPointsB                       GameLength    `json:"DemonsHand_MissionPointsB"`
	PerkPrimaryStyle                               GameLength    `json:"PERK_PRIMARY_STYLE"`
	EventBrawlMinions                              GameLength    `json:"Event_Brawl_Minions"`
	Perk1Var2                                      GameLength    `json:"PERK1_VAR2"`
	EnemyMissingPings                              GameLength    `json:"ENEMY_MISSING_PINGS"`
	TimePlayed                                     GameLength    `json:"TIME_PLAYED"`
	EventS1A2AprilFoolsDragon                      GameLength    `json:"Event_S1_A2_AprilFools_Dragon"`
	The2026_S1A1SkinsLillia                        GameLength    `json:"2026_S1A1_Skins_Lillia"`
	TotalDamageDealtToObjectives                   GameLength    `json:"TOTAL_DAMAGE_DEALT_TO_OBJECTIVES"`
	PlayerSubteamPlacement                         GameLength    `json:"PLAYER_SUBTEAM_PLACEMENT"`
	MissionsGoldFromStructuresDestroyed            GameLength    `json:"Missions_GoldFromStructuresDestroyed"`
	WasEarlySurrenderAccomplice                    GameLength    `json:"WAS_EARLY_SURRENDER_ACCOMPLICE"`
	Perk3                                          GameLength    `json:"PERK3"`
	EventS2A2ChampDamageAutos                      GameLength    `json:"Event_S2A2Champ_DamageAutos"`
	TotalDamageSelfMitigated                       GameLength    `json:"TOTAL_DAMAGE_SELF_MITIGATED"`
	Perk2Var1                                      GameLength    `json:"PERK2_VAR1"`
	MissionsVoidMitesSummoned                      GameLength    `json:"Missions_VoidMitesSummoned"`
	EventS2A2MV                                    GameLength    `json:"Event_S2A2_MV"`
	HoLJungleCampsStolen                           GameLength    `json:"HoL_JungleCampsStolen"`
	MagicDamageDealtPlayer                         GameLength    `json:"MAGIC_DAMAGE_DEALT_PLAYER"`
	ActMissionS1A2FeatsOfStrength                  GameLength    `json:"ActMission_S1_A2_FeatsOfStrength"`
	AssistMePings                                  GameLength    `json:"ASSIST_ME_PINGS"`
	MutedAll                                       GameLength    `json:"MUTED_ALL"`
	NodeCaptureAssist                              GameLength    `json:"NODE_CAPTURE_ASSIST"`
	ItemsPurchased                                 GameLength    `json:"ITEMS_PURCHASED"`
	Level                                          GameLength    `json:"LEVEL"`
	Perk0Var1                                      GameLength    `json:"PERK0_VAR1"`
	MissionsPrecisionRune                          GameLength    `json:"Missions_PrecisionRune"`
	HoLShutdownGoldCollected                       GameLength    `json:"HoL_ShutdownGoldCollected"`
	GameEndedInEarlySurrender                      GameLength    `json:"GAME_ENDED_IN_EARLY_SURRENDER"`
	WasAfk                                         GameLength    `json:"WAS_AFK"`
	NeutralMinionsKilledYourJungle                 GameLength    `json:"NEUTRAL_MINIONS_KILLED_YOUR_JUNGLE"`
	ChampionMissionStat3                           GameLength    `json:"CHAMPION_MISSION_STAT_3"`
	SeasonalMissionsTakedownAtakhan                GameLength    `json:"SeasonalMissions_TakedownAtakhan"`
	MissionsSnowballsHit                           GameLength    `json:"Missions_SnowballsHit"`
	Perk4Var1                                      GameLength    `json:"PERK4_VAR1"`
	MissionsTakedownDragons                        GameLength    `json:"Missions_TakedownDragons"`
	UnrealKills                                    GameLength    `json:"UNREAL_KILLS"`
	EventS1A1AprilFoolsSnowball                    GameLength    `json:"Event_S1_A1_AprilFools_Snowball"`
	TotalTimeCrowdControlDealt                     GameLength    `json:"TOTAL_TIME_CROWD_CONTROL_DEALT"`
	QuadraKills                                    GameLength    `json:"QUADRA_KILLS"`
	EventS1A2AprilFoolsGarenPlay                   GameLength    `json:"Event_S1_A2_AprilFools_Garen_Play"`
	MissionsCreepScoreBy10Minutes                  GameLength    `json:"Missions_CreepScoreBy10Minutes"`
	MissionsDamageToChampsWithItems                GameLength    `json:"Missions_DamageToChampsWithItems"`
	HoLHiddenEnemiesDamaged                        GameLength    `json:"HoL_HiddenEnemiesDamaged"`
	StatPerk1                                      GameLength    `json:"STAT_PERK_1"`
	HoLControlWardsKilled                          GameLength    `json:"HoL_ControlWardsKilled"`
	RiotIDGameName                                 GameLength    `json:"RIOT_ID_GAME_NAME"`
	ChampionsKilled                                GameLength    `json:"CHAMPIONS_KILLED"`
	Perk1Var1                                      GameLength    `json:"PERK1_VAR1"`
	Perk4Var3                                      GameLength    `json:"PERK4_VAR3"`
	S3A2ZaahenUnlock                               GameLength    `json:"S3A2_ZaahenUnlock"`
	The2026_S1A1SkinsCamille                       GameLength    `json:"2026_S1A1_Skins_Camille"`
	SummonSpell1Cast                               GameLength    `json:"SUMMON_SPELL1_CAST"`
	SummonerSpell2                                 GameLength    `json:"SUMMONER_SPELL_2"`
	Event2025LRStructuresEpicMonsters              GameLength    `json:"Event_2025LR_StructuresEpicMonsters"`
	WeeklyMissionS2FeatsOfStrength                 GameLength    `json:"WeeklyMission_S2_FeatsOfStrength"`
	HoLEliteLucianCullingHits                      GameLength    `json:"HoL_Elite_LucianCullingHits"`
	GetBackPings                                   GameLength    `json:"GET_BACK_PINGS"`
	Team                                           GameLength    `json:"TEAM"`
	PlayersThatMutedMe                             GameLength    `json:"PLAYERS_THAT_MUTED_ME"`
	ID                                             GameLength    `json:"ID"`
	TeamObjective                                  GameLength    `json:"TEAM_OBJECTIVE"`
	TotalDamageTaken                               GameLength    `json:"TOTAL_DAMAGE_TAKEN"`
	ActMissionS1A2ArenaRoundsWon                   GameLength    `json:"ActMission_S1_A2_ArenaRoundsWon"`
	Item1                                          GameLength    `json:"ITEM1"`
	StatPerk2                                      GameLength    `json:"STAT_PERK_2"`
	PlayerScore10                                  GameLength    `json:"PLAYER_SCORE_10"`
	The2026_S1A1SkinsKatarina                      GameLength    `json:"2026_S1A1_Skins_Katarina"`
	WeeklyMissionS2DamagingAbilities               GameLength    `json:"WeeklyMission_S2_DamagingAbilities"`
	RetreatPings                                   GameLength    `json:"RETREAT_PINGS"`
	Puuid                                          GameLength    `json:"PUUID"`
	MissionsTrueDamageToStructures                 GameLength    `json:"Missions_TrueDamageToStructures"`
	GoldEarned                                     GameLength    `json:"GOLD_EARNED"`
	EventARAMHexgates                              GameLength    `json:"Event_ARAM_Hexgates"`
	HoLSoloKills                                   GameLength    `json:"HoL_SoloKills"`
	HoLTurretsTakenWithinMinutes                   GameLength    `json:"HoL_TurretsTakenWithinMinutes"`
	The2026_S1A1SkinsNautilus                      GameLength    `json:"2026_S1A1_Skins_Nautilus"`
	KillingSprees                                  GameLength    `json:"KILLING_SPREES"`
	MinionsKilled                                  GameLength    `json:"MINIONS_KILLED"`
	EnemyVisionPings                               GameLength    `json:"ENEMY_VISION_PINGS"`
	HoLEliteLucianPiercingLightMultiHit            GameLength    `json:"HoL_Elite_LucianPiercingLightMultiHit"`
	MissionsPlaceUsefulWards                       GameLength    `json:"Missions_PlaceUsefulWards"`
	MissionsTakedownStructures                     GameLength    `json:"Missions_TakedownStructures"`
	LastTakedownTime                               GameLength    `json:"LAST_TAKEDOWN_TIME"`
	Perk3Var2                                      GameLength    `json:"PERK3_VAR2"`
	Perk4                                          GameLength    `json:"PERK4"`
	StatPerk0                                      GameLength    `json:"STAT_PERK_0"`
	DemonsHandMissionPointsA                       GameLength    `json:"DemonsHand_MissionPointsA"`
	TimeSpentDisconnected                          GameLength    `json:"TIME_SPENT_DISCONNECTED"`
	Perk2Var2                                      GameLength    `json:"PERK2_VAR2"`
	IndividualPosition                             GameLength    `json:"INDIVIDUAL_POSITION"`
	FriendlyDampenLost                             GameLength    `json:"FRIENDLY_DAMPEN_LOST"`
	HordeKills                                     GameLength    `json:"HORDE_KILLS"`
	MissionsPlaceUsefulControlWards                GameLength    `json:"Missions_PlaceUsefulControlWards"`
	Perk5Var1                                      GameLength    `json:"PERK5_VAR1"`
	Skin                                           GameLength    `json:"SKIN"`
	NodeCapture                                    GameLength    `json:"NODE_CAPTURE"`
	FriendlyHqLost                                 GameLength    `json:"FRIENDLY_HQ_LOST"`
	MissionsInspirationRune                        GameLength    `json:"Missions_InspirationRune"`
	TotalDamageShieldedOnTeammates                 GameLength    `json:"TOTAL_DAMAGE_SHIELDED_ON_TEAMMATES"`
	WardPlaced                                     GameLength    `json:"WARD_PLACED"`
	Item5                                          GameLength    `json:"ITEM5"`
	PlayerScore5                                   GameLength    `json:"PLAYER_SCORE_5"`
	HoLEliteEzrealEssenceFluxDetonated             GameLength    `json:"HoL_Elite_EzrealEssenceFluxDetonated"`
	MissionsPeriodicDamage                         GameLength    `json:"Missions_PeriodicDamage"`
	NeutralMinionsKilled                           GameLength    `json:"NEUTRAL_MINIONS_KILLED"`
	The2026_S1A1SRGrowthSmashed                    GameLength    `json:"2026_S1A1_SR_GrowthSmashed"`
	BarracksTakedowns                              GameLength    `json:"BARRACKS_TAKEDOWNS"`
	S3A1PlayAsDemaciansOrAgainstNoxians            GameLength    `json:"S3A1_PlayAsDemaciansOrAgainstNoxians"`
	TrueDamageTaken                                GameLength    `json:"TRUE_DAMAGE_TAKEN"`
	MissionsTakedownsAfterExhausting               GameLength    `json:"Missions_TakedownsAfterExhausting"`
	BasicPings                                     GameLength    `json:"BASIC_PINGS"`
	EventS2A2ChampDamageAbilities                  GameLength    `json:"Event_S2A2Champ_DamageAbilities"`
	MissionsDamageToStructures                     GameLength    `json:"Missions_DamageToStructures"`
	PlayerScore3                                   GameLength    `json:"PLAYER_SCORE_3"`
	MissionsDominationRune                         GameLength    `json:"Missions_DominationRune"`
	TotalDamageDealtToBuildings                    GameLength    `json:"TOTAL_DAMAGE_DEALT_TO_BUILDINGS"`
	PlayerAugment5                                 GameLength    `json:"PLAYER_AUGMENT_5"`
	The2026_S1A1SkinsSamira                        GameLength    `json:"2026_S1A1_Skins_Samira"`
	ObjectivesStolenAssists                        GameLength    `json:"OBJECTIVES_STOLEN_ASSISTS"`
	The2026_S1A1SkinsPoppy                         GameLength    `json:"2026_S1A1_Skins_Poppy"`
	PlayerScore2                                   GameLength    `json:"PLAYER_SCORE_2"`
	EventS1A2AprilFoolsGarenTakedown               GameLength    `json:"Event_S1_A2_AprilFools_Garen_Takedown"`
	FriendlyTurretLost                             GameLength    `json:"FRIENDLY_TURRET_LOST"`
	PlayerScore4                                   GameLength    `json:"PLAYER_SCORE_4"`
	EventBrawlJungle                               GameLength    `json:"Event_Brawl_Jungle"`
	MissionsCrepeDamageDealtSpeedZone              GameLength    `json:"Missions_Crepe_DamageDealtSpeedZone"`
	MissionsCannonMinionsKilled                    GameLength    `json:"Missions_CannonMinionsKilled"`
	Perk0Var3                                      GameLength    `json:"PERK0_VAR3"`
	HoLEliteVayneCondemnStun                       GameLength    `json:"HoL_Elite_VayneCondemnStun"`
	DemonsHandMissionPointsF                       GameLength    `json:"DemonsHand_MissionPointsF"`
	PlayerAugment1                                 GameLength    `json:"PLAYER_AUGMENT_1"`
	TotalDamageDealtToChampions                    GameLength    `json:"TOTAL_DAMAGE_DEALT_TO_CHAMPIONS"`
	TotalDamageDealtToEpicMonsters                 GameLength    `json:"TOTAL_DAMAGE_DEALT_TO_EPIC_MONSTERS"`
	TotalTimeSpentDead                             GameLength    `json:"TOTAL_TIME_SPENT_DEAD"`
	TotalHealOnTeammates                           GameLength    `json:"TOTAL_HEAL_ON_TEAMMATES"`
	DemonsHandMissionPointsD                       GameLength    `json:"DemonsHand_MissionPointsD"`
	PlayerAugment4                                 GameLength    `json:"PLAYER_AUGMENT_4"`
	WasAfkAfterFailedSurrender                     GameLength    `json:"WAS_AFK_AFTER_FAILED_SURRENDER"`
	PlayerAugment3                                 GameLength    `json:"PLAYER_AUGMENT_3"`
	HoLEliteAsheHawkshotChampsRevealed             GameLength    `json:"HoL_Elite_AsheHawkshotChampsRevealed"`
	LargestKillingSpree                            GameLength    `json:"LARGEST_KILLING_SPREE"`
	ChampionTransform                              GameLength    `json:"CHAMPION_TRANSFORM"`
	WasSurrenderDueToAfk                           GameLength    `json:"WAS_SURRENDER_DUE_TO_AFK"`
	HoLOutnumberedTakedowns                        GameLength    `json:"HoL_OutnumberedTakedowns"`
	EventS1A2ArenaBraveryChampions                 GameLength    `json:"Event_S1_A2_Arena_BraveryChampions"`
	MissionsDestroyPlants                          GameLength    `json:"Missions_DestroyPlants"`
	PlayerScore7                                   GameLength    `json:"PLAYER_SCORE_7"`
	MissionsTotalGold                              GameLength    `json:"Missions_TotalGold"`
	EventS2A2PetalPoints                           GameLength    `json:"Event_S2A2_PetalPoints"`
	LargestAbilityDamage                           GameLength    `json:"LARGEST_ABILITY_DAMAGE"`
	HoLChampionsDamagedWhileHidden                 GameLength    `json:"HoL_ChampionsDamagedWhileHidden"`
	Name                                           GameLength    `json:"NAME"`
	MissionsHealingFromLevelObjects                GameLength    `json:"Missions_HealingFromLevelObjects"`
	Perk1                                          GameLength    `json:"PERK1"`
	ChampionMissionStat2                           GameLength    `json:"CHAMPION_MISSION_STAT_2"`
	The2026_S1A1SkinsGalio                         GameLength    `json:"2026_S1A1_Skins_Galio"`
	Item0                                          GameLength    `json:"ITEM0"`
	MissionsMinionsKilled                          GameLength    `json:"Missions_MinionsKilled"`
	MissionsGoldPerMinute                          GameLength    `json:"Missions_GoldPerMinute"`
	Assists                                        GameLength    `json:"ASSISTS"`
	RiftHeraldKills                                GameLength    `json:"RIFT_HERALD_KILLS"`
	SummonerID                                     GameLength    `json:"SUMMONER_ID"`
	PlayerScore9                                   GameLength    `json:"PLAYER_SCORE_9"`
	MissionsResolveRune                            GameLength    `json:"Missions_ResolveRune"`
	MissionsTakedownsUnderTurret                   GameLength    `json:"Missions_TakedownsUnderTurret"`
	NumDeaths                                      GameLength    `json:"NUM_DEATHS"`
	TurretsKilled                                  GameLength    `json:"TURRETS_KILLED"`
	MissionsTakedownGold                           GameLength    `json:"Missions_TakedownGold"`
	PerkSubStyle                                   GameLength    `json:"PERK_SUB_STYLE"`
	PhysicalDamageDealtPlayer                      GameLength    `json:"PHYSICAL_DAMAGE_DEALT_PLAYER"`
	KeystoneID                                     GameLength    `json:"KEYSTONE_ID"`
	NeedVisionPings                                GameLength    `json:"NEED_VISION_PINGS"`
	ChampionMissionStat0                           GameLength    `json:"CHAMPION_MISSION_STAT_0"`
	TotalTimeCrowdControlDealtToChampions          GameLength    `json:"TOTAL_TIME_CROWD_CONTROL_DEALT_TO_CHAMPIONS"`
	Perk5Var2                                      GameLength    `json:"PERK5_VAR2"`
	MissionsTakedownWards                          GameLength    `json:"Missions_TakedownWards"`
	PlayerSubteam                                  GameLength    `json:"PLAYER_SUBTEAM"`
	MissionsHexgatesUsed                           GameLength    `json:"Missions_HexgatesUsed"`
	EventS1A2AprilFoolsSnowball                    GameLength    `json:"Event_S1_A2_AprilFools_Snowball"`
	PlayerAugment2                                 GameLength    `json:"PLAYER_AUGMENT_2"`
	MissionsTurretPlatesDestroyed                  GameLength    `json:"Missions_TurretPlatesDestroyed"`
	MissionsCreepScore                             GameLength    `json:"Missions_CreepScore"`
	HoLEliteKaiSaAbilitiesUpgraded                 GameLength    `json:"HoL_Elite_KaiSaAbilitiesUpgraded"`
	Spell4Cast                                     GameLength    `json:"SPELL4_CAST"`
	PlayerRole                                     GameLength    `json:"PLAYER_ROLE"`
	Spell3Cast                                     GameLength    `json:"SPELL3_CAST"`
	MissionsGoldFromTurretPlatesTaken              GameLength    `json:"Missions_GoldFromTurretPlatesTaken"`
	CommandPings                                   GameLength    `json:"COMMAND_PINGS"`
	MissionsImmobilizeChampions                    GameLength    `json:"Missions_ImmobilizeChampions"`
	MissionsTakedownBaronsElderDragons             GameLength    `json:"Missions_TakedownBaronsElderDragons"`
	The2026_S1A1SkinsOrnn                          GameLength    `json:"2026_S1A1_Skins_Ornn"`
	ActMissionS1A2BloodyPetalsCollected            GameLength    `json:"ActMission_S1_A2_BloodyPetalsCollected"`
	Item6                                          GameLength    `json:"ITEM6"`
	BarracksKilled                                 GameLength    `json:"BARRACKS_KILLED"`
}

type RiotIDTagLine struct {
	Type []Type `json:"type"`
}

type Type string

const (
	Integer Type = "integer"
	String  Type = "string"
)
