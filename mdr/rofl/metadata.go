package rofl

/// This file is the Go representation of the JSON schema.
//
// REGENERATION WORKFLOW:
// 1. Run: python3 scripts/extract_metadata_schema.py <rofl_file> -o analysis_tmp/replay-schema.json --cast-numbers
// 2. Go to https://app.quicktype.io/ and generate Go code from the schema
// 3. Replace the StatsJSON struct fields with the generated ones
// 4. IMPORTANT: Keep the following custom types and methods:
//    - FlexInt64 type and its UnmarshalJSON/MarshalJSON methods
//    - rawMetadata struct
//    - Metadata.UnmarshalJSON custom method
//    - RiotIDTagLine and its union marshaling methods
//    - Use FlexInt64 instead of int64 for numeric fields in StatsJSON
//
// NOTE: The statsJson field in ROFL files is a JSON-encoded string, not a direct array.
// The custom UnmarshalJSON on Metadata handles this two-step parsing.
// Additionally, numeric fields may appear as strings in the JSON (e.g., "0" instead of 0),
// which is why we use FlexInt64 instead of int64.
//
// Code generated from JSON Schema using quicktype. Parts manually modified.

import (
	"bytes"
	"encoding/json"
	"errors"
	"strconv"
)

// FlexInt64 is a type that can unmarshal from either a JSON number or a JSON string
type FlexInt64 int64

func (f *FlexInt64) UnmarshalJSON(data []byte) error {
	// Try to unmarshal as int64 first
	var i int64
	if err := json.Unmarshal(data, &i); err == nil {
		*f = FlexInt64(i)
		return nil
	}

	// Try to unmarshal as string and convert
	var s string
	if err := json.Unmarshal(data, &s); err == nil {
		if s == "" {
			*f = 0
			return nil
		}
		i, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return err
		}
		*f = FlexInt64(i)
		return nil
	}

	return errors.New("FlexInt64: cannot unmarshal value")
}

func (f FlexInt64) MarshalJSON() ([]byte, error) {
	return json.Marshal(int64(f))
}

func UnmarshalMetadata(data []byte) (Metadata, error) {
	var r Metadata
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Metadata) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// rawMetadata is used for initial JSON parsing where statsJson is a string
type rawMetadata struct {
	GameLength      FlexInt64 `json:"gameLength"`
	LastGameChunkID FlexInt64 `json:"lastGameChunkId"`
	LastKeyFrameID  FlexInt64 `json:"lastKeyFrameId"`
	StatsJSON       string    `json:"statsJson"`
}

type Metadata struct {
	GameLength      FlexInt64   `json:"gameLength"`
	LastGameChunkID FlexInt64   `json:"lastGameChunkId"`
	LastKeyFrameID  FlexInt64   `json:"lastKeyFrameId"`
	StatsJSON       []StatsJSON `json:"statsJson"`
}

// UnmarshalJSON implements custom unmarshaling for Metadata.
// The statsJson field in ROFL files is a JSON-encoded string, not a direct array,
// so we need to unmarshal it in two steps.
func (m *Metadata) UnmarshalJSON(data []byte) error {
	var raw rawMetadata
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	m.GameLength = raw.GameLength
	m.LastGameChunkID = raw.LastGameChunkID
	m.LastKeyFrameID = raw.LastKeyFrameID

	// statsJson is a JSON string that needs to be parsed again
	if raw.StatsJSON != "" {
		if err := json.Unmarshal([]byte(raw.StatsJSON), &m.StatsJSON); err != nil {
			return err
		}
	}

	return nil
}

