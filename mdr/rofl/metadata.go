package rofl

/// This file is the Go representation of the JSON schema.
// To make it please use the scripts in the scripts/ folder and go to https://app.quicktype.io/ to generate it from
// the JSON schema created with the scripts.
// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    metadata, err := UnmarshalMetadata(bytes)
//    bytes, err = metadata.Marshal()

import "bytes"
import "errors"

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
	GameLength      int64       `json:"gameLength"`
	LastGameChunkID int64       `json:"lastGameChunkId"`
	LastKeyFrameID  int64       `json:"lastKeyFrameId"`
	StatsJSON       []StatsJSON `json:"statsJson"`
}

type StatsJSON struct {
	The2026_S1A1SkinsAshe                          int64          `json:"2026_S1A1_Skins_Ashe"`
	The2026_S1A1SkinsBriar                         int64          `json:"2026_S1A1_Skins_Briar"`
	The2026_S1A1SkinsCaitlyn                       int64          `json:"2026_S1A1_Skins_Caitlyn"`
	The2026_S1A1SkinsCamille                       int64          `json:"2026_S1A1_Skins_Camille"`
	The2026_S1A1SkinsGalio                         int64          `json:"2026_S1A1_Skins_Galio"`
	The2026_S1A1SkinsJayce                         int64          `json:"2026_S1A1_Skins_Jayce"`
	The2026_S1A1SkinsKatarina                      int64          `json:"2026_S1A1_Skins_Katarina"`
	The2026_S1A1SkinsLillia                        int64          `json:"2026_S1A1_Skins_Lillia"`
	The2026_S1A1SkinsNautilus                      int64          `json:"2026_S1A1_Skins_Nautilus"`
	The2026_S1A1SkinsOrnn                          int64          `json:"2026_S1A1_Skins_Ornn"`
	The2026_S1A1SkinsPoppy                         int64          `json:"2026_S1A1_Skins_Poppy"`
	The2026_S1A1SkinsSamira                        int64          `json:"2026_S1A1_Skins_Samira"`
	The2026_S1A1SkinsSeraphine                     int64          `json:"2026_S1A1_Skins_Seraphine"`
	The2026_S1A1SkinsYasuo                         int64          `json:"2026_S1A1_Skins_Yasuo"`
	The2026_S1A1SkinsYuumi                         int64          `json:"2026_S1A1_Skins_Yuumi"`
	The2026_S1A1SkinsZiggs                         int64          `json:"2026_S1A1_Skins_Ziggs"`
	The2026_S1A1SRFaerieWards                      int64          `json:"2026_S1A1_SR_FaerieWards"`
	The2026_S1A1SRGrowthSmashed                    int64          `json:"2026_S1A1_SR_GrowthSmashed"`
	The2026_S1A1SRRoleQuestComplete                int64          `json:"2026_S1A1_SR_RoleQuestComplete"`
	ActMissionS1A2ArenaRoundsWon                   int64          `json:"ActMission_S1_A2_ArenaRoundsWon"`
	ActMissionS1A2BloodyPetalsCollected            int64          `json:"ActMission_S1_A2_BloodyPetalsCollected"`
	ActMissionS1A2FeatsOfStrength                  int64          `json:"ActMission_S1_A2_FeatsOfStrength"`
	AllInPings                                     int64          `json:"ALL_IN_PINGS"`
	AssistMePings                                  int64          `json:"ASSIST_ME_PINGS"`
	Assists                                        int64          `json:"ASSISTS"`
	AtakhanKills                                   int64          `json:"ATAKHAN_KILLS"`
	BaronKills                                     int64          `json:"BARON_KILLS"`
	BarracksKilled                                 int64          `json:"BARRACKS_KILLED"`
	BarracksTakedowns                              int64          `json:"BARRACKS_TAKEDOWNS"`
	BasicPings                                     int64          `json:"BASIC_PINGS"`
	ChampionMissionStat0                           int64          `json:"CHAMPION_MISSION_STAT_0"`
	ChampionMissionStat1                           int64          `json:"CHAMPION_MISSION_STAT_1"`
	ChampionMissionStat2                           int64          `json:"CHAMPION_MISSION_STAT_2"`
	ChampionMissionStat3                           int64          `json:"CHAMPION_MISSION_STAT_3"`
	ChampionTransform                              int64          `json:"CHAMPION_TRANSFORM"`
	ChampionsKilled                                int64          `json:"CHAMPIONS_KILLED"`
	CommandPings                                   int64          `json:"COMMAND_PINGS"`
	ConsumablesPurchased                           int64          `json:"CONSUMABLES_PURCHASED"`
	DangerPings                                    int64          `json:"DANGER_PINGS"`
	DemonsHandMissionPointsA                       int64          `json:"DemonsHand_MissionPointsA"`
	DemonsHandMissionPointsB                       int64          `json:"DemonsHand_MissionPointsB"`
	DemonsHandMissionPointsC                       int64          `json:"DemonsHand_MissionPointsC"`
	DemonsHandMissionPointsD                       int64          `json:"DemonsHand_MissionPointsD"`
	DemonsHandMissionPointsE                       int64          `json:"DemonsHand_MissionPointsE"`
	DemonsHandMissionPointsF                       int64          `json:"DemonsHand_MissionPointsF"`
	DoubleKills                                    int64          `json:"DOUBLE_KILLS"`
	DragonKills                                    int64          `json:"DRAGON_KILLS"`
	EnemyMissingPings                              int64          `json:"ENEMY_MISSING_PINGS"`
	EnemyVisionPings                               int64          `json:"ENEMY_VISION_PINGS"`
	Event2025LRStructuresEpicMonsters              int64          `json:"Event_2025LR_StructuresEpicMonsters"`
	EventARAMDocks                                 int64          `json:"Event_ARAM_Docks"`
	EventARAMHexgates                              int64          `json:"Event_ARAM_Hexgates"`
	EventBrawlJungle                               int64          `json:"Event_Brawl_Jungle"`
	EventBrawlMinions                              int64          `json:"Event_Brawl_Minions"`
	EventS1A1AprilFoolsDragon                      int64          `json:"Event_S1_A1_AprilFools_Dragon"`
	EventS1A1AprilFoolsSnowball                    int64          `json:"Event_S1_A1_AprilFools_Snowball"`
	EventS1A2AprilFoolsDragon                      int64          `json:"Event_S1_A2_AprilFools_Dragon"`
	EventS1A2AprilFoolsGarenPlay                   int64          `json:"Event_S1_A2_AprilFools_Garen_Play"`
	EventS1A2AprilFoolsGarenTakedown               int64          `json:"Event_S1_A2_AprilFools_Garen_Takedown"`
	EventS1A2AprilFoolsSnowball                    int64          `json:"Event_S1_A2_AprilFools_Snowball"`
	EventS1A2ArenaBraveryChampions                 int64          `json:"Event_S1_A2_Arena_BraveryChampions"`
	EventS1A2ArenaNoxianChampions                  int64          `json:"Event_S1_A2_Arena_NoxianChampions"`
	EventS1A2ArenaReviveAllies                     int64          `json:"Event_S1_A2_Arena_ReviveAllies"`
	EventS1A2EsportsTakedownEpicMonstersSingleGame int64          `json:"Event_S1_A2_Esports_TakedownEpicMonstersSingleGame"`
	EventS1A2Mordekaiser                           int64          `json:"Event_S1_A2_Mordekaiser"`
	EventS2A2Exalted                               int64          `json:"Event_S2A2_Exalted"`
	EventS2A2MV                                    int64          `json:"Event_S2A2_MV"`
	EventS2A2PetalPoints                           int64          `json:"Event_S2A2_PetalPoints"`
	EventS2A2ChampDamageAbilities                  int64          `json:"Event_S2A2Champ_DamageAbilities"`
	EventS2A2ChampDamageAutos                      int64          `json:"Event_S2A2Champ_DamageAutos"`
	Exp                                            int64          `json:"EXP"`
	FriendlyDampenLost                             int64          `json:"FRIENDLY_DAMPEN_LOST"`
	FriendlyHqLost                                 int64          `json:"FRIENDLY_HQ_LOST"`
	FriendlyTurretLost                             int64          `json:"FRIENDLY_TURRET_LOST"`
	GameEndedInEarlySurrender                      int64          `json:"GAME_ENDED_IN_EARLY_SURRENDER"`
	GameEndedInSurrender                           int64          `json:"GAME_ENDED_IN_SURRENDER"`
	GetBackPings                                   int64          `json:"GET_BACK_PINGS"`
	GoldEarned                                     int64          `json:"GOLD_EARNED"`
	GoldSpent                                      int64          `json:"GOLD_SPENT"`
	HoLChampionsDamagedWhileHidden                 int64          `json:"HoL_ChampionsDamagedWhileHidden"`
	HoLControlWardsKilled                          int64          `json:"HoL_ControlWardsKilled"`
	HoLEliteAsheCrystalArrowTakedowns              int64          `json:"HoL_Elite_AsheCrystalArrowTakedowns"`
	HoLEliteAsheHawkshotChampsRevealed             int64          `json:"HoL_Elite_AsheHawkshotChampsRevealed"`
	HoLEliteEzrealEssenceFluxDetonated             int64          `json:"HoL_Elite_EzrealEssenceFluxDetonated"`
	HoLEliteEzrealTrueshotBarrageMultiHit          int64          `json:"HoL_Elite_EzrealTrueshotBarrageMultiHit"`
	HoLEliteKaiSaAbilitiesUpgraded                 int64          `json:"HoL_Elite_KaiSaAbilitiesUpgraded"`
	HoLEliteKaiSaKillerInstinctKills               int64          `json:"HoL_Elite_KaiSaKillerInstinctKills"`
	HoLEliteLucianCullingHits                      int64          `json:"HoL_Elite_LucianCullingHits"`
	HoLEliteLucianPiercingLightMultiHit            int64          `json:"HoL_Elite_LucianPiercingLightMultiHit"`
	HoLEliteVayneCondemnStun                       int64          `json:"HoL_Elite_VayneCondemnStun"`
	HoLEliteVayneTumbleDodge                       int64          `json:"HoL_Elite_VayneTumbleDodge"`
	HoLEnemyTakedownUnderTower                     int64          `json:"HoL_EnemyTakedownUnderTower"`
	HoLFightsSurvivedWhileLowHealth                int64          `json:"HoL_FightsSurvivedWhileLowHealth"`
	HoLHiddenEnemiesDamaged                        int64          `json:"HoL_HiddenEnemiesDamaged"`
	HoLJungleCampsStolen                           int64          `json:"HoL_JungleCampsStolen"`
	HoLKillsWhileLowHealth                         int64          `json:"HoL_KillsWhileLowHealth"`
	HoLOutnumberedTakedowns                        int64          `json:"HoL_OutnumberedTakedowns"`
	HoLShutdownGoldCollected                       int64          `json:"HoL_ShutdownGoldCollected"`
	HoLSoloKills                                   int64          `json:"HoL_SoloKills"`
	HoLTurretsTakenWithinMinutes                   int64          `json:"HoL_TurretsTakenWithinMinutes"`
	HoldPings                                      int64          `json:"HOLD_PINGS"`
	HordeKills                                     int64          `json:"HORDE_KILLS"`
	HqKilled                                       int64          `json:"HQ_KILLED"`
	HqTakedowns                                    int64          `json:"HQ_TAKEDOWNS"`
	ID                                             int64          `json:"ID"`
	IndividualPosition                             string         `json:"INDIVIDUAL_POSITION"`
	Item0                                          int64          `json:"ITEM0"`
	Item1                                          int64          `json:"ITEM1"`
	Item2                                          int64          `json:"ITEM2"`
	Item3                                          int64          `json:"ITEM3"`
	Item4                                          int64          `json:"ITEM4"`
	Item5                                          int64          `json:"ITEM5"`
	Item6                                          int64          `json:"ITEM6"`
	ItemsPurchased                                 int64          `json:"ITEMS_PURCHASED"`
	KeystoneID                                     int64          `json:"KEYSTONE_ID"`
	KillingSprees                                  int64          `json:"KILLING_SPREES"`
	LargestAbilityDamage                           int64          `json:"LARGEST_ABILITY_DAMAGE"`
	LargestAttackDamage                            int64          `json:"LARGEST_ATTACK_DAMAGE"`
	LargestCriticalStrike                          int64          `json:"LARGEST_CRITICAL_STRIKE"`
	LargestKillingSpree                            int64          `json:"LARGEST_KILLING_SPREE"`
	LargestMultiKill                               int64          `json:"LARGEST_MULTI_KILL"`
	LastTakedownTime                               int64          `json:"LAST_TAKEDOWN_TIME"`
	Level                                          int64          `json:"LEVEL"`
	LongestTimeSpentLiving                         int64          `json:"LONGEST_TIME_SPENT_LIVING"`
	MagicDamageDealtPlayer                         int64          `json:"MAGIC_DAMAGE_DEALT_PLAYER"`
	MagicDamageDealtToChampions                    int64          `json:"MAGIC_DAMAGE_DEALT_TO_CHAMPIONS"`
	MagicDamageTaken                               int64          `json:"MAGIC_DAMAGE_TAKEN"`
	MinionsKilled                                  int64          `json:"MINIONS_KILLED"`
	MissionsBXPEarnedPerGame                       int64          `json:"Missions_BXP_EarnedPerGame"`
	MissionsCannonMinionsKilled                    int64          `json:"Missions_CannonMinionsKilled"`
	MissionsChampionsHitWithAbilitiesEarlyGame     int64          `json:"Missions_ChampionsHitWithAbilitiesEarlyGame"`
	MissionsChampionsKilled                        int64          `json:"Missions_ChampionsKilled"`
	MissionsChampionTakedownsWhileGhosted          int64          `json:"Missions_ChampionTakedownsWhileGhosted"`
	MissionsChampionTakedownsWithIgnite            int64          `json:"Missions_ChampionTakedownsWithIgnite"`
	MissionsCreepScore                             int64          `json:"Missions_CreepScore"`
	MissionsCreepScoreBy10Minutes                  int64          `json:"Missions_CreepScoreBy10Minutes"`
	MissionsCrepeDamageDealtSpeedZone              int64          `json:"Missions_Crepe_DamageDealtSpeedZone"`
	MissionsCrepeSnowballLanded                    int64          `json:"Missions_Crepe_SnowballLanded"`
	MissionsCrepeTakedownsWithInhibBuff            int64          `json:"Missions_Crepe_TakedownsWithInhibBuff"`
	MissionsDamageToChampsWithItems                int64          `json:"Missions_DamageToChampsWithItems"`
	MissionsDamageToStructures                     int64          `json:"Missions_DamageToStructures"`
	MissionsDestroyPlants                          int64          `json:"Missions_DestroyPlants"`
	MissionsDominationRune                         int64          `json:"Missions_DominationRune"`
	MissionsGoldFromStructuresDestroyed            int64          `json:"Missions_GoldFromStructuresDestroyed"`
	MissionsGoldFromTurretPlatesTaken              int64          `json:"Missions_GoldFromTurretPlatesTaken"`
	MissionsGoldPerMinute                          int64          `json:"Missions_GoldPerMinute"`
	MissionsHealingFromLevelObjects                int64          `json:"Missions_HealingFromLevelObjects"`
	MissionsHexgatesUsed                           int64          `json:"Missions_HexgatesUsed"`
	MissionsImmobilizeChampions                    int64          `json:"Missions_ImmobilizeChampions"`
	MissionsInspirationRune                        int64          `json:"Missions_InspirationRune"`
	MissionsLegendaryItems                         int64          `json:"Missions_LegendaryItems"`
	MissionsMinionsKilled                          int64          `json:"Missions_MinionsKilled"`
	MissionsPeriodicDamage                         int64          `json:"Missions_PeriodicDamage"`
	MissionsPlaceUsefulControlWards                int64          `json:"Missions_PlaceUsefulControlWards"`
	MissionsPlaceUsefulWards                       int64          `json:"Missions_PlaceUsefulWards"`
	MissionsPorosFed                               int64          `json:"Missions_PorosFed"`
	MissionsPrecisionRune                          int64          `json:"Missions_PrecisionRune"`
	MissionsResolveRune                            int64          `json:"Missions_ResolveRune"`
	MissionsSnowballsHit                           int64          `json:"Missions_SnowballsHit"`
	MissionsSorceryRune                            int64          `json:"Missions_SorceryRune"`
	MissionsTakedownBaronsElderDragons             int64          `json:"Missions_TakedownBaronsElderDragons"`
	MissionsTakedownDragons                        int64          `json:"Missions_TakedownDragons"`
	MissionsTakedownEpicMonsters                   int64          `json:"Missions_TakedownEpicMonsters"`
	MissionsTakedownEpicMonstersSingleGame         int64          `json:"Missions_TakedownEpicMonstersSingleGame"`
	MissionsTakedownGold                           int64          `json:"Missions_TakedownGold"`
	MissionsTakedownsAfterExhausting               int64          `json:"Missions_TakedownsAfterExhausting"`
	MissionsTakedownsAfterTeleporting              int64          `json:"Missions_TakedownsAfterTeleporting"`
	MissionsTakedownsBefore15Min                   int64          `json:"Missions_TakedownsBefore15Min"`
	MissionsTakedownStructures                     int64          `json:"Missions_TakedownStructures"`
	MissionsTakedownsUnderTurret                   int64          `json:"Missions_TakedownsUnderTurret"`
	MissionsTakedownsWithHelpFromMonsters          int64          `json:"Missions_TakedownsWithHelpFromMonsters"`
	MissionsTakedownWards                          int64          `json:"Missions_TakedownWards"`
	MissionsTimeSpentActivelyPlaying               int64          `json:"Missions_TimeSpentActivelyPlaying"`
	MissionsTotalGold                              int64          `json:"Missions_TotalGold"`
	MissionsTrueDamageToStructures                 int64          `json:"Missions_TrueDamageToStructures"`
	MissionsTurretPlatesDestroyed                  int64          `json:"Missions_TurretPlatesDestroyed"`
	MissionsTwoChampsKilledWithSameAbility         int64          `json:"Missions_TwoChampsKilledWithSameAbility"`
	MissionsVoidMitesSummoned                      int64          `json:"Missions_VoidMitesSummoned"`
	MutedAll                                       int64          `json:"MUTED_ALL"`
	Name                                           string         `json:"NAME"`
	NeedVisionPings                                int64          `json:"NEED_VISION_PINGS"`
	NeutralMinionsKilled                           int64          `json:"NEUTRAL_MINIONS_KILLED"`
	NeutralMinionsKilledEnemyJungle                int64          `json:"NEUTRAL_MINIONS_KILLED_ENEMY_JUNGLE"`
	NeutralMinionsKilledYourJungle                 int64          `json:"NEUTRAL_MINIONS_KILLED_YOUR_JUNGLE"`
	NodeCapture                                    int64          `json:"NODE_CAPTURE"`
	NodeCaptureAssist                              int64          `json:"NODE_CAPTURE_ASSIST"`
	NodeNeutralize                                 int64          `json:"NODE_NEUTRALIZE"`
	NodeNeutralizeAssist                           int64          `json:"NODE_NEUTRALIZE_ASSIST"`
	NumDeaths                                      int64          `json:"NUM_DEATHS"`
	ObjectivesStolen                               int64          `json:"OBJECTIVES_STOLEN"`
	ObjectivesStolenAssists                        int64          `json:"OBJECTIVES_STOLEN_ASSISTS"`
	OnMyWayPings                                   int64          `json:"ON_MY_WAY_PINGS"`
	PentaKills                                     int64          `json:"PENTA_KILLS"`
	Perk0                                          int64          `json:"PERK0"`
	Perk0Var1                                      int64          `json:"PERK0_VAR1"`
	Perk0Var2                                      int64          `json:"PERK0_VAR2"`
	Perk0Var3                                      int64          `json:"PERK0_VAR3"`
	Perk1                                          int64          `json:"PERK1"`
	Perk1Var1                                      int64          `json:"PERK1_VAR1"`
	Perk1Var2                                      int64          `json:"PERK1_VAR2"`
	Perk1Var3                                      int64          `json:"PERK1_VAR3"`
	Perk2                                          int64          `json:"PERK2"`
	Perk2Var1                                      int64          `json:"PERK2_VAR1"`
	Perk2Var2                                      int64          `json:"PERK2_VAR2"`
	Perk2Var3                                      int64          `json:"PERK2_VAR3"`
	Perk3                                          int64          `json:"PERK3"`
	Perk3Var1                                      int64          `json:"PERK3_VAR1"`
	Perk3Var2                                      int64          `json:"PERK3_VAR2"`
	Perk3Var3                                      int64          `json:"PERK3_VAR3"`
	Perk4                                          int64          `json:"PERK4"`
	Perk4Var1                                      int64          `json:"PERK4_VAR1"`
	Perk4Var2                                      int64          `json:"PERK4_VAR2"`
	Perk4Var3                                      int64          `json:"PERK4_VAR3"`
	Perk5                                          int64          `json:"PERK5"`
	Perk5Var1                                      int64          `json:"PERK5_VAR1"`
	Perk5Var2                                      int64          `json:"PERK5_VAR2"`
	Perk5Var3                                      int64          `json:"PERK5_VAR3"`
	PerkPrimaryStyle                               int64          `json:"PERK_PRIMARY_STYLE"`
	PerkSubStyle                                   int64          `json:"PERK_SUB_STYLE"`
	PhysicalDamageDealtPlayer                      int64          `json:"PHYSICAL_DAMAGE_DEALT_PLAYER"`
	PhysicalDamageDealtToChampions                 int64          `json:"PHYSICAL_DAMAGE_DEALT_TO_CHAMPIONS"`
	PhysicalDamageTaken                            int64          `json:"PHYSICAL_DAMAGE_TAKEN"`
	Ping                                           int64          `json:"PING"`
	PlayerAugment1                                 int64          `json:"PLAYER_AUGMENT_1"`
	PlayerAugment2                                 int64          `json:"PLAYER_AUGMENT_2"`
	PlayerAugment3                                 int64          `json:"PLAYER_AUGMENT_3"`
	PlayerAugment4                                 int64          `json:"PLAYER_AUGMENT_4"`
	PlayerAugment5                                 int64          `json:"PLAYER_AUGMENT_5"`
	PlayerAugment6                                 int64          `json:"PLAYER_AUGMENT_6"`
	PlayerPosition                                 int64          `json:"PLAYER_POSITION"`
	PlayerRole                                     int64          `json:"PLAYER_ROLE"`
	PlayerScore0                                   int64          `json:"PLAYER_SCORE_0"`
	PlayerScore1                                   int64          `json:"PLAYER_SCORE_1"`
	PlayerScore10                                  int64          `json:"PLAYER_SCORE_10"`
	PlayerScore11                                  int64          `json:"PLAYER_SCORE_11"`
	PlayerScore2                                   int64          `json:"PLAYER_SCORE_2"`
	PlayerScore3                                   int64          `json:"PLAYER_SCORE_3"`
	PlayerScore4                                   int64          `json:"PLAYER_SCORE_4"`
	PlayerScore5                                   int64          `json:"PLAYER_SCORE_5"`
	PlayerScore6                                   int64          `json:"PLAYER_SCORE_6"`
	PlayerScore7                                   int64          `json:"PLAYER_SCORE_7"`
	PlayerScore8                                   int64          `json:"PLAYER_SCORE_8"`
	PlayerScore9                                   int64          `json:"PLAYER_SCORE_9"`
	PlayerSubteam                                  int64          `json:"PLAYER_SUBTEAM"`
	PlayerSubteamPlacement                         int64          `json:"PLAYER_SUBTEAM_PLACEMENT"`
	PlayersIMuted                                  int64          `json:"PLAYERS_I_MUTED"`
	PlayersThatMutedMe                             int64          `json:"PLAYERS_THAT_MUTED_ME"`
	PushPings                                      int64          `json:"PUSH_PINGS"`
	Puuid                                          string         `json:"PUUID"`
	QuadraKills                                    int64          `json:"QUADRA_KILLS"`
	RetreatPings                                   int64          `json:"RETREAT_PINGS"`
	RiftHeraldKills                                int64          `json:"RIFT_HERALD_KILLS"`
	RiotIDGameName                                 string         `json:"RIOT_ID_GAME_NAME"`
	RiotIDTagLine                                  *RiotIDTagLine `json:"RIOT_ID_TAG_LINE"`
	S3A1EventDoombotsTakenDownBefore5              int64          `json:"S3A1_Event_DoombotsTakenDownBefore5"`
	S3A1PlayAsDemaciansOrAgainstNoxians            int64          `json:"S3A1_PlayAsDemaciansOrAgainstNoxians"`
	S3A1Takedowns                                  int64          `json:"S3A1_Takedowns"`
	S3A2PrismaticAug                               int64          `json:"S3A2_PrismaticAug"`
	S3A2ZaahenUnlock                               int64          `json:"S3A2_ZaahenUnlock"`
	SeasonalMissionsTakedownAtakhan                int64          `json:"SeasonalMissions_TakedownAtakhan"`
	SightWardsBoughtInGame                         int64          `json:"SIGHT_WARDS_BOUGHT_IN_GAME"`
	Skin                                           string         `json:"SKIN"`
	Spell1Cast                                     int64          `json:"SPELL1_CAST"`
	Spell2Cast                                     int64          `json:"SPELL2_CAST"`
	Spell3Cast                                     int64          `json:"SPELL3_CAST"`
	Spell4Cast                                     int64          `json:"SPELL4_CAST"`
	StatPerk0                                      int64          `json:"STAT_PERK_0"`
	StatPerk1                                      int64          `json:"STAT_PERK_1"`
	StatPerk2                                      int64          `json:"STAT_PERK_2"`
	SummonSpell1Cast                               int64          `json:"SUMMON_SPELL1_CAST"`
	SummonSpell2Cast                               int64          `json:"SUMMON_SPELL2_CAST"`
	SummonerID                                     int64          `json:"SUMMONER_ID"`
	SummonerSpell1                                 int64          `json:"SUMMONER_SPELL_1"`
	SummonerSpell2                                 int64          `json:"SUMMONER_SPELL_2"`
	Team                                           int64          `json:"TEAM"`
	TeamEarlySurrendered                           int64          `json:"TEAM_EARLY_SURRENDERED"`
	TeamObjective                                  int64          `json:"TEAM_OBJECTIVE"`
	TeamPosition                                   string         `json:"TEAM_POSITION"`
	TimeCcingOthers                                int64          `json:"TIME_CCING_OTHERS"`
	TimeOfFromLastDisconnect                       int64          `json:"TIME_OF_FROM_LAST_DISCONNECT"`
	TimePlayed                                     int64          `json:"TIME_PLAYED"`
	TimeSpentDisconnected                          int64          `json:"TIME_SPENT_DISCONNECTED"`
	TotalDamageDealt                               int64          `json:"TOTAL_DAMAGE_DEALT"`
	TotalDamageDealtToBuildings                    int64          `json:"TOTAL_DAMAGE_DEALT_TO_BUILDINGS"`
	TotalDamageDealtToChampions                    int64          `json:"TOTAL_DAMAGE_DEALT_TO_CHAMPIONS"`
	TotalDamageDealtToEpicMonsters                 int64          `json:"TOTAL_DAMAGE_DEALT_TO_EPIC_MONSTERS"`
	TotalDamageDealtToObjectives                   int64          `json:"TOTAL_DAMAGE_DEALT_TO_OBJECTIVES"`
	TotalDamageDealtToTurrets                      int64          `json:"TOTAL_DAMAGE_DEALT_TO_TURRETS"`
	TotalDamageSelfMitigated                       int64          `json:"TOTAL_DAMAGE_SELF_MITIGATED"`
	TotalDamageShieldedOnTeammates                 int64          `json:"TOTAL_DAMAGE_SHIELDED_ON_TEAMMATES"`
	TotalDamageTaken                               int64          `json:"TOTAL_DAMAGE_TAKEN"`
	TotalHeal                                      int64          `json:"TOTAL_HEAL"`
	TotalHealOnTeammates                           int64          `json:"TOTAL_HEAL_ON_TEAMMATES"`
	TotalTimeCrowdControlDealt                     int64          `json:"TOTAL_TIME_CROWD_CONTROL_DEALT"`
	TotalTimeCrowdControlDealtToChampions          int64          `json:"TOTAL_TIME_CROWD_CONTROL_DEALT_TO_CHAMPIONS"`
	TotalTimeSpentDead                             int64          `json:"TOTAL_TIME_SPENT_DEAD"`
	TotalUnitsHealed                               int64          `json:"TOTAL_UNITS_HEALED"`
	TripleKills                                    int64          `json:"TRIPLE_KILLS"`
	TrueDamageDealtPlayer                          int64          `json:"TRUE_DAMAGE_DEALT_PLAYER"`
	TrueDamageDealtToChampions                     int64          `json:"TRUE_DAMAGE_DEALT_TO_CHAMPIONS"`
	TrueDamageTaken                                int64          `json:"TRUE_DAMAGE_TAKEN"`
	TurretTakedowns                                int64          `json:"TURRET_TAKEDOWNS"`
	TurretsKilled                                  int64          `json:"TURRETS_KILLED"`
	UnrealKills                                    int64          `json:"UNREAL_KILLS"`
	VictoryPointTotal                              int64          `json:"VICTORY_POINT_TOTAL"`
	VisionClearedPings                             int64          `json:"VISION_CLEARED_PINGS"`
	VisionScore                                    int64          `json:"VISION_SCORE"`
	VisionWardsBoughtInGame                        int64          `json:"VISION_WARDS_BOUGHT_IN_GAME"`
	WardKilled                                     int64          `json:"WARD_KILLED"`
	WardPlaced                                     int64          `json:"WARD_PLACED"`
	WardPlacedDetector                             int64          `json:"WARD_PLACED_DETECTOR"`
	WasAfk                                         int64          `json:"WAS_AFK"`
	WasAfkAfterFailedSurrender                     int64          `json:"WAS_AFK_AFTER_FAILED_SURRENDER"`
	WasEarlySurrenderAccomplice                    int64          `json:"WAS_EARLY_SURRENDER_ACCOMPLICE"`
	WasLeaver                                      int64          `json:"WAS_LEAVER"`
	WasSurrenderDueToAfk                           int64          `json:"WAS_SURRENDER_DUE_TO_AFK"`
	WeeklyMissionS2DamagingAbilities               int64          `json:"WeeklyMission_S2_DamagingAbilities"`
	WeeklyMissionS2FeatsOfStrength                 int64          `json:"WeeklyMission_S2_FeatsOfStrength"`
	WeeklyMissionS2SpiritPetals                    int64          `json:"WeeklyMission_S2_SpiritPetals"`
	Win                                            string         `json:"WIN"`
}

type RiotIDTagLine struct {
	Integer *int64
	String  *string
}

func (x *RiotIDTagLine) UnmarshalJSON(data []byte) error {
	object, err := unmarshalUnion(data, &x.Integer, nil, nil, &x.String, false, nil, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *RiotIDTagLine) MarshalJSON() ([]byte, error) {
	return marshalUnion(x.Integer, nil, nil, x.String, false, nil, false, nil, false, nil, false, nil, false)
}

func unmarshalUnion(data []byte, pi **int64, pf **float64, pb **bool, ps **string, haveArray bool, pa interface{}, haveObject bool, pc interface{}, haveMap bool, pm interface{}, haveEnum bool, pe interface{}, nullable bool) (bool, error) {
	if pi != nil {
			*pi = nil
	}
	if pf != nil {
			*pf = nil
	}
	if pb != nil {
			*pb = nil
	}
	if ps != nil {
			*ps = nil
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.UseNumber()
	tok, err := dec.Token()
	if err != nil {
			return false, err
	}

	switch v := tok.(type) {
	case json.Number:
			if pi != nil {
					i, err := v.Int64()
					if err == nil {
							*pi = &i
							return false, nil
					}
			}
			if pf != nil {
					f, err := v.Float64()
					if err == nil {
							*pf = &f
							return false, nil
					}
					return false, errors.New("Unparsable number")
			}
			return false, errors.New("Union does not contain number")
	case float64:
			return false, errors.New("Decoder should not return float64")
	case bool:
			if pb != nil {
					*pb = &v
					return false, nil
			}
			return false, errors.New("Union does not contain bool")
	case string:
			if haveEnum {
					return false, json.Unmarshal(data, pe)
			}
			if ps != nil {
					*ps = &v
					return false, nil
			}
			return false, errors.New("Union does not contain string")
	case nil:
			if nullable {
					return false, nil
			}
			return false, errors.New("Union does not contain null")
	case json.Delim:
			if v == '{' {
					if haveObject {
							return true, json.Unmarshal(data, pc)
					}
					if haveMap {
							return false, json.Unmarshal(data, pm)
					}
					return false, errors.New("Union does not contain object")
			}
			if v == '[' {
					if haveArray {
							return false, json.Unmarshal(data, pa)
					}
					return false, errors.New("Union does not contain array")
			}
			return false, errors.New("Cannot handle delimiter")
	}
	return false, errors.New("Cannot unmarshal union")
}

func marshalUnion(pi *int64, pf *float64, pb *bool, ps *string, haveArray bool, pa interface{}, haveObject bool, pc interface{}, haveMap bool, pm interface{}, haveEnum bool, pe interface{}, nullable bool) ([]byte, error) {
	if pi != nil {
			return json.Marshal(*pi)
	}
	if pf != nil {
			return json.Marshal(*pf)
	}
	if pb != nil {
			return json.Marshal(*pb)
	}
	if ps != nil {
			return json.Marshal(*ps)
	}
	if haveArray {
			return json.Marshal(pa)
	}
	if haveObject {
			return json.Marshal(pc)
	}
	if haveMap {
			return json.Marshal(pm)
	}
	if haveEnum {
			return json.Marshal(pe)
	}
	if nullable {
			return json.Marshal(nil)
	}
	return nil, errors.New("Union must not be null")
}