type StatsJSON struct {
	The2026_S1A1SkinsAshe                          FlexInt64      `json:"2026_S1A1_Skins_Ashe"`
	The2026_S1A1SkinsBriar                         FlexInt64      `json:"2026_S1A1_Skins_Briar"`
	The2026_S1A1SkinsCaitlyn                       FlexInt64      `json:"2026_S1A1_Skins_Caitlyn"`
	The2026_S1A1SkinsCamille                       FlexInt64      `json:"2026_S1A1_Skins_Camille"`
	The2026_S1A1SkinsGalio                         FlexInt64      `json:"2026_S1A1_Skins_Galio"`
	The2026_S1A1SkinsJayce                         FlexInt64      `json:"2026_S1A1_Skins_Jayce"`
	The2026_S1A1SkinsKatarina                      FlexInt64      `json:"2026_S1A1_Skins_Katarina"`
	The2026_S1A1SkinsLillia                        FlexInt64      `json:"2026_S1A1_Skins_Lillia"`
	The2026_S1A1SkinsNautilus                      FlexInt64      `json:"2026_S1A1_Skins_Nautilus"`
	The2026_S1A1SkinsOrnn                          FlexInt64      `json:"2026_S1A1_Skins_Ornn"`
	The2026_S1A1SkinsPoppy                         FlexInt64      `json:"2026_S1A1_Skins_Poppy"`
	The2026_S1A1SkinsSamira                        FlexInt64      `json:"2026_S1A1_Skins_Samira"`
	The2026_S1A1SkinsSeraphine                     FlexInt64      `json:"2026_S1A1_Skins_Seraphine"`
	The2026_S1A1SkinsYasuo                         FlexInt64      `json:"2026_S1A1_Skins_Yasuo"`
	The2026_S1A1SkinsYuumi                         FlexInt64      `json:"2026_S1A1_Skins_Yuumi"`
	The2026_S1A1SkinsZiggs                         FlexInt64      `json:"2026_S1A1_Skins_Ziggs"`
	The2026_S1A1SRFaerieWards                      FlexInt64      `json:"2026_S1A1_SR_FaerieWards"`
	The2026_S1A1SRGrowthSmashed                    FlexInt64      `json:"2026_S1A1_SR_GrowthSmashed"`
	The2026_S1A1SRRoleQuestComplete                FlexInt64      `json:"2026_S1A1_SR_RoleQuestComplete"`
	ActMissionS1A2ArenaRoundsWon                   FlexInt64      `json:"ActMission_S1_A2_ArenaRoundsWon"`
	ActMissionS1A2BloodyPetalsCollected            FlexInt64      `json:"ActMission_S1_A2_BloodyPetalsCollected"`
	ActMissionS1A2FeatsOfStrength                  FlexInt64      `json:"ActMission_S1_A2_FeatsOfStrength"`
	AllInPings                                     FlexInt64      `json:"ALL_IN_PINGS"`
	AssistMePings                                  FlexInt64      `json:"ASSIST_ME_PINGS"`
	Assists                                        FlexInt64      `json:"ASSISTS"`
	AtakhanKills                                   FlexInt64      `json:"ATAKHAN_KILLS"`
	BaronKills                                     FlexInt64      `json:"BARON_KILLS"`
	BarracksKilled                                 FlexInt64      `json:"BARRACKS_KILLED"`
	BarracksTakedowns                              FlexInt64      `json:"BARRACKS_TAKEDOWNS"`
	BasicPings                                     FlexInt64      `json:"BASIC_PINGS"`
	ChampionMissionStat0                           FlexInt64      `json:"CHAMPION_MISSION_STAT_0"`
	ChampionMissionStat1                           FlexInt64      `json:"CHAMPION_MISSION_STAT_1"`
	ChampionMissionStat2                           FlexInt64      `json:"CHAMPION_MISSION_STAT_2"`
	ChampionMissionStat3                           FlexInt64      `json:"CHAMPION_MISSION_STAT_3"`
	ChampionTransform                              FlexInt64      `json:"CHAMPION_TRANSFORM"`
	ChampionsKilled                                FlexInt64      `json:"CHAMPIONS_KILLED"`
	CommandPings                                   FlexInt64      `json:"COMMAND_PINGS"`
	ConsumablesPurchased                           FlexInt64      `json:"CONSUMABLES_PURCHASED"`
	DangerPings                                    FlexInt64      `json:"DANGER_PINGS"`
	DemonsHandMissionPointsA                       FlexInt64      `json:"DemonsHand_MissionPointsA"`
	DemonsHandMissionPointsB                       FlexInt64      `json:"DemonsHand_MissionPointsB"`
	DemonsHandMissionPointsC                       FlexInt64      `json:"DemonsHand_MissionPointsC"`
	DemonsHandMissionPointsD                       FlexInt64      `json:"DemonsHand_MissionPointsD"`
	DemonsHandMissionPointsE                       FlexInt64      `json:"DemonsHand_MissionPointsE"`
	DemonsHandMissionPointsF                       FlexInt64      `json:"DemonsHand_MissionPointsF"`
	DoubleKills                                    FlexInt64      `json:"DOUBLE_KILLS"`
	DragonKills                                    FlexInt64      `json:"DRAGON_KILLS"`
	EnemyMissingPings                              FlexInt64      `json:"ENEMY_MISSING_PINGS"`
	EnemyVisionPings                               FlexInt64      `json:"ENEMY_VISION_PINGS"`
	Event2025LRStructuresEpicMonsters              FlexInt64      `json:"Event_2025LR_StructuresEpicMonsters"`
	EventARAMDocks                                 FlexInt64      `json:"Event_ARAM_Docks"`
	EventARAMHexgates                              FlexInt64      `json:"Event_ARAM_Hexgates"`
	EventBrawlJungle                               FlexInt64      `json:"Event_Brawl_Jungle"`
	EventBrawlMinions                              FlexInt64      `json:"Event_Brawl_Minions"`
	EventS1A1AprilFoolsDragon                      FlexInt64      `json:"Event_S1_A1_AprilFools_Dragon"`
	EventS1A1AprilFoolsSnowball                    FlexInt64      `json:"Event_S1_A1_AprilFools_Snowball"`
	EventS1A2AprilFoolsDragon                      FlexInt64      `json:"Event_S1_A2_AprilFools_Dragon"`
	EventS1A2AprilFoolsGarenPlay                   FlexInt64      `json:"Event_S1_A2_AprilFools_Garen_Play"`
	EventS1A2AprilFoolsGarenTakedown               FlexInt64      `json:"Event_S1_A2_AprilFools_Garen_Takedown"`
	EventS1A2AprilFoolsSnowball                    FlexInt64      `json:"Event_S1_A2_AprilFools_Snowball"`
	EventS1A2ArenaBraveryChampions                 FlexInt64      `json:"Event_S1_A2_Arena_BraveryChampions"`
	EventS1A2ArenaNoxianChampions                  FlexInt64      `json:"Event_S1_A2_Arena_NoxianChampions"`
	EventS1A2ArenaReviveAllies                     FlexInt64      `json:"Event_S1_A2_Arena_ReviveAllies"`
	EventS1A2EsportsTakedownEpicMonstersSingleGame FlexInt64      `json:"Event_S1_A2_Esports_TakedownEpicMonstersSingleGame"`
	EventS1A2Mordekaiser                           FlexInt64      `json:"Event_S1_A2_Mordekaiser"`
	EventS2A2Exalted                               FlexInt64      `json:"Event_S2A2_Exalted"`
	EventS2A2MV                                    FlexInt64      `json:"Event_S2A2_MV"`
	EventS2A2PetalPoints                           FlexInt64      `json:"Event_S2A2_PetalPoints"`
	EventS2A2ChampDamageAbilities                  FlexInt64      `json:"Event_S2A2Champ_DamageAbilities"`
	EventS2A2ChampDamageAutos                      FlexInt64      `json:"Event_S2A2Champ_DamageAutos"`
	Exp                                            FlexInt64      `json:"EXP"`
	FriendlyDampenLost                             FlexInt64      `json:"FRIENDLY_DAMPEN_LOST"`
	FriendlyHqLost                                 FlexInt64      `json:"FRIENDLY_HQ_LOST"`
	FriendlyTurretLost                             FlexInt64      `json:"FRIENDLY_TURRET_LOST"`
	GameEndedInEarlySurrender                      FlexInt64      `json:"GAME_ENDED_IN_EARLY_SURRENDER"`
	GameEndedInSurrender                           FlexInt64      `json:"GAME_ENDED_IN_SURRENDER"`
	GetBackPings                                   FlexInt64      `json:"GET_BACK_PINGS"`
	GoldEarned                                     FlexInt64      `json:"GOLD_EARNED"`
	GoldSpent                                      FlexInt64      `json:"GOLD_SPENT"`
	HoLChampionsDamagedWhileHidden                 FlexInt64      `json:"HoL_ChampionsDamagedWhileHidden"`
	HoLControlWardsKilled                          FlexInt64      `json:"HoL_ControlWardsKilled"`
	HoLEliteAsheCrystalArrowTakedowns              FlexInt64      `json:"HoL_Elite_AsheCrystalArrowTakedowns"`
	HoLEliteAsheHawkshotChampsRevealed             FlexInt64      `json:"HoL_Elite_AsheHawkshotChampsRevealed"`
	HoLEliteEzrealEssenceFluxDetonated             FlexInt64      `json:"HoL_Elite_EzrealEssenceFluxDetonated"`
	HoLEliteEzrealTrueshotBarrageMultiHit          FlexInt64      `json:"HoL_Elite_EzrealTrueshotBarrageMultiHit"`
	HoLEliteKaiSaAbilitiesUpgraded                 FlexInt64      `json:"HoL_Elite_KaiSaAbilitiesUpgraded"`
	HoLEliteKaiSaKillerInstinctKills               FlexInt64      `json:"HoL_Elite_KaiSaKillerInstinctKills"`
	HoLEliteLucianCullingHits                      FlexInt64      `json:"HoL_Elite_LucianCullingHits"`
	HoLEliteLucianPiercingLightMultiHit            FlexInt64      `json:"HoL_Elite_LucianPiercingLightMultiHit"`
	HoLEliteVayneCondemnStun                       FlexInt64      `json:"HoL_Elite_VayneCondemnStun"`
	HoLEliteVayneTumbleDodge                       FlexInt64      `json:"HoL_Elite_VayneTumbleDodge"`
	HoLEnemyTakedownUnderTower                     FlexInt64      `json:"HoL_EnemyTakedownUnderTower"`
	HoLFightsSurvivedWhileLowHealth                FlexInt64      `json:"HoL_FightsSurvivedWhileLowHealth"`
	HoLHiddenEnemiesDamaged                        FlexInt64      `json:"HoL_HiddenEnemiesDamaged"`
	HoLJungleCampsStolen                           FlexInt64      `json:"HoL_JungleCampsStolen"`
	HoLKillsWhileLowHealth                         FlexInt64      `json:"HoL_KillsWhileLowHealth"`
	HoLOutnumberedTakedowns                        FlexInt64      `json:"HoL_OutnumberedTakedowns"`
	HoLShutdownGoldCollected                       FlexInt64      `json:"HoL_ShutdownGoldCollected"`
	HoLSoloKills                                   FlexInt64      `json:"HoL_SoloKills"`
	HoLTurretsTakenWithinMinutes                   FlexInt64      `json:"HoL_TurretsTakenWithinMinutes"`
	HoldPings                                      FlexInt64      `json:"HOLD_PINGS"`
	HordeKills                                     FlexInt64      `json:"HORDE_KILLS"`
	HqKilled                                       FlexInt64      `json:"HQ_KILLED"`
	HqTakedowns                                    FlexInt64      `json:"HQ_TAKEDOWNS"`
	ID                                             FlexInt64      `json:"ID"`
	IndividualPosition                             string         `json:"INDIVIDUAL_POSITION"`
	Item0                                          FlexInt64      `json:"ITEM0"`
	Item1                                          FlexInt64      `json:"ITEM1"`
	Item2                                          FlexInt64      `json:"ITEM2"`
	Item3                                          FlexInt64      `json:"ITEM3"`
	Item4                                          FlexInt64      `json:"ITEM4"`
	Item5                                          FlexInt64      `json:"ITEM5"`
	Item6                                          FlexInt64      `json:"ITEM6"`
	ItemsPurchased                                 FlexInt64      `json:"ITEMS_PURCHASED"`
	KeystoneID                                     FlexInt64      `json:"KEYSTONE_ID"`
	KillingSprees                                  FlexInt64      `json:"KILLING_SPREES"`
	LargestAbilityDamage                           FlexInt64      `json:"LARGEST_ABILITY_DAMAGE"`
	LargestAttackDamage                            FlexInt64      `json:"LARGEST_ATTACK_DAMAGE"`
	LargestCriticalStrike                          FlexInt64      `json:"LARGEST_CRITICAL_STRIKE"`
	LargestKillingSpree                            FlexInt64      `json:"LARGEST_KILLING_SPREE"`
	LargestMultiKill                               FlexInt64      `json:"LARGEST_MULTI_KILL"`
	LastTakedownTime                               FlexInt64      `json:"LAST_TAKEDOWN_TIME"`
	Level                                          FlexInt64      `json:"LEVEL"`
	LongestTimeSpentLiving                         FlexInt64      `json:"LONGEST_TIME_SPENT_LIVING"`
	MagicDamageDealtPlayer                         FlexInt64      `json:"MAGIC_DAMAGE_DEALT_PLAYER"`
	MagicDamageDealtToChampions                    FlexInt64      `json:"MAGIC_DAMAGE_DEALT_TO_CHAMPIONS"`
	MagicDamageTaken                               FlexInt64      `json:"MAGIC_DAMAGE_TAKEN"`
	MinionsKilled                                  FlexInt64      `json:"MINIONS_KILLED"`
	MissionsBXPEarnedPerGame                       FlexInt64      `json:"Missions_BXP_EarnedPerGame"`
	MissionsCannonMinionsKilled                    FlexInt64      `json:"Missions_CannonMinionsKilled"`
	MissionsChampionsHitWithAbilitiesEarlyGame     FlexInt64      `json:"Missions_ChampionsHitWithAbilitiesEarlyGame"`
	MissionsChampionsKilled                        FlexInt64      `json:"Missions_ChampionsKilled"`
	MissionsChampionTakedownsWhileGhosted          FlexInt64      `json:"Missions_ChampionTakedownsWhileGhosted"`
	MissionsChampionTakedownsWithIgnite            FlexInt64      `json:"Missions_ChampionTakedownsWithIgnite"`
	MissionsCreepScore                             FlexInt64      `json:"Missions_CreepScore"`
	MissionsCreepScoreBy10Minutes                  FlexInt64      `json:"Missions_CreepScoreBy10Minutes"`
	MissionsCrepeDamageDealtSpeedZone              FlexInt64      `json:"Missions_Crepe_DamageDealtSpeedZone"`
	MissionsCrepeSnowballLanded                    FlexInt64      `json:"Missions_Crepe_SnowballLanded"`
	MissionsCrepeTakedownsWithInhibBuff            FlexInt64      `json:"Missions_Crepe_TakedownsWithInhibBuff"`
	MissionsDamageToChampsWithItems                FlexInt64      `json:"Missions_DamageToChampsWithItems"`
	MissionsDamageToStructures                     FlexInt64      `json:"Missions_DamageToStructures"`
	MissionsDestroyPlants                          FlexInt64      `json:"Missions_DestroyPlants"`
	MissionsDominationRune                         FlexInt64      `json:"Missions_DominationRune"`
	MissionsGoldFromStructuresDestroyed            FlexInt64      `json:"Missions_GoldFromStructuresDestroyed"`
	MissionsGoldFromTurretPlatesTaken              FlexInt64      `json:"Missions_GoldFromTurretPlatesTaken"`
	MissionsGoldPerMinute                          FlexInt64      `json:"Missions_GoldPerMinute"`
	MissionsHealingFromLevelObjects                FlexInt64      `json:"Missions_HealingFromLevelObjects"`
	MissionsHexgatesUsed                           FlexInt64      `json:"Missions_HexgatesUsed"`
	MissionsImmobilizeChampions                    FlexInt64      `json:"Missions_ImmobilizeChampions"`
	MissionsInspirationRune                        FlexInt64      `json:"Missions_InspirationRune"`
	MissionsLegendaryItems                         FlexInt64      `json:"Missions_LegendaryItems"`
	MissionsMinionsKilled                          FlexInt64      `json:"Missions_MinionsKilled"`
	MissionsPeriodicDamage                         FlexInt64      `json:"Missions_PeriodicDamage"`
	MissionsPlaceUsefulControlWards                FlexInt64      `json:"Missions_PlaceUsefulControlWards"`
	MissionsPlaceUsefulWards                       FlexInt64      `json:"Missions_PlaceUsefulWards"`
	MissionsPorosFed                               FlexInt64      `json:"Missions_PorosFed"`
	MissionsPrecisionRune                          FlexInt64      `json:"Missions_PrecisionRune"`
	MissionsResolveRune                            FlexInt64      `json:"Missions_ResolveRune"`
	MissionsSnowballsHit                           FlexInt64      `json:"Missions_SnowballsHit"`
	MissionsSorceryRune                            FlexInt64      `json:"Missions_SorceryRune"`
	MissionsTakedownBaronsElderDragons             FlexInt64      `json:"Missions_TakedownBaronsElderDragons"`
	MissionsTakedownDragons                        FlexInt64      `json:"Missions_TakedownDragons"`
	MissionsTakedownEpicMonsters                   FlexInt64      `json:"Missions_TakedownEpicMonsters"`
	MissionsTakedownEpicMonstersSingleGame         FlexInt64      `json:"Missions_TakedownEpicMonstersSingleGame"`
	MissionsTakedownGold                           FlexInt64      `json:"Missions_TakedownGold"`
	MissionsTakedownsAfterExhausting               FlexInt64      `json:"Missions_TakedownsAfterExhausting"`
	MissionsTakedownsAfterTeleporting              FlexInt64      `json:"Missions_TakedownsAfterTeleporting"`
	MissionsTakedownsBefore15Min                   FlexInt64      `json:"Missions_TakedownsBefore15Min"`
	MissionsTakedownStructures                     FlexInt64      `json:"Missions_TakedownStructures"`
	MissionsTakedownsUnderTurret                   FlexInt64      `json:"Missions_TakedownsUnderTurret"`
	MissionsTakedownsWithHelpFromMonsters          FlexInt64      `json:"Missions_TakedownsWithHelpFromMonsters"`
	MissionsTakedownWards                          FlexInt64      `json:"Missions_TakedownWards"`
	MissionsTimeSpentActivelyPlaying               FlexInt64      `json:"Missions_TimeSpentActivelyPlaying"`
	MissionsTotalGold                              FlexInt64      `json:"Missions_TotalGold"`
	MissionsTrueDamageToStructures                 FlexInt64      `json:"Missions_TrueDamageToStructures"`
	MissionsTurretPlatesDestroyed                  FlexInt64      `json:"Missions_TurretPlatesDestroyed"`
	MissionsTwoChampsKilledWithSameAbility         FlexInt64      `json:"Missions_TwoChampsKilledWithSameAbility"`
	MissionsVoidMitesSummoned                      FlexInt64      `json:"Missions_VoidMitesSummoned"`
	MutedAll                                       FlexInt64      `json:"MUTED_ALL"`
	Name                                           string         `json:"NAME"`
	NeedVisionPings                                FlexInt64      `json:"NEED_VISION_PINGS"`
	NeutralMinionsKilled                           FlexInt64      `json:"NEUTRAL_MINIONS_KILLED"`
	NeutralMinionsKilledEnemyJungle                FlexInt64      `json:"NEUTRAL_MINIONS_KILLED_ENEMY_JUNGLE"`
	NeutralMinionsKilledYourJungle                 FlexInt64      `json:"NEUTRAL_MINIONS_KILLED_YOUR_JUNGLE"`
	NodeCapture                                    FlexInt64      `json:"NODE_CAPTURE"`
	NodeCaptureAssist                              FlexInt64      `json:"NODE_CAPTURE_ASSIST"`
	NodeNeutralize                                 FlexInt64      `json:"NODE_NEUTRALIZE"`
	NodeNeutralizeAssist                           FlexInt64      `json:"NODE_NEUTRALIZE_ASSIST"`
	NumDeaths                                      FlexInt64      `json:"NUM_DEATHS"`
	ObjectivesStolen                               FlexInt64      `json:"OBJECTIVES_STOLEN"`
	ObjectivesStolenAssists                        FlexInt64      `json:"OBJECTIVES_STOLEN_ASSISTS"`
	OnMyWayPings                                   FlexInt64      `json:"ON_MY_WAY_PINGS"`
	PentaKills                                     FlexInt64      `json:"PENTA_KILLS"`
	Perk0                                          FlexInt64      `json:"PERK0"`
	Perk0Var1                                      FlexInt64      `json:"PERK0_VAR1"`
	Perk0Var2                                      FlexInt64      `json:"PERK0_VAR2"`
	Perk0Var3                                      FlexInt64      `json:"PERK0_VAR3"`
	Perk1                                          FlexInt64      `json:"PERK1"`
	Perk1Var1                                      FlexInt64      `json:"PERK1_VAR1"`
	Perk1Var2                                      FlexInt64      `json:"PERK1_VAR2"`
	Perk1Var3                                      FlexInt64      `json:"PERK1_VAR3"`
	Perk2                                          FlexInt64      `json:"PERK2"`
	Perk2Var1                                      FlexInt64      `json:"PERK2_VAR1"`
	Perk2Var2                                      FlexInt64      `json:"PERK2_VAR2"`
	Perk2Var3                                      FlexInt64      `json:"PERK2_VAR3"`
	Perk3                                          FlexInt64      `json:"PERK3"`
	Perk3Var1                                      FlexInt64      `json:"PERK3_VAR1"`
	Perk3Var2                                      FlexInt64      `json:"PERK3_VAR2"`
	Perk3Var3                                      FlexInt64      `json:"PERK3_VAR3"`
	Perk4                                          FlexInt64      `json:"PERK4"`
	Perk4Var1                                      FlexInt64      `json:"PERK4_VAR1"`
	Perk4Var2                                      FlexInt64      `json:"PERK4_VAR2"`
	Perk4Var3                                      FlexInt64      `json:"PERK4_VAR3"`
	Perk5                                          FlexInt64      `json:"PERK5"`
	Perk5Var1                                      FlexInt64      `json:"PERK5_VAR1"`
	Perk5Var2                                      FlexInt64      `json:"PERK5_VAR2"`
	Perk5Var3                                      FlexInt64      `json:"PERK5_VAR3"`
	PerkPrimaryStyle                               FlexInt64      `json:"PERK_PRIMARY_STYLE"`
	PerkSubStyle                                   FlexInt64      `json:"PERK_SUB_STYLE"`
	PhysicalDamageDealtPlayer                      FlexInt64      `json:"PHYSICAL_DAMAGE_DEALT_PLAYER"`
	PhysicalDamageDealtToChampions                 FlexInt64      `json:"PHYSICAL_DAMAGE_DEALT_TO_CHAMPIONS"`
	PhysicalDamageTaken                            FlexInt64      `json:"PHYSICAL_DAMAGE_TAKEN"`
	Ping                                           FlexInt64      `json:"PING"`
	PlayerAugment1                                 FlexInt64      `json:"PLAYER_AUGMENT_1"`
	PlayerAugment2                                 FlexInt64      `json:"PLAYER_AUGMENT_2"`
	PlayerAugment3                                 FlexInt64      `json:"PLAYER_AUGMENT_3"`
	PlayerAugment4                                 FlexInt64      `json:"PLAYER_AUGMENT_4"`
	PlayerAugment5                                 FlexInt64      `json:"PLAYER_AUGMENT_5"`
	PlayerAugment6                                 FlexInt64      `json:"PLAYER_AUGMENT_6"`
	PlayerPosition                                 FlexInt64      `json:"PLAYER_POSITION"`
	PlayerRole                                     FlexInt64      `json:"PLAYER_ROLE"`
	PlayerScore0                                   FlexInt64      `json:"PLAYER_SCORE_0"`
	PlayerScore1                                   FlexInt64      `json:"PLAYER_SCORE_1"`
	PlayerScore10                                  FlexInt64      `json:"PLAYER_SCORE_10"`
	PlayerScore11                                  FlexInt64      `json:"PLAYER_SCORE_11"`
	PlayerScore2                                   FlexInt64      `json:"PLAYER_SCORE_2"`
	PlayerScore3                                   FlexInt64      `json:"PLAYER_SCORE_3"`
	PlayerScore4                                   FlexInt64      `json:"PLAYER_SCORE_4"`
	PlayerScore5                                   FlexInt64      `json:"PLAYER_SCORE_5"`
	PlayerScore6                                   FlexInt64      `json:"PLAYER_SCORE_6"`
	PlayerScore7                                   FlexInt64      `json:"PLAYER_SCORE_7"`
	PlayerScore8                                   FlexInt64      `json:"PLAYER_SCORE_8"`
	PlayerScore9                                   FlexInt64      `json:"PLAYER_SCORE_9"`
	PlayerSubteam                                  FlexInt64      `json:"PLAYER_SUBTEAM"`
	PlayerSubteamPlacement                         FlexInt64      `json:"PLAYER_SUBTEAM_PLACEMENT"`
	PlayersIMuted                                  FlexInt64      `json:"PLAYERS_I_MUTED"`
	PlayersThatMutedMe                             FlexInt64      `json:"PLAYERS_THAT_MUTED_ME"`
	PushPings                                      FlexInt64      `json:"PUSH_PINGS"`
	Puuid                                          string         `json:"PUUID"`
	QuadraKills                                    FlexInt64      `json:"QUADRA_KILLS"`
	RetreatPings                                   FlexInt64      `json:"RETREAT_PINGS"`
	RiftHeraldKills                                FlexInt64      `json:"RIFT_HERALD_KILLS"`
	RiotIDGameName                                 string         `json:"RIOT_ID_GAME_NAME"`
	RiotIDTagLine                                  *RiotIDTagLine `json:"RIOT_ID_TAG_LINE"`
	S3A1EventDoombotsTakenDownBefore5              FlexInt64      `json:"S3A1_Event_DoombotsTakenDownBefore5"`
	S3A1PlayAsDemaciansOrAgainstNoxians            FlexInt64      `json:"S3A1_PlayAsDemaciansOrAgainstNoxians"`
	S3A1Takedowns                                  FlexInt64      `json:"S3A1_Takedowns"`
	S3A2PrismaticAug                               FlexInt64      `json:"S3A2_PrismaticAug"`
	S3A2ZaahenUnlock                               FlexInt64      `json:"S3A2_ZaahenUnlock"`
	SeasonalMissionsTakedownAtakhan                FlexInt64      `json:"SeasonalMissions_TakedownAtakhan"`
	SightWardsBoughtInGame                         FlexInt64      `json:"SIGHT_WARDS_BOUGHT_IN_GAME"`
	Skin                                           string         `json:"SKIN"`
	Spell1Cast                                     FlexInt64      `json:"SPELL1_CAST"`
	Spell2Cast                                     FlexInt64      `json:"SPELL2_CAST"`
	Spell3Cast                                     FlexInt64      `json:"SPELL3_CAST"`
	Spell4Cast                                     FlexInt64      `json:"SPELL4_CAST"`
	StatPerk0                                      FlexInt64      `json:"STAT_PERK_0"`
	StatPerk1                                      FlexInt64      `json:"STAT_PERK_1"`
	StatPerk2                                      FlexInt64      `json:"STAT_PERK_2"`
	SummonSpell1Cast                               FlexInt64      `json:"SUMMON_SPELL1_CAST"`
	SummonSpell2Cast                               FlexInt64      `json:"SUMMON_SPELL2_CAST"`
	SummonerID                                     FlexInt64      `json:"SUMMONER_ID"`
	SummonerSpell1                                 FlexInt64      `json:"SUMMONER_SPELL_1"`
	SummonerSpell2                                 FlexInt64      `json:"SUMMONER_SPELL_2"`
	Team                                           FlexInt64      `json:"TEAM"`
	TeamEarlySurrendered                           FlexInt64      `json:"TEAM_EARLY_SURRENDERED"`
	TeamObjective                                  FlexInt64      `json:"TEAM_OBJECTIVE"`
	TeamPosition                                   string         `json:"TEAM_POSITION"`
	TimeCcingOthers                                FlexInt64      `json:"TIME_CCING_OTHERS"`
	TimeOfFromLastDisconnect                       FlexInt64      `json:"TIME_OF_FROM_LAST_DISCONNECT"`
	TimePlayed                                     FlexInt64      `json:"TIME_PLAYED"`
	TimeSpentDisconnected                          FlexInt64      `json:"TIME_SPENT_DISCONNECTED"`
	TotalDamageDealt                               FlexInt64      `json:"TOTAL_DAMAGE_DEALT"`
	TotalDamageDealtToBuildings                    FlexInt64      `json:"TOTAL_DAMAGE_DEALT_TO_BUILDINGS"`
	TotalDamageDealtToChampions                    FlexInt64      `json:"TOTAL_DAMAGE_DEALT_TO_CHAMPIONS"`
	TotalDamageDealtToEpicMonsters                 FlexInt64      `json:"TOTAL_DAMAGE_DEALT_TO_EPIC_MONSTERS"`
	TotalDamageDealtToObjectives                   FlexInt64      `json:"TOTAL_DAMAGE_DEALT_TO_OBJECTIVES"`
	TotalDamageDealtToTurrets                      FlexInt64      `json:"TOTAL_DAMAGE_DEALT_TO_TURRETS"`
	TotalDamageSelfMitigated                       FlexInt64      `json:"TOTAL_DAMAGE_SELF_MITIGATED"`
	TotalDamageShieldedOnTeammates                 FlexInt64      `json:"TOTAL_DAMAGE_SHIELDED_ON_TEAMMATES"`
	TotalDamageTaken                               FlexInt64      `json:"TOTAL_DAMAGE_TAKEN"`
	TotalHeal                                      FlexInt64      `json:"TOTAL_HEAL"`
	TotalHealOnTeammates                           FlexInt64      `json:"TOTAL_HEAL_ON_TEAMMATES"`
	TotalTimeCrowdControlDealt                     FlexInt64      `json:"TOTAL_TIME_CROWD_CONTROL_DEALT"`
	TotalTimeCrowdControlDealtToChampions          FlexInt64      `json:"TOTAL_TIME_CROWD_CONTROL_DEALT_TO_CHAMPIONS"`
	TotalTimeSpentDead                             FlexInt64      `json:"TOTAL_TIME_SPENT_DEAD"`
	TotalUnitsHealed                               FlexInt64      `json:"TOTAL_UNITS_HEALED"`
	TripleKills                                    FlexInt64      `json:"TRIPLE_KILLS"`
	TrueDamageDealtPlayer                          FlexInt64      `json:"TRUE_DAMAGE_DEALT_PLAYER"`
	TrueDamageDealtToChampions                     FlexInt64      `json:"TRUE_DAMAGE_DEALT_TO_CHAMPIONS"`
	TrueDamageTaken                                FlexInt64      `json:"TRUE_DAMAGE_TAKEN"`
	TurretTakedowns                                FlexInt64      `json:"TURRET_TAKEDOWNS"`
	TurretsKilled                                  FlexInt64      `json:"TURRETS_KILLED"`
	UnrealKills                                    FlexInt64      `json:"UNREAL_KILLS"`
	VictoryPointTotal                              FlexInt64      `json:"VICTORY_POINT_TOTAL"`
	VisionClearedPings                             FlexInt64      `json:"VISION_CLEARED_PINGS"`
	VisionScore                                    FlexInt64      `json:"VISION_SCORE"`
	VisionWardsBoughtInGame                        FlexInt64      `json:"VISION_WARDS_BOUGHT_IN_GAME"`
	WardKilled                                     FlexInt64      `json:"WARD_KILLED"`
	WardPlaced                                     FlexInt64      `json:"WARD_PLACED"`
	WardPlacedDetector                             FlexInt64      `json:"WARD_PLACED_DETECTOR"`
	WasAfk                                         FlexInt64      `json:"WAS_AFK"`
	WasAfkAfterFailedSurrender                     FlexInt64      `json:"WAS_AFK_AFTER_FAILED_SURRENDER"`
	WasEarlySurrenderAccomplice                    FlexInt64      `json:"WAS_EARLY_SURRENDER_ACCOMPLICE"`
	WasLeaver                                      FlexInt64      `json:"WAS_LEAVER"`
	WasSurrenderDueToAfk                           FlexInt64      `json:"WAS_SURRENDER_DUE_TO_AFK"`
	WeeklyMissionS2DamagingAbilities               FlexInt64      `json:"WeeklyMission_S2_DamagingAbilities"`
	WeeklyMissionS2FeatsOfStrength                 FlexInt64      `json:"WeeklyMission_S2_FeatsOfStrength"`
	WeeklyMissionS2SpiritPetals                    FlexInt64      `json:"WeeklyMission_S2_SpiritPetals"`
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
