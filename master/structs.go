// Generated code. DO NOT EDIT!
package master

import "time"

type AdvDatas struct {
  Id int `yaml:"Id"`
  AdvSeriesId int `yaml:"AdvSeriesId"`
  Name string `yaml:"Name"`
  Description string `yaml:"Description"`
  ScriptId int `yaml:"ScriptId"`
  OpenSeasonFanLevel int `yaml:"OpenSeasonFanLevel"`
  RewardType string `yaml:"RewardType"`
  WatchRewardId string `yaml:"WatchRewardId"`
  WatchRewardNum string `yaml:"WatchRewardNum"`
  RewardTextId string `yaml:"RewardTextId"`
  OrderId int `yaml:"OrderId"`
  SubTitleName string `yaml:"SubTitleName"`
  StartTime time.Time `yaml:"StartTime"`
  EndTime time.Time `yaml:"EndTime"`
  IsPeriod int `yaml:"IsPeriod"`
}

type AdvSeries struct {
  Id int `yaml:"Id"`
  SeasonsId int `yaml:"SeasonsId"`
  Name string `yaml:"Name"`
  Description string `yaml:"Description"`
  StartTime time.Time `yaml:"StartTime"`
  EndTime time.Time `yaml:"EndTime"`
  FiscalYearDisplay int `yaml:"FiscalYearDisplay"`
}

type AdvStoryDigestMovies struct {
  Id int `yaml:"Id"`
  Title string `yaml:"Title"`
  OrderId int `yaml:"OrderId"`
  DigestFiscalYearDisplay int `yaml:"DigestFiscalYearDisplay"`
}

type BeginnerMissionBannerRewards struct {
  Id int `yaml:"Id"`
  MissionType int `yaml:"MissionType"`
  ItemType int `yaml:"ItemType"`
  ItemId int `yaml:"ItemId"`
  Num int `yaml:"Num"`
}

type BeginnerMissionsHintImages struct {
  Id int `yaml:"Id"`
  BeginnerMissionsHintId int `yaml:"BeginnerMissionsHintId"`
  ImageId int `yaml:"ImageId"`
}

type BeginnerMissionsHint struct {
  Id int `yaml:"Id"`
  Name string `yaml:"Name"`
  Description string `yaml:"Description"`
}

type BirthdayRareBonuses struct {
  Id int `yaml:"Id"`
  CardSeriesId int `yaml:"CardSeriesId"`
  SkillName string `yaml:"SkillName"`
  LimitBreakTimes int `yaml:"LimitBreakTimes"`
  MentalBonus int `yaml:"MentalBonus"`
  VoltageBonus int `yaml:"VoltageBonus"`
  HeartBonus int `yaml:"HeartBonus"`
  LoveBonus int `yaml:"LoveBonus"`
}

type CampaignAddRewardSeries struct {
  Id int `yaml:"Id"`
  AddRewardOddsId_1 int64 `yaml:"AddRewardOddsId_1"`
  AddRewardOddsId_2 int64 `yaml:"AddRewardOddsId_2"`
  AddRewardOddsId_3 int64 `yaml:"AddRewardOddsId_3"`
}

type CampaignAddRewards struct {
  Id int `yaml:"Id"`
  AddRewardOddsId int64 `yaml:"AddRewardOddsId"`
  RewardType int `yaml:"RewardType"`
  AddRewardItemId int `yaml:"AddRewardItemId"`
  AddRewardItemQuantity int `yaml:"AddRewardItemQuantity"`
  AddRewardItemOdds int `yaml:"AddRewardItemOdds"`
  AddRewardItemOddsSum int `yaml:"AddRewardItemOddsSum"`
}

type Campaign struct {
  Id int `yaml:"Id"`
  CampaignType int `yaml:"CampaignType"`
  TargetContents int `yaml:"TargetContents"`
  Name string `yaml:"Name"`
  Description string `yaml:"Description"`
  EffectValue int `yaml:"EffectValue"`
  StartTime time.Time `yaml:"StartTime"`
  EndTime time.Time `yaml:"EndTime"`
}

type CardCoordinates struct {
  Id int64 `yaml:"Id"`
  CardDatasId int `yaml:"CardDatasId"`
  CardCoordSceneType int `yaml:"CardCoordSceneType"`
  XCoord int `yaml:"XCoord"`
  YCoord int `yaml:"YCoord"`
  Scale int `yaml:"Scale"`
}

type CardDatas struct {
  Id int `yaml:"Id"`
  CardSeriesId int `yaml:"CardSeriesId"`
  Name string `yaml:"Name"`
  Description string `yaml:"Description"`
  CharactersId int `yaml:"CharactersId"`
  Rarity int `yaml:"Rarity"`
  EvolveTimes int `yaml:"EvolveTimes"`
  CardLevelLimitAddition int `yaml:"CardLevelLimitAddition"`
  Style int `yaml:"Style"`
  Mood int `yaml:"Mood"`
  ExperienceType int `yaml:"ExperienceType"`
  InitialSmile int `yaml:"InitialSmile"`
  InitialPure int `yaml:"InitialPure"`
  InitialCool int `yaml:"InitialCool"`
  InitialMental int `yaml:"InitialMental"`
  MaxSmile int `yaml:"MaxSmile"`
  MaxPure int `yaml:"MaxPure"`
  MaxCool int `yaml:"MaxCool"`
  MaxMental int `yaml:"MaxMental"`
  BeatPoint int `yaml:"BeatPoint"`
  SpecialAppealSeriesId int `yaml:"SpecialAppealSeriesId"`
  SkillSeriesId int `yaml:"SkillSeriesId"`
  AttributeId int `yaml:"AttributeId"`
  SpineId int `yaml:"SpineId"`
  OrderId int `yaml:"OrderId"`
}

type CardDuetVoice struct {
  Id int `yaml:"Id"`
  CardSeriesId int `yaml:"CardSeriesId"`
  CharacterIds string `yaml:"CharacterIds"`
}

type CardEvolutionMaterials struct {
  Id int `yaml:"Id"`
  CostItemsId1 int `yaml:"CostItemsId1"`
  CostNum1 int `yaml:"CostNum1"`
  CostItemsId2 int `yaml:"CostItemsId2"`
  CostNum2 int `yaml:"CostNum2"`
  CostItemsId3 int `yaml:"CostItemsId3"`
  CostNum3 int `yaml:"CostNum3"`
}

type CardGetMovieSettings struct {
  Id int `yaml:"Id"`
  CardInfoPositionType int `yaml:"CardInfoPositionType"`
  CardInfoDisplayStartTimeSeconds int `yaml:"CardInfoDisplayStartTimeSeconds"`
  UrCardEffectBackgroundId int `yaml:"UrCardEffectBackgroundId"`
}

type CardLevels struct {
  Id int `yaml:"Id"`
  ExperienceType int `yaml:"ExperienceType"`
  CardLevel int `yaml:"CardLevel"`
  Experience int `yaml:"Experience"`
  CumulativeExperience int `yaml:"CumulativeExperience"`
}

type CardLimitBreakMaterials struct {
  Id int `yaml:"Id"`
  CardSeriesId int `yaml:"CardSeriesId"`
  LimitBreakTimes int `yaml:"LimitBreakTimes"`
  CostItemsId int `yaml:"CostItemsId"`
  CostNum int `yaml:"CostNum"`
}

type CardRarities struct {
  Id int `yaml:"Id"`
  RarityName string `yaml:"RarityName"`
  Evolution0_MaxLevel int `yaml:"Evolution0_MaxLevel"`
  Evolution1_MaxLevel int `yaml:"Evolution1_MaxLevel"`
  Evolution2_MaxLevel int `yaml:"Evolution2_MaxLevel"`
  Evolution3_MaxLevel int `yaml:"Evolution3_MaxLevel"`
  Evolution4_MaxLevel int `yaml:"Evolution4_MaxLevel"`
}

type CardSeries struct {
  Id int `yaml:"Id"`
  Evolution0Id int `yaml:"Evolution0Id"`
  Evolution1Id int `yaml:"Evolution1Id"`
  Evolution2Id int `yaml:"Evolution2Id"`
  ObtainFanLvPt int `yaml:"ObtainFanLvPt"`
  Evolution1FanLvPt int `yaml:"Evolution1FanLvPt"`
  Evolution2FanLvPt int `yaml:"Evolution2FanLvPt"`
  LimitBreak1FanLvPt int `yaml:"LimitBreak1FanLvPt"`
  LimitBreak2FanLvPt int `yaml:"LimitBreak2FanLvPt"`
  LimitBreak3FanLvPt int `yaml:"LimitBreak3FanLvPt"`
  LimitBreak4FanLvPt int `yaml:"LimitBreak4FanLvPt"`
  LimitedType int `yaml:"LimitedType"`
  SideStyleSettingCharacterId string `yaml:"SideStyleSettingCharacterId"`
  Evolution3Id int `yaml:"Evolution3Id"`
  Evolution4Id int `yaml:"Evolution4Id"`
  Evolution3FanLvPt int `yaml:"Evolution3FanLvPt"`
  Evolution4FanLvPt int `yaml:"Evolution4FanLvPt"`
  Evolution3RequiredFanLv int `yaml:"Evolution3RequiredFanLv"`
  Evolution4RequiredFanLv int `yaml:"Evolution4RequiredFanLv"`
}

type CardSkillEffectDetailParams struct {
  Id int64 `yaml:"Id"`
  ParamType string `yaml:"ParamType"`
  ParamValue string `yaml:"ParamValue"`
}

type CardSkillEffectDetails struct {
  Id int64 `yaml:"Id"`
  SkillEffectDetailType string `yaml:"SkillEffectDetailType"`
  TargetMood int `yaml:"TargetMood"`
  EffectValue int64 `yaml:"EffectValue"`
}

type CardSkillEffects struct {
  Id int64 `yaml:"Id"`
  ActionType int `yaml:"ActionType"`
  OrderId int `yaml:"OrderId"`
}

type CardSkillLevelUpMaterials struct {
  Id int64 `yaml:"Id"`
  CardSeriesId int `yaml:"CardSeriesId"`
  SkillType int `yaml:"SkillType"`
  SkillLevel int `yaml:"SkillLevel"`
  Cost_ItemsId1 int `yaml:"Cost_ItemsId1"`
  CostNum1 int `yaml:"CostNum1"`
  Cost_ItemsId2 int `yaml:"Cost_ItemsId2"`
  CostNum2 int `yaml:"CostNum2"`
  Cost_ItemsId3 int `yaml:"Cost_ItemsId3"`
  CostNum3 int `yaml:"CostNum3"`
}

type CardSkillModes struct {
  Id int64 `yaml:"Id"`
  ModeOffName string `yaml:"ModeOffName"`
  ModeOnName string `yaml:"ModeOnName"`
  CharactersId int `yaml:"CharactersId"`
}

type CardSkillSeries struct {
  Id int `yaml:"Id"`
  Name string `yaml:"Name"`
  SkillIcon int `yaml:"SkillIcon"`
  SkillMainEffect int `yaml:"SkillMainEffect"`
}

type CardSkills struct {
  Id int64 `yaml:"Id"`
  CardSkillSeriesId int `yaml:"CardSkillSeriesId"`
  SkillLevel int `yaml:"SkillLevel"`
  SkillCost int `yaml:"SkillCost"`
  ApperanceType int `yaml:"ApperanceType"`
  CardSkillEffectId string `yaml:"CardSkillEffectId"`
  Description string `yaml:"Description"`
}

type ChallengeModeEffectDetails struct {
  Id int64 `yaml:"Id"`
  EffectDetailType string `yaml:"EffectDetailType"`
  TargetMood int `yaml:"TargetMood"`
  EffectValue int `yaml:"EffectValue"`
}

type ChallengeModeEffects struct {
  Id int64 `yaml:"Id"`
  StandardQuestStagesId int `yaml:"StandardQuestStagesId"`
  ActionType int `yaml:"ActionType"`
  OrderId int `yaml:"OrderId"`
  Description string `yaml:"Description"`
}

type ChallengeModeReleaseCondition struct {
  Id int64 `yaml:"Id"`
  ReleaseChallengeModeId int `yaml:"ReleaseChallengeModeId"`
  ConditionsType int `yaml:"ConditionsType"`
  ConditionsValue int `yaml:"ConditionsValue"`
  StartTime time.Time `yaml:"StartTime"`
  EndTime time.Time `yaml:"EndTime"`
}

type ChallengeModeStages struct {
  Id int `yaml:"Id"`
  ChallengeModeAreasId int `yaml:"ChallengeModeAreasId"`
  CorrespondedQuestStageId int `yaml:"CorrespondedQuestStageId"`
  Name string `yaml:"Name"`
  Description string `yaml:"Description"`
  Hint string `yaml:"Hint"`
  MapNumber int `yaml:"MapNumber"`
  StageType int `yaml:"StageType"`
  UseType int `yaml:"UseType"`
  UseItem int `yaml:"UseItem"`
  UseNum int `yaml:"UseNum"`
  LiveStagesId int `yaml:"LiveStagesId"`
  QuestMusicsType int `yaml:"QuestMusicsType"`
  QuestMusicsDetail int `yaml:"QuestMusicsDetail"`
  DeckRestrictedType int `yaml:"DeckRestrictedType"`
  DeckRestrictedDetail int `yaml:"DeckRestrictedDetail"`
  ChallengeModeEffectId int64 `yaml:"ChallengeModeEffectId"`
  QuestLevel int `yaml:"QuestLevel"`
  FirstClearRewardSeriesId int64 `yaml:"FirstClearRewardSeriesId"`
  ChallengeModeScore int64 `yaml:"ChallengeModeScore"`
}

type CharacterFavoriteGifts struct {
  Id int `yaml:"Id"`
  ItemsId int `yaml:"ItemsId"`
  CharactersId int `yaml:"CharactersId"`
  FavoriteRank int `yaml:"FavoriteRank"`
}

type Characters struct {
  Id int `yaml:"Id"`
  NameLast string `yaml:"NameLast"`
  NameFirst string `yaml:"NameFirst"`
  LatinAlphabetNameLast string `yaml:"LatinAlphabetNameLast"`
  LatinAlphabetNameFirst string `yaml:"LatinAlphabetNameFirst"`
  GenerationsId int `yaml:"GenerationsId"`
  SeriesType int `yaml:"SeriesType"`
  IconOrderId int `yaml:"IconOrderId"`
  CharacterVoice string `yaml:"CharacterVoice"`
  ThemeColor string `yaml:"ThemeColor"`
  Introduction string `yaml:"Introduction"`
  ShowSeasonFanLvStartTime time.Time `yaml:"ShowSeasonFanLvStartTime"`
  ShowSeasonFanLvEndTime time.Time `yaml:"ShowSeasonFanLvEndTime"`
  IsExistFanLv int `yaml:"IsExistFanLv"`
  StyleType int `yaml:"StyleType"`
}

type Comics struct {
  Id int `yaml:"Id"`
  Name string `yaml:"Name"`
  ViewType int `yaml:"ViewType"`
  StartTime time.Time `yaml:"StartTime"`
  EndTime time.Time `yaml:"EndTime"`
  TabListId int `yaml:"TabListId"`
  AppearanceCharacterIds string `yaml:"AppearanceCharacterIds"`
}

type ContentGuidances struct {
  Id int `yaml:"Id"`
  ContentId int `yaml:"ContentId"`
  Priority int `yaml:"Priority"`
  IsSkip int `yaml:"IsSkip"`
  IsEnable int `yaml:"IsEnable"`
  ConditionValues string `yaml:"ConditionValues"`
}

type ContentsReleaseConditions struct {
  Id int `yaml:"Id"`
  ReleaseContentsId int `yaml:"ReleaseContentsId"`
  ConditionsType int `yaml:"ConditionsType"`
  ConditionsValue int `yaml:"ConditionsValue"`
  StartTime time.Time `yaml:"StartTime"`
  EndTime time.Time `yaml:"EndTime"`
}

type CostumeModels struct {
  Id int64 `yaml:"Id"`
  Label string `yaml:"Label"`
  CharactersId int `yaml:"CharactersId"`
  CostumesId int `yaml:"CostumesId"`
  HairStyleId int `yaml:"HairStyleId"`
}

type Costumes struct {
  Id int `yaml:"Id"`
  Label string `yaml:"Label"`
}

type CustomComplementMaterials struct {
  Id int `yaml:"Id"`
  Name string `yaml:"Name"`
  OrderId int `yaml:"OrderId"`
  StartTime time.Time `yaml:"StartTime"`
  EndTime time.Time `yaml:"EndTime"`
}

type DailyLiveReleaseConditions struct {
  Id int64 `yaml:"Id"`
  ReleaseDailyLivesId int `yaml:"ReleaseDailyLivesId"`
  ConditionsType int `yaml:"ConditionsType"`
  ConditionsValue int `yaml:"ConditionsValue"`
  StartTime time.Time `yaml:"StartTime"`
  EndTime time.Time `yaml:"EndTime"`
}

type DailyQuestSeries struct {
  Id int `yaml:"Id"`
  Name string `yaml:"Name"`
  Description string `yaml:"Description"`
  IsSunday int `yaml:"IsSunday"`
  IsMonday int `yaml:"IsMonday"`
  IsTuesday int `yaml:"IsTuesday"`
  IsWednesday int `yaml:"IsWednesday"`
  IsThursday int `yaml:"IsThursday"`
  IsFriday int `yaml:"IsFriday"`
  IsSaturday int `yaml:"IsSaturday"`
  StartTime time.Time `yaml:"StartTime"`
  EndTime time.Time `yaml:"EndTime"`
}

type DailyQuestStages struct {
  Id int `yaml:"Id"`
  DailyQuestSeriesId int `yaml:"DailyQuestSeriesId"`
  Name string `yaml:"Name"`
  Description string `yaml:"Description"`
  Hint string `yaml:"Hint"`
  StageType int `yaml:"StageType"`
  UseType int `yaml:"UseType"`
  UseItem int `yaml:"UseItem"`
  UseNum int `yaml:"UseNum"`
  LiveStagesId int `yaml:"LiveStagesId"`
  QuestMusicsType int `yaml:"QuestMusicsType"`
  QuestMusicsDetail int `yaml:"QuestMusicsDetail"`
  DeckRestrictedType int `yaml:"DeckRestrictedType"`
  DeckRestrictedDetail int `yaml:"DeckRestrictedDetail"`
  QuestLevel int `yaml:"QuestLevel"`
  QuestRank int `yaml:"QuestRank"`
  FirstClearRewardSeriesId int64 `yaml:"FirstClearRewardSeriesId"`
  CompleteRewardSeriesId int64 `yaml:"CompleteRewardSeriesId"`
  DropRewardSeriesId int64 `yaml:"DropRewardSeriesId"`
  RandomDropRewardSeriesId int64 `yaml:"RandomDropRewardSeriesId"`
  Score1 int64 `yaml:"Score1"`
  Score2 int64 `yaml:"Score2"`
  Score3 int64 `yaml:"Score3"`
  GainStylePoint int `yaml:"GainStylePoint"`
  GainMusicExp int `yaml:"GainMusicExp"`
  ItemSourceIconId int `yaml:"ItemSourceIconId"`
}

type DeckMemberPositions struct {
  Id int `yaml:"Id"`
  GenerationsId int `yaml:"GenerationsId"`
  CharactersId int `yaml:"CharactersId"`
  OrderId int `yaml:"OrderId"`
  StartTime time.Time `yaml:"StartTime"`
  EndTime time.Time `yaml:"EndTime"`
}

type DownloadImages struct {
  Id int64 `yaml:"Id"`
  DownloadType int `yaml:"DownloadType"`
  Title string `yaml:"Title"`
  OrderId int `yaml:"OrderId"`
  StartTime time.Time `yaml:"StartTime"`
  EndTime time.Time `yaml:"EndTime"`
}

type DreamLiveReleaseConditions struct {
  Id int64 `yaml:"Id"`
  ReleaseDreamLiveId int `yaml:"ReleaseDreamLiveId"`
  ConditionsType int `yaml:"ConditionsType"`
  ConditionsValue int `yaml:"ConditionsValue"`
  StartTime time.Time `yaml:"StartTime"`
  EndTime time.Time `yaml:"EndTime"`
}

type DreamLiveSeriesList struct {
  Id int `yaml:"Id"`
  DisplayPosition int `yaml:"DisplayPosition"`
  ImageId int `yaml:"ImageId"`
  StartTime time.Time `yaml:"StartTime"`
  EndTime time.Time `yaml:"EndTime"`
}

type DreamQuestSeries struct {
  Id int `yaml:"Id"`
  CharactersId int `yaml:"CharactersId"`
  Name string `yaml:"Name"`
  OrderId int `yaml:"OrderId"`
  StartTime time.Time `yaml:"StartTime"`
  EndTime time.Time `yaml:"EndTime"`
  DreamLiveSeriesListId int `yaml:"DreamLiveSeriesListId"`
}

type DreamQuestStages struct {
  Id int `yaml:"Id"`
  DreamQuestSeriesId int `yaml:"DreamQuestSeriesId"`
  Name string `yaml:"Name"`
  Description string `yaml:"Description"`
  Hint string `yaml:"Hint"`
  StageType int `yaml:"StageType"`
  UseType int `yaml:"UseType"`
  UseItem int `yaml:"UseItem"`
  UseNum int `yaml:"UseNum"`
  LiveStagesId int `yaml:"LiveStagesId"`
  QuestMusicsType int `yaml:"QuestMusicsType"`
  QuestMusicsDetail int `yaml:"QuestMusicsDetail"`
  DeckRestrictedType int `yaml:"DeckRestrictedType"`
  DeckRestrictedDetail int `yaml:"DeckRestrictedDetail"`
  FirstClearRewardSeriesId int64 `yaml:"FirstClearRewardSeriesId"`
  DropRewardSeriesId int64 `yaml:"DropRewardSeriesId"`
  RandomDropRewardSeriesId int64 `yaml:"RandomDropRewardSeriesId"`
  Score1 int64 `yaml:"Score1"`
}

type EmojiCategory struct {
  Id int `yaml:"Id"`
  Name string `yaml:"Name"`
  OrderId int `yaml:"OrderId"`
  StartTime time.Time `yaml:"StartTime"`
  EndTime time.Time `yaml:"EndTime"`
}

type Emojis struct {
  Id int `yaml:"Id"`
  Name string `yaml:"Name"`
  Alias string `yaml:"Alias"`
  Category int `yaml:"Category"`
  OrderId int `yaml:"OrderId"`
  RequirementType int `yaml:"RequirementType"`
  RequirementDetail string `yaml:"RequirementDetail"`
  RequirementValue int `yaml:"RequirementValue"`
  RequirementText string `yaml:"RequirementText"`
  IsVisibleOnlyPossess int `yaml:"IsVisibleOnlyPossess"`
  StartTime time.Time `yaml:"StartTime"`
  EndTime time.Time `yaml:"EndTime"`
  AvailableStartTime time.Time `yaml:"AvailableStartTime"`
  AvailableEndTime time.Time `yaml:"AvailableEndTime"`
}

type EventLoginBonuses struct {
  Id int `yaml:"Id"`
  Name string `yaml:"Name"`
  LoginBonusRewardSeriesId int `yaml:"LoginBonusRewardSeriesId"`
  EventLoginBonusType int `yaml:"EventLoginBonusType"`
  LoginBonusTextId string `yaml:"LoginBonusTextId"`
  StartTime time.Time `yaml:"StartTime"`
  EndTime time.Time `yaml:"EndTime"`
}

type EventMissionAchieveRewards struct {
  Id int `yaml:"Id"`
  EvemtMissionSeriesId int `yaml:"EvemtMissionSeriesId"`
  AchieveMarkNum int `yaml:"AchieveMarkNum"`
  RewardCategory int `yaml:"RewardCategory"`
  RewardType int `yaml:"RewardType"`
  ItemsId int `yaml:"ItemsId"`
  RewardNum int `yaml:"RewardNum"`
  RewardTextId int `yaml:"RewardTextId"`
  SortOrder int `yaml:"SortOrder"`
}

type EventMissionRewards struct {
  Id int `yaml:"Id"`
  EventMissionSeriesId int `yaml:"EventMissionSeriesId"`
  EventMissionsId int `yaml:"EventMissionsId"`
  RewardCategory int `yaml:"RewardCategory"`
  RewardType int `yaml:"RewardType"`
  ItemsId int `yaml:"ItemsId"`
  RewardNum int `yaml:"RewardNum"`
  RewardTextId int `yaml:"RewardTextId"`
  SortOrder int `yaml:"SortOrder"`
}

type EventMissionSeries struct {
  Id int `yaml:"Id"`
  Name string `yaml:"Name"`
  Description string `yaml:"Description"`
  GrandPrixesId int `yaml:"GrandPrixesId"`
  StartTime time.Time `yaml:"StartTime"`
  EndTime time.Time `yaml:"EndTime"`
}

type EventMissions struct {
  Id int `yaml:"Id"`
  EventMissionSeriesId int `yaml:"EventMissionSeriesId"`
  Name string `yaml:"Name"`
  Description string `yaml:"Description"`
  MissionType int `yaml:"MissionType"`
  MissionCondition int `yaml:"MissionCondition"`
  MissionConditionNum int `yaml:"MissionConditionNum"`
  OpenType int `yaml:"OpenType"`
  NextMissionsId int `yaml:"NextMissionsId"`
  SortOrder int `yaml:"SortOrder"`
  TransitionContentsId int `yaml:"TransitionContentsId"`
  MissionConditionDetail int `yaml:"MissionConditionDetail"`
}

type ExchangePointConvert struct {
  Id int `yaml:"Id"`
  ConvertItemType int `yaml:"ConvertItemType"`
  ConvertItemId int `yaml:"ConvertItemId"`
  ConvertItemQuantity int `yaml:"ConvertItemQuantity"`
  ConvertTime time.Time `yaml:"ConvertTime"`
}

type ExchangePointRate struct {
  Id int `yaml:"Id"`
  ExchangePointId int `yaml:"ExchangePointId"`
  ExchangeItemType int `yaml:"ExchangeItemType"`
  ExchangeItemId int `yaml:"ExchangeItemId"`
  ExchangeItemQuantity int `yaml:"ExchangeItemQuantity"`
  ExchangePrice int `yaml:"ExchangePrice"`
  LimitedCount int `yaml:"LimitedCount"`
  StartTime time.Time `yaml:"StartTime"`
  EndTime time.Time `yaml:"EndTime"`
  GachaSeriesId int `yaml:"GachaSeriesId"`
  BonusItemQuantity int `yaml:"BonusItemQuantity"`
}

type FlowerStandColors struct {
  Id int `yaml:"Id"`
  Name string `yaml:"Name"`
  ColorCode string `yaml:"ColorCode"`
  OrderId int `yaml:"OrderId"`
}

type FlowerStandIdolPictures struct {
  Id int `yaml:"Id"`
  Name string `yaml:"Name"`
  CharactersId int `yaml:"CharactersId"`
}

type FlowerStandTypes struct {
  Id int `yaml:"Id"`
  Name string `yaml:"Name"`
}

type GachaCampaigns struct {
  Id int `yaml:"Id"`
  CampaignName string `yaml:"CampaignName"`
  CampaignType int `yaml:"CampaignType"`
  ConsectiveTimesType int `yaml:"ConsectiveTimesType"`
  ResetType int `yaml:"ResetType"`
  PerDayCampaignTimes int `yaml:"PerDayCampaignTimes"`
  GachaSeriesId_1 int64 `yaml:"GachaSeriesId_1"`
  GachaSeriesId_2 int64 `yaml:"GachaSeriesId_2"`
  GachaSeriesId_3 int64 `yaml:"GachaSeriesId_3"`
  GachaSeriesId_4 int64 `yaml:"GachaSeriesId_4"`
  GachaSeriesId_5 int64 `yaml:"GachaSeriesId_5"`
  StartTime time.Time `yaml:"StartTime"`
  EndTime time.Time `yaml:"EndTime"`
  MiniBannerPopId int `yaml:"MiniBannerPopId"`
}

type GachaSeries struct {
  Id int `yaml:"Id"`
  GachaSeriesName string `yaml:"GachaSeriesName"`
  Description string `yaml:"Description"`
  GachaType int `yaml:"GachaType"`
  LimitedGachaCount int `yaml:"LimitedGachaCount"`
  LimitedGachaResetType int `yaml:"LimitedGachaResetType"`
  GachaExchangePointId int `yaml:"GachaExchangePointId"`
  ExchangePointNoticeNum int `yaml:"ExchangePointNoticeNum"`
  ExchangePointLockFlag int `yaml:"ExchangePointLockFlag"`
  OrderId int `yaml:"OrderId"`
  FilterType int `yaml:"FilterType"`
  PickUpCardSeriesId_1 int `yaml:"PickUpCardSeriesId_1"`
  PickUpCardBonusItemQuantity_1 int `yaml:"PickUpCardBonusItemQuantity_1"`
  PickUpCardSeriesId_2 int `yaml:"PickUpCardSeriesId_2"`
  PickUpCardBonusItemQuantity_2 int `yaml:"PickUpCardBonusItemQuantity_2"`
  PickUpCardSeriesId_3 int `yaml:"PickUpCardSeriesId_3"`
  PickUpCardBonusItemQuantity_3 int `yaml:"PickUpCardBonusItemQuantity_3"`
  PickUpCardSeriesId_4 int `yaml:"PickUpCardSeriesId_4"`
  PickUpCardBonusItemQuantity_4 int `yaml:"PickUpCardBonusItemQuantity_4"`
  PickUpCardSeriesId_5 int `yaml:"PickUpCardSeriesId_5"`
  PickUpCardBonusItemQuantity_5 int `yaml:"PickUpCardBonusItemQuantity_5"`
  PickUpCardSeriesId_6 int `yaml:"PickUpCardSeriesId_6"`
  PickUpCardBonusItemQuantity_6 int `yaml:"PickUpCardBonusItemQuantity_6"`
  BgType int `yaml:"BgType"`
  StartTime time.Time `yaml:"StartTime"`
  EndTime time.Time `yaml:"EndTime"`
  NoticeText string `yaml:"NoticeText"`
  GachaTimeLimitType int `yaml:"GachaTimeLimitType"`
  AvailableTime int `yaml:"AvailableTime"`
}

type Generations struct {
  Id int `yaml:"Id"`
  Name string `yaml:"Name"`
  StartTime time.Time `yaml:"StartTime"`
  EndTime time.Time `yaml:"EndTime"`
}

type GiftBonusGachas struct {
  Id int `yaml:"Id"`
  SingleGachaPrice int `yaml:"SingleGachaPrice"`
  SingleGachaPopId int `yaml:"SingleGachaPopId"`
  ConsectiveGachaPrice int `yaml:"ConsectiveGachaPrice"`
  ConsectiveGachaTimes int `yaml:"ConsectiveGachaTimes"`
  ConsectiveGachaPopId int `yaml:"ConsectiveGachaPopId"`
  PaidSIsCaOnlyGachaFlag int `yaml:"PaidSIsCaOnlyGachaFlag"`
  PaidSIsCaOnlyGachaPrice int `yaml:"PaidSIsCaOnlyGachaPrice"`
  PaidSIsCaOnlyGachaTimes int `yaml:"PaidSIsCaOnlyGachaTimes"`
  PaidSIsCaOnlyGachaPointFlag int `yaml:"PaidSIsCaOnlyGachaPointFlag"`
  PaidSIsCaOnlyGachaaPopId int `yaml:"PaidSIsCaOnlyGachaaPopId"`
  LimitedPaidGachaCount int `yaml:"LimitedPaidGachaCount"`
  LimitedPaidGachaResetType int `yaml:"LimitedPaidGachaResetType"`
  LimitedPaidButtonDesignType int `yaml:"LimitedPaidButtonDesignType"`
  MiniBannerPopId int `yaml:"MiniBannerPopId"`
}

type GiftlessGachas struct {
  Id int `yaml:"Id"`
  SingleGachaPrice int `yaml:"SingleGachaPrice"`
  SingleGachaPopId int `yaml:"SingleGachaPopId"`
  ConsectiveGachaPrice int `yaml:"ConsectiveGachaPrice"`
  ConsectiveGachaTimes int `yaml:"ConsectiveGachaTimes"`
  ConsectiveGachaPopId int `yaml:"ConsectiveGachaPopId"`
  PaidSIsCaOnlyGachaFlag int `yaml:"PaidSIsCaOnlyGachaFlag"`
  PaidSIsCaOnlyGachaPrice int `yaml:"PaidSIsCaOnlyGachaPrice"`
  PaidSIsCaOnlyGachaTimes int `yaml:"PaidSIsCaOnlyGachaTimes"`
  PaidSIsCaOnlyGachaPointFlag int `yaml:"PaidSIsCaOnlyGachaPointFlag"`
  PaidSIsCaOnlyGachaaPopId int `yaml:"PaidSIsCaOnlyGachaaPopId"`
  LimitedPaidGachaCount int `yaml:"LimitedPaidGachaCount"`
  LimitedPaidGachaResetType int `yaml:"LimitedPaidGachaResetType"`
  LimitedPaidButtonDesignType int `yaml:"LimitedPaidButtonDesignType"`
  MiniBannerPopId int `yaml:"MiniBannerPopId"`
}

type GpPrizeExchanges struct {
  Id int `yaml:"Id"`
  ProductItemType int `yaml:"ProductItemType"`
  ProductItemId int `yaml:"ProductItemId"`
  ProductItemNum int `yaml:"ProductItemNum"`
  MaterialItemType int `yaml:"MaterialItemType"`
  MaterialItemId int `yaml:"MaterialItemId"`
  MaterialItemNum int `yaml:"MaterialItemNum"`
  LimitNum int `yaml:"LimitNum"`
  ResetType int `yaml:"ResetType"`
  OrderId int `yaml:"OrderId"`
  StartTime time.Time `yaml:"StartTime"`
  EndTime time.Time `yaml:"EndTime"`
}

type GradeAddSkillEffectDetails struct {
  Id int64 `yaml:"Id"`
  SkillEffectDetailType string `yaml:"SkillEffectDetailType"`
  TargetMood int `yaml:"TargetMood"`
  EffectValue int64 `yaml:"EffectValue"`
}

type GradeAddSkillEffects struct {
  Id int64 `yaml:"Id"`
  ActionType int `yaml:"ActionType"`
  OrderId int `yaml:"OrderId"`
}

type GradeAddSkills struct {
  Id int64 `yaml:"Id"`
  Name string `yaml:"Name"`
  Description string `yaml:"Description"`
  SkillIcon int `yaml:"SkillIcon"`
  GradeAddSkillEffectsId string `yaml:"GradeAddSkillEffectsId"`
  OrderId int `yaml:"OrderId"`
}

type GradeChalQuestStageRewardDatas struct {
  Id int64 `yaml:"Id"`
  GradeChalQuestStagesRewardsId int64 `yaml:"GradeChalQuestStagesRewardsId"`
  RewardType int `yaml:"RewardType"`
  RewardItemId int `yaml:"RewardItemId"`
  RewardNum int `yaml:"RewardNum"`
  OrderId int `yaml:"OrderId"`
}

type GradeChalQuestStages struct {
  Id int `yaml:"Id"`
  GradeChalSeasonId int `yaml:"GradeChalSeasonId"`
  Name string `yaml:"Name"`
  Description string `yaml:"Description"`
  Hint string `yaml:"Hint"`
  CharacterId int `yaml:"CharacterId"`
  OrderId int `yaml:"OrderId"`
  StageType int `yaml:"StageType"`
  StageIconId int `yaml:"StageIconId"`
  LiveStagesId int `yaml:"LiveStagesId"`
  QuestMusicsType int `yaml:"QuestMusicsType"`
  QuestMusicsDetail int `yaml:"QuestMusicsDetail"`
  DeckRestrictedType int `yaml:"DeckRestrictedType"`
  DeckRestrictedDetail int `yaml:"DeckRestrictedDetail"`
}

type GradeChalQuestStagesRewards struct {
  Id int64 `yaml:"Id"`
  GradeChalQuestStagesId int `yaml:"GradeChalQuestStagesId"`
  OrderId int `yaml:"OrderId"`
  ConditionsType int `yaml:"ConditionsType"`
  ConditionsValue int `yaml:"ConditionsValue"`
}

type GradeChalSeason struct {
  Id int `yaml:"Id"`
  Name string `yaml:"Name"`
  GiveSeasonGradeId int `yaml:"GiveSeasonGradeId"`
  BgImageId int `yaml:"BgImageId"`
  StartTime time.Time `yaml:"StartTime"`
  EndTime time.Time `yaml:"EndTime"`
}

type GradeChalTotalScoreRewardDatas struct {
  Id int `yaml:"Id"`
  GradeChalTotalScoreRewardsId int `yaml:"GradeChalTotalScoreRewardsId"`
  RewardType int `yaml:"RewardType"`
  RewardItemId int `yaml:"RewardItemId"`
  RewardNum int `yaml:"RewardNum"`
  OrderId int `yaml:"OrderId"`
}

type GradeChalTotalScoreRewards struct {
  Id int `yaml:"Id"`
  GradeChalSeasonId int `yaml:"GradeChalSeasonId"`
  TotalScore int64 `yaml:"TotalScore"`
  OrderId int `yaml:"OrderId"`
}

type GradeDatas struct {
  Id int `yaml:"Id"`
  GradeNum int `yaml:"GradeNum"`
  GradeId int `yaml:"GradeId"`
  GradeTier int `yaml:"GradeTier"`
  RequireGradePoints int `yaml:"RequireGradePoints"`
  GradeLiveBonus int `yaml:"GradeLiveBonus"`
}

type GradeQuestLivePointBonus struct {
  Id int `yaml:"Id"`
  GradeQuestSeriesId int `yaml:"GradeQuestSeriesId"`
  OrderId int `yaml:"OrderId"`
  Description string `yaml:"Description"`
  ConditionsType int `yaml:"ConditionsType"`
  ConditionsValue int `yaml:"ConditionsValue"`
  BonusNum int `yaml:"BonusNum"`
  BonusLimitUp int `yaml:"BonusLimitUp"`
}

type GradeQuestRewardsDatas struct {
  Id int64 `yaml:"Id"`
  GradeQuestRewardsId int `yaml:"GradeQuestRewardsId"`
  RewardType int `yaml:"RewardType"`
  RewardItemId int `yaml:"RewardItemId"`
  RewardNum int `yaml:"RewardNum"`
}

type GradeQuestRewards struct {
  Id int `yaml:"Id"`
  GradeQuestSeriesId int `yaml:"GradeQuestSeriesId"`
  OrderId int `yaml:"OrderId"`
  ConditionsType int `yaml:"ConditionsType"`
  ConditionsValue int `yaml:"ConditionsValue"`
  ConditionsValue2 int `yaml:"ConditionsValue2"`
  ConditionsDescription string `yaml:"ConditionsDescription"`
  DisplayType int `yaml:"DisplayType"`
}

type GradeQuestSeason struct {
  Id int `yaml:"Id"`
  Name string `yaml:"Name"`
  OrderId int `yaml:"OrderId"`
  Generation int `yaml:"Generation"`
  Season int `yaml:"Season"`
  StartTime time.Time `yaml:"StartTime"`
  EndTime time.Time `yaml:"EndTime"`
}

type GradeQuestSeasonReleaseCond struct {
  Id int `yaml:"Id"`
  GradeQuestSeasonId int `yaml:"GradeQuestSeasonId"`
  ConditionsType int `yaml:"ConditionsType"`
  ConditionsValue int `yaml:"ConditionsValue"`
  ConditionsDescription string `yaml:"ConditionsDescription"`
  StartTime time.Time `yaml:"StartTime"`
  EndTime time.Time `yaml:"EndTime"`
}

type GradeQuestSeries struct {
  Id int `yaml:"Id"`
  Name string `yaml:"Name"`
  GradeQuestSeasonId int `yaml:"GradeQuestSeasonId"`
  OrderId int `yaml:"OrderId"`
  BackGroundId int `yaml:"BackGroundId"`
  MapImageId string `yaml:"MapImageId"`
  SoundId int `yaml:"SoundId"`
  IsTutorial int `yaml:"IsTutorial"`
  DefaultGradeAddSkillsId string `yaml:"DefaultGradeAddSkillsId"`
  DefaultActionPoint int `yaml:"DefaultActionPoint"`
  LivePointBonusLimit int `yaml:"LivePointBonusLimit"`
  StartTime time.Time `yaml:"StartTime"`
  EndTime time.Time `yaml:"EndTime"`
}

type GradeQuestSeriesReleaseCond struct {
  Id int `yaml:"Id"`
  GradeQuestSeriesId int `yaml:"GradeQuestSeriesId"`
  ConditionsType int `yaml:"ConditionsType"`
  ConditionsValue int `yaml:"ConditionsValue"`
  StartTime time.Time `yaml:"StartTime"`
  EndTime time.Time `yaml:"EndTime"`
}

type GradeQuestSquareDatas struct {
  Id int `yaml:"Id"`
  Name string `yaml:"Name"`
  SquareType int `yaml:"SquareType"`
  TargetId int64 `yaml:"TargetId"`
  MinActionPoint int `yaml:"MinActionPoint"`
  MaxActionPoint int `yaml:"MaxActionPoint"`
}

type GradeQuestSquare struct {
  Id int `yaml:"Id"`
  GradeQuestSeriesId int `yaml:"GradeQuestSeriesId"`
  SquareId int `yaml:"SquareId"`
  XCoord int `yaml:"XCoord"`
  YCoord int `yaml:"YCoord"`
  OpenGradeQuestSquareIds string `yaml:"OpenGradeQuestSquareIds"`
}

type GradeQuestStages struct {
  Id int `yaml:"Id"`
  Name string `yaml:"Name"`
  Description string `yaml:"Description"`
  Hint string `yaml:"Hint"`
  StageType int `yaml:"StageType"`
  LiveStagesId int `yaml:"LiveStagesId"`
  QuestMusicsType int `yaml:"QuestMusicsType"`
  QuestMusicsDetail int `yaml:"QuestMusicsDetail"`
  DeckRestrictedType int `yaml:"DeckRestrictedType"`
  DeckRestrictedDetail int `yaml:"DeckRestrictedDetail"`
  LivePoint int64 `yaml:"LivePoint"`
}

type Grade struct {
  Id int `yaml:"Id"`
  Name string `yaml:"Name"`
  MaxGradeTier int `yaml:"MaxGradeTier"`
}

type GradeRewardDatas struct {
  Id int `yaml:"Id"`
  GradeRewardsId int `yaml:"GradeRewardsId"`
  RewardType int `yaml:"RewardType"`
  RewardItemId int `yaml:"RewardItemId"`
  RewardNum int `yaml:"RewardNum"`
}

type GradeRewards struct {
  Id int `yaml:"Id"`
  GradeNum int `yaml:"GradeNum"`
}

type GrandPrixPointBonuses struct {
  Id int64 `yaml:"Id"`
  GrandPrixesId int `yaml:"GrandPrixesId"`
  TargetType int `yaml:"TargetType"`
  TargetDetail int `yaml:"TargetDetail"`
  TargetNum int `yaml:"TargetNum"`
  BonusValue int `yaml:"BonusValue"`
}

type GrandPrixQuestSeries struct {
  Id int `yaml:"Id"`
  Name string `yaml:"Name"`
  Description string `yaml:"Description"`
  GrandPrixesId int `yaml:"GrandPrixesId"`
  PlayLimitCount int `yaml:"PlayLimitCount"`
  RetireLimitCount int `yaml:"RetireLimitCount"`
  OrderId int `yaml:"OrderId"`
  SeriesNum int `yaml:"SeriesNum"`
  StartTime time.Time `yaml:"StartTime"`
  EndTime time.Time `yaml:"EndTime"`
}

type GrandPrixQuestStages struct {
  Id int `yaml:"Id"`
  GrandPrixSeriesId int `yaml:"GrandPrixSeriesId"`
  Name string `yaml:"Name"`
  Description string `yaml:"Description"`
  Hint string `yaml:"Hint"`
  OrderId int `yaml:"OrderId"`
  StageType int `yaml:"StageType"`
  LiveStagesId int `yaml:"LiveStagesId"`
  QuestMusicsType int `yaml:"QuestMusicsType"`
  QuestMusicsDetail int `yaml:"QuestMusicsDetail"`
  DeckRestrictedType int `yaml:"DeckRestrictedType"`
  DeckRestrictedDetail int `yaml:"DeckRestrictedDetail"`
  QuestLevel int `yaml:"QuestLevel"`
  QuestRank int `yaml:"QuestRank"`
  FirstClearRewardSeriesId int64 `yaml:"FirstClearRewardSeriesId"`
  CompleteRewardSeriesId int64 `yaml:"CompleteRewardSeriesId"`
  DropRewardSeriesId int64 `yaml:"DropRewardSeriesId"`
  RandomDropRewardSeriesId int64 `yaml:"RandomDropRewardSeriesId"`
  Score1 int64 `yaml:"Score1"`
  Score2 int64 `yaml:"Score2"`
  Score3 int64 `yaml:"Score3"`
  StylePoint int `yaml:"StylePoint"`
  GainMusicExp int `yaml:"GainMusicExp"`
  ScoreBonusValue0 int `yaml:"ScoreBonusValue0"`
  ScoreBonusValue1 int `yaml:"ScoreBonusValue1"`
  ScoreBonusValue2 int `yaml:"ScoreBonusValue2"`
  ScoreBonusValue3 int `yaml:"ScoreBonusValue3"`
}

type GrandPrix struct {
  Id int `yaml:"Id"`
  Name string `yaml:"Name"`
  Description string `yaml:"Description"`
  GrandPrixType int `yaml:"GrandPrixType"`
  GuildRankingTabs string `yaml:"GuildRankingTabs"`
  PersonalRankingTabs string `yaml:"PersonalRankingTabs"`
  GuildPresentCommentId int `yaml:"GuildPresentCommentId"`
  PersonalPresentCommentId int `yaml:"PersonalPresentCommentId"`
  InfoStartTime time.Time `yaml:"InfoStartTime"`
  InfoEndTime time.Time `yaml:"InfoEndTime"`
  StartTime time.Time `yaml:"StartTime"`
  EndTime time.Time `yaml:"EndTime"`
}

type GrandPrixReleaseCondition struct {
  Id int64 `yaml:"Id"`
  ReleaseGrandPrixId int `yaml:"ReleaseGrandPrixId"`
  ConditionsType int `yaml:"ConditionsType"`
  ConditionsValue int `yaml:"ConditionsValue"`
  StartTime time.Time `yaml:"StartTime"`
  EndTime time.Time `yaml:"EndTime"`
}

type GrandPrixRewardDatas struct {
  Id int64 `yaml:"Id"`
  GrandPrixRewardsId int `yaml:"GrandPrixRewardsId"`
  RewardType int `yaml:"RewardType"`
  RewardItemId int `yaml:"RewardItemId"`
  RewardNum int `yaml:"RewardNum"`
  IsEmphasize int `yaml:"IsEmphasize"`
  IsUsePresentBox int `yaml:"IsUsePresentBox"`
  LifeTimeDay int `yaml:"LifeTimeDay"`
  RewardTextId int `yaml:"RewardTextId"`
}

type GrandPrixRewards struct {
  Id int `yaml:"Id"`
  GrandPrixesId int `yaml:"GrandPrixesId"`
  GrandPrixRewardType int `yaml:"GrandPrixRewardType"`
  MinTargetNum int `yaml:"MinTargetNum"`
  MaxTargetNum int `yaml:"MaxTargetNum"`
  IsDisplay int `yaml:"IsDisplay"`
}

type HelpImages struct {
  Id int `yaml:"Id"`
  Name string `yaml:"Name"`
  Scene int `yaml:"Scene"`
  OrderId int `yaml:"OrderId"`
  TitleTextId int `yaml:"TitleTextId"`
}

type HomeBgms struct {
  Id int `yaml:"Id"`
  DaytimeBgmId int `yaml:"DaytimeBgmId"`
  NighttimeBgmId int `yaml:"NighttimeBgmId"`
  StartTime time.Time `yaml:"StartTime"`
  EndTime time.Time `yaml:"EndTime"`
}

type ItemExchangeCategoryDatas struct {
  Id int `yaml:"Id"`
  ItemExchangeCategoryName string `yaml:"ItemExchangeCategoryName"`
  ItemId int `yaml:"ItemId"`
}

type ItemExchanges struct {
  Id int `yaml:"Id"`
  ItemExchangeCategoryId int `yaml:"ItemExchangeCategoryId"`
  ProductItemType int `yaml:"ProductItemType"`
  ProductItemId int `yaml:"ProductItemId"`
  ProductItemNum int `yaml:"ProductItemNum"`
  MaterialItemType int `yaml:"MaterialItemType"`
  MaterialItemId int `yaml:"MaterialItemId"`
  MaterialItemNum int `yaml:"MaterialItemNum"`
  LimitNum int `yaml:"LimitNum"`
  ResetType int `yaml:"ResetType"`
  IsVisibleEndTime int `yaml:"IsVisibleEndTime"`
  OrderId int `yaml:"OrderId"`
  StartTime time.Time `yaml:"StartTime"`
  EndTime time.Time `yaml:"EndTime"`
}

type ItemSources struct {
  ItemsId int `yaml:"ItemsId"`
  StandardQuestStagesId string `yaml:"StandardQuestStagesId"`
  DailyQuestStagesId string `yaml:"DailyQuestStagesId"`
}

type Items struct {
  Id int `yaml:"Id"`
  Name string `yaml:"Name"`
  NameFurigana string `yaml:"NameFurigana"`
  ItemType int `yaml:"ItemType"`
  ItemCategory int `yaml:"ItemCategory"`
  Rarity int `yaml:"Rarity"`
  EffectValue int `yaml:"EffectValue"`
  LimitNum int `yaml:"LimitNum"`
  RequestableNum int `yaml:"RequestableNum"`
  Description string `yaml:"Description"`
  StartTime time.Time `yaml:"StartTime"`
  EndTime time.Time `yaml:"EndTime"`
}

type LauncherBanners struct {
  Id int `yaml:"Id"`
  DisplayPosition int `yaml:"DisplayPosition"`
  ImageId int `yaml:"ImageId"`
  StartTime time.Time `yaml:"StartTime"`
  EndTime time.Time `yaml:"EndTime"`
}

type LearningLiveReleaseConditions struct {
  Id int64 `yaml:"Id"`
  ReleaseLearningLiveId int `yaml:"ReleaseLearningLiveId"`
  ConditionsType int `yaml:"ConditionsType"`
  ConditionsValue int `yaml:"ConditionsValue"`
  StartTime time.Time `yaml:"StartTime"`
  EndTime time.Time `yaml:"EndTime"`
}

type LimitBreakMaterialConvertRate struct {
  Id int `yaml:"Id"`
  LimitBreakMaterialRarity int `yaml:"LimitBreakMaterialRarity"`
  LimitBreakMaterialNum int `yaml:"LimitBreakMaterialNum"`
  ProductItemType int `yaml:"ProductItemType"`
  ProductItemId int `yaml:"ProductItemId"`
  ProductItemNum int `yaml:"ProductItemNum"`
}

type LimitBreakMaterialRate struct {
  Id int `yaml:"Id"`
  CardSeriesId int `yaml:"CardSeriesId"`
  LimitBreakMaterialId int `yaml:"LimitBreakMaterialId"`
  LimitBreakMaterialQuantity int `yaml:"LimitBreakMaterialQuantity"`
}

type LiveChannels struct {
  Id int `yaml:"Id"`
  Name string `yaml:"Name"`
  Description string `yaml:"Description"`
  OrderId int `yaml:"OrderId"`
  StartTime time.Time `yaml:"StartTime"`
  EndTime time.Time `yaml:"EndTime"`
}

type LiveStages struct {
  Id int `yaml:"Id"`
  Name string `yaml:"Name"`
  Description string `yaml:"Description"`
  BaseUseMental int `yaml:"BaseUseMental"`
  MentalStepClass int `yaml:"MentalStepClass"`
  BaseGainVoltage int `yaml:"BaseGainVoltage"`
  VoltageStepClass int `yaml:"VoltageStepClass"`
  BackGroundId int `yaml:"BackGroundId"`
  StageSkillConditionId int64 `yaml:"StageSkillConditionId"`
  StageSkillEffectId int64 `yaml:"StageSkillEffectId"`
  StageSkillDescription string `yaml:"StageSkillDescription"`
  StageSkillSetIds string `yaml:"StageSkillSetIds"`
}

type LoginBonuses struct {
  Id int `yaml:"Id"`
  Name string `yaml:"Name"`
  LoginBonusRewardSeriesID int `yaml:"LoginBonusRewardSeriesID"`
  IsLoop int `yaml:"IsLoop"`
  LoginBonusTextId string `yaml:"LoginBonusTextId"`
  StartTime time.Time `yaml:"StartTime"`
  EndTime time.Time `yaml:"EndTime"`
}

type LoginBonusRewardDatas struct {
  Id int `yaml:"Id"`
  LoginBonusRewardSeriesId int `yaml:"LoginBonusRewardSeriesId"`
  DayCount int `yaml:"DayCount"`
  RewardType int `yaml:"RewardType"`
  RewardId int `yaml:"RewardId"`
  RewardNum int `yaml:"RewardNum"`
  LifeTimeDay int `yaml:"LifeTimeDay"`
  RewardTextId int `yaml:"RewardTextId"`
}

type MemberFanLevels struct {
  Id int `yaml:"Id"`
  MemberFanLevel int `yaml:"MemberFanLevel"`
  Experience int `yaml:"Experience"`
  CumulativeExperience int `yaml:"CumulativeExperience"`
}

type MemberMovies struct {
  Id int `yaml:"Id"`
  CharactersId int `yaml:"CharactersId"`
  MovieType int `yaml:"MovieType"`
  Name string `yaml:"Name"`
  Priority int `yaml:"Priority"`
  ReleaseConditionText string `yaml:"ReleaseConditionText"`
}

type MemberVoices struct {
  Id int `yaml:"Id"`
  CharactersId int `yaml:"CharactersId"`
  Name string `yaml:"Name"`
  Priority int `yaml:"Priority"`
  VoiceName string `yaml:"VoiceName"`
  ReleaseConditionText string `yaml:"ReleaseConditionText"`
}

type MissionAchieveRewards struct {
  Id int `yaml:"Id"`
  MissionType int `yaml:"MissionType"`
  AchieveMarkNum int `yaml:"AchieveMarkNum"`
  RewardType int `yaml:"RewardType"`
  ItemsId int `yaml:"ItemsId"`
  RewardNum int `yaml:"RewardNum"`
  RewardTextId int `yaml:"RewardTextId"`
  SortOrder int `yaml:"SortOrder"`
}

type MissionRewards struct {
  Id int `yaml:"Id"`
  MissionsId int `yaml:"MissionsId"`
  RewardCategory int `yaml:"RewardCategory"`
  RewardType int `yaml:"RewardType"`
  ItemsId int `yaml:"ItemsId"`
  RewardNum int `yaml:"RewardNum"`
  RewardTextId int `yaml:"RewardTextId"`
}

type Missions struct {
  Id int `yaml:"Id"`
  Name string `yaml:"Name"`
  Description string `yaml:"Description"`
  MissionType int `yaml:"MissionType"`
  MissionConditionType int `yaml:"MissionConditionType"`
  MissionConditionNum int `yaml:"MissionConditionNum"`
  MissionConditionDetail int `yaml:"MissionConditionDetail"`
  OpenType int `yaml:"OpenType"`
  NextMissionsId int `yaml:"NextMissionsId"`
  SortOrder int `yaml:"SortOrder"`
  TransitionContentsId int `yaml:"TransitionContentsId"`
  StartTime time.Time `yaml:"StartTime"`
  EndTime time.Time `yaml:"EndTime"`
  Hint int `yaml:"Hint"`
}

type MusicLearningQuestSeries struct {
  Id int `yaml:"Id"`
  MusicsId int `yaml:"MusicsId"`
  Name string `yaml:"Name"`
  StartTime time.Time `yaml:"StartTime"`
  EndTime time.Time `yaml:"EndTime"`
}

type MusicLearningQuestStages struct {
  Id int `yaml:"Id"`
  LearningLiveSeriesId int `yaml:"LearningLiveSeriesId"`
  Name string `yaml:"Name"`
  Description string `yaml:"Description"`
  Hint string `yaml:"Hint"`
  StageType int `yaml:"StageType"`
  UseType int `yaml:"UseType"`
  UseItem int `yaml:"UseItem"`
  UseNum int `yaml:"UseNum"`
  LiveStagesId int `yaml:"LiveStagesId"`
  QuestMusicsType int `yaml:"QuestMusicsType"`
  QuestMusicsDetail int `yaml:"QuestMusicsDetail"`
  DeckRestrictedType int `yaml:"DeckRestrictedType"`
  DeckRestrictedDetail int `yaml:"DeckRestrictedDetail"`
  QuestLevel int `yaml:"QuestLevel"`
  QuestRank int `yaml:"QuestRank"`
  FirstClearRewardSeriesId int64 `yaml:"FirstClearRewardSeriesId"`
  DropRewardSeriesId int64 `yaml:"DropRewardSeriesId"`
  RandomDropRewardSeriesId int64 `yaml:"RandomDropRewardSeriesId"`
  Score1 int64 `yaml:"Score1"`
  GainStylePoint int `yaml:"GainStylePoint"`
  GainMusicExp int `yaml:"GainMusicExp"`
}

type MusicLevels struct {
  Id int `yaml:"Id"`
  ExperienceType int `yaml:"ExperienceType"`
  Level int `yaml:"Level"`
  Experience int `yaml:"Experience"`
  CumulativeExperience int `yaml:"CumulativeExperience"`
}

type MusicMasteryHeartBonuses struct {
  Id int `yaml:"Id"`
  Level int `yaml:"Level"`
  LoveRate int `yaml:"LoveRate"`
}

type MusicMasteryLevels struct {
  Id int `yaml:"Id"`
  MusicsId int `yaml:"MusicsId"`
  Level int `yaml:"Level"`
  MusicMasterySkillsId int `yaml:"MusicMasterySkillsId"`
}

type MusicMasteryLoveBonuses struct {
  Id int `yaml:"Id"`
  Level int `yaml:"Level"`
  LoveRate int `yaml:"LoveRate"`
}

type MusicMasteryMentalBonuses struct {
  Id int `yaml:"Id"`
  Level int `yaml:"Level"`
  DemandDamagePt int `yaml:"DemandDamagePt"`
}

type MusicMasterySkill struct {
  Id int `yaml:"Id"`
  MusicMasterySkillsName string `yaml:"MusicMasterySkillsName"`
}

type MusicMasteryVoltageBonuses struct {
  Id int `yaml:"Id"`
  Level int `yaml:"Level"`
  DemandVoltagePt int `yaml:"DemandVoltagePt"`
}

type Musics struct {
  Id int `yaml:"Id"`
  OrderId int `yaml:"OrderId"`
  Title string `yaml:"Title"`
  TitleFurigana string `yaml:"TitleFurigana"`
  JacketId int `yaml:"JacketId"`
  SoundId int `yaml:"SoundId"`
  Description string `yaml:"Description"`
  GenerationsId int `yaml:"GenerationsId"`
  UnitId int `yaml:"UnitId"`
  CenterCharacterId int `yaml:"CenterCharacterId"`
  SingerCharacterId string `yaml:"SingerCharacterId"`
  SupportCharacterId string `yaml:"SupportCharacterId"`
  MusicType int `yaml:"MusicType"`
  ExperienceType int `yaml:"ExperienceType"`
  BeatPointCoefficient int `yaml:"BeatPointCoefficient"`
  ApIncrement int `yaml:"ApIncrement"`
  SongTime int `yaml:"SongTime"`
  PlayTime int `yaml:"PlayTime"`
  FeverSectionNo int `yaml:"FeverSectionNo"`
  PreviewStartTime int `yaml:"PreviewStartTime"`
  PreviewEndTime int `yaml:"PreviewEndTime"`
  PreviewFadeInTime int `yaml:"PreviewFadeInTime"`
  PreviewFadeOutTime int `yaml:"PreviewFadeOutTime"`
  ReleaseConditionType int `yaml:"ReleaseConditionType"`
  ReleaseConditionDetail int `yaml:"ReleaseConditionDetail"`
  ReleaseConditionText string `yaml:"ReleaseConditionText"`
  StartTime time.Time `yaml:"StartTime"`
  EndTime time.Time `yaml:"EndTime"`
  MaxAp int `yaml:"MaxAp"`
  IsVideoMode int `yaml:"IsVideoMode"`
  VideoBgId int `yaml:"VideoBgId"`
}

type PetalCoinExchangeRate struct {
  Id int `yaml:"Id"`
  Rarity int `yaml:"Rarity"`
  PetalCoinQuantity int `yaml:"PetalCoinQuantity"`
}

type PetalExchangeRates struct {
  Id int `yaml:"Id"`
  Rarity int `yaml:"Rarity"`
  Price int `yaml:"Price"`
  ExchangeLimitLower int `yaml:"ExchangeLimitLower"`
  ExchangeLimitUpper int `yaml:"ExchangeLimitUpper"`
}

type PresentTexts struct {
  Id int `yaml:"Id"`
  Description string `yaml:"Description"`
}

type QuestAreaReleaseConditions struct {
  Id int `yaml:"Id"`
  ReleaseAreaId int `yaml:"ReleaseAreaId"`
  ConditionsType int `yaml:"ConditionsType"`
  ConditionsValue int `yaml:"ConditionsValue"`
  StartTime time.Time `yaml:"StartTime"`
  EndTime time.Time `yaml:"EndTime"`
}

type QuestLiveDownloads struct {
  Id int64 `yaml:"Id"`
  Title string `yaml:"Title"`
  OrderId int `yaml:"OrderId"`
  StartTime time.Time `yaml:"StartTime"`
  EndTime time.Time `yaml:"EndTime"`
}

type QuestLiveLoadings struct {
  Id int64 `yaml:"Id"`
  Title string `yaml:"Title"`
  OrderId int `yaml:"OrderId"`
  StartTime time.Time `yaml:"StartTime"`
  EndTime time.Time `yaml:"EndTime"`
}

type QuestLiveReleaseConditions struct {
  Id int64 `yaml:"Id"`
  ReleaseQuestId int `yaml:"ReleaseQuestId"`
  ConditionsType int `yaml:"ConditionsType"`
  ConditionsValue int `yaml:"ConditionsValue"`
  StartTime time.Time `yaml:"StartTime"`
  EndTime time.Time `yaml:"EndTime"`
}

type QuestSections struct {
  Id int64 `yaml:"Id"`
  SectionNo int `yaml:"SectionNo"`
  QuestStagesId int `yaml:"QuestStagesId"`
  SectionSkillsId int64 `yaml:"SectionSkillsId"`
}

type RaidEvents struct {
  Id int `yaml:"Id"`
  Name string `yaml:"Name"`
  Description string `yaml:"Description"`
  DropNumUp int `yaml:"DropNumUp"`
  PersonalRankingTabs string `yaml:"PersonalRankingTabs"`
  InfoStartTime time.Time `yaml:"InfoStartTime"`
  InfoEndTime time.Time `yaml:"InfoEndTime"`
  StartTime time.Time `yaml:"StartTime"`
  EndTime time.Time `yaml:"EndTime"`
}

type RaidQuestDropRateUp struct {
  Id int `yaml:"Id"`
  Name string `yaml:"Name"`
  Flame int `yaml:"Flame"`
}

type RaidQuestReleaseCondition struct {
  Id int64 `yaml:"Id"`
  ReleaseRaidId int `yaml:"ReleaseRaidId"`
  ConditionsType int `yaml:"ConditionsType"`
  ConditionsValue int `yaml:"ConditionsValue"`
  StartTime time.Time `yaml:"StartTime"`
  EndTime time.Time `yaml:"EndTime"`
}

type RaidQuestSeries struct {
  Id int `yaml:"Id"`
  Name string `yaml:"Name"`
  Description string `yaml:"Description"`
  RaidEventId int `yaml:"RaidEventId"`
  PointAddLimit int64 `yaml:"PointAddLimit"`
  StartTime time.Time `yaml:"StartTime"`
  EndTime time.Time `yaml:"EndTime"`
}

type RaidQuestStages struct {
  Id int `yaml:"Id"`
  RaidQuestSeriesId int `yaml:"RaidQuestSeriesId"`
  Name string `yaml:"Name"`
  Description string `yaml:"Description"`
  Hint string `yaml:"Hint"`
  OrderId int `yaml:"OrderId"`
  StageType int `yaml:"StageType"`
  LiveStagesId int `yaml:"LiveStagesId"`
  QuestMusicsType int `yaml:"QuestMusicsType"`
  QuestMusicsDetail int `yaml:"QuestMusicsDetail"`
  DeckRestrictedType int `yaml:"DeckRestrictedType"`
  DeckRestrictedDetail int `yaml:"DeckRestrictedDetail"`
  QuestLevel int `yaml:"QuestLevel"`
  QuestRank int `yaml:"QuestRank"`
  FirstClearRewardSeriesId int64 `yaml:"FirstClearRewardSeriesId"`
  CompleteRewardSeriesId int64 `yaml:"CompleteRewardSeriesId"`
  DropRewardSeriesId int64 `yaml:"DropRewardSeriesId"`
  RandomDropRewardSeriesId int64 `yaml:"RandomDropRewardSeriesId"`
  Score1 int64 `yaml:"Score1"`
  Score2 int64 `yaml:"Score2"`
  Score3 int64 `yaml:"Score3"`
  StylePoint int `yaml:"StylePoint"`
  GainMusicExp int `yaml:"GainMusicExp"`
}

type RaidResourceAddDate struct {
  Id int `yaml:"Id"`
  RaidResourceId int `yaml:"RaidResourceId"`
  AddNum int `yaml:"AddNum"`
  AddTime time.Time `yaml:"AddTime"`
}

type RaidResource struct {
  Id int `yaml:"Id"`
  RaidEventId int `yaml:"RaidEventId"`
  RaidResourceDefaultNum int `yaml:"RaidResourceDefaultNum"`
  RaidResourceLimit int `yaml:"RaidResourceLimit"`
  RaidResourceAddLimit int `yaml:"RaidResourceAddLimit"`
}

type RaidResourceRecoveryDatas struct {
  Id int `yaml:"Id"`
  RaidResourceId int `yaml:"RaidResourceId"`
  Order int `yaml:"Order"`
  RequireItemType int `yaml:"RequireItemType"`
  RequireItemId int `yaml:"RequireItemId"`
  RequireItemNum int `yaml:"RequireItemNum"`
}

type RaidRewardDatas struct {
  Id int64 `yaml:"Id"`
  RaidRewardsId int `yaml:"RaidRewardsId"`
  RewardType int `yaml:"RewardType"`
  RewardItemId int `yaml:"RewardItemId"`
  RewardNum int `yaml:"RewardNum"`
}

type RaidRewards struct {
  Id int `yaml:"Id"`
  RaidEventId int `yaml:"RaidEventId"`
  RewardType int `yaml:"RewardType"`
  RequirePointAmount int64 `yaml:"RequirePointAmount"`
}

type RaidTopProgressImage struct {
  Id int `yaml:"Id"`
  RaidEventId int `yaml:"RaidEventId"`
  Order int `yaml:"Order"`
  RequirePoint int64 `yaml:"RequirePoint"`
}

type RentalCardDatas struct {
  Id int64 `yaml:"Id"`
  CardDatasId int `yaml:"CardDatasId"`
  StyleLevel int `yaml:"StyleLevel"`
  LimitBreakLevel int `yaml:"LimitBreakLevel"`
  SpecialAppealLevel int `yaml:"SpecialAppealLevel"`
  SkillLevel int `yaml:"SkillLevel"`
}

type RentalDeckCards struct {
  Id int `yaml:"Id"`
  RentalDecksId int `yaml:"RentalDecksId"`
  CharactersId int `yaml:"CharactersId"`
  RentalCardId_Main int64 `yaml:"RentalCardId_Main"`
  RentalCardId_Side1 int64 `yaml:"RentalCardId_Side1"`
  RentalCardId_Side2 int64 `yaml:"RentalCardId_Side2"`
}

type RentalDecks struct {
  Id int `yaml:"Id"`
  DeckName string `yaml:"DeckName"`
  GenerationsId int `yaml:"GenerationsId"`
  DeckNumber int `yaml:"DeckNumber"`
  AceCardId int64 `yaml:"AceCardId"`
  ReleaseConditionsType int `yaml:"ReleaseConditionsType"`
  ConditionsValue int `yaml:"ConditionsValue"`
  ReleaseConditonsDescription string `yaml:"ReleaseConditonsDescription"`
  ReleaseDeckStartTime time.Time `yaml:"ReleaseDeckStartTime"`
  ReleaseDeckEndTime time.Time `yaml:"ReleaseDeckEndTime"`
}

type SeasonFanLevels struct {
  Id int `yaml:"Id"`
  SeasonsId int `yaml:"SeasonsId"`
  SeasonFanLevel int `yaml:"SeasonFanLevel"`
  Experience int `yaml:"Experience"`
  CumulativeExperience int `yaml:"CumulativeExperience"`
}

type SeasonGrade struct {
  Id int `yaml:"Id"`
  Description string `yaml:"Description"`
  PersonalRankingTabs string `yaml:"PersonalRankingTabs"`
  CountStartTime time.Time `yaml:"CountStartTime"`
  CountEndTime time.Time `yaml:"CountEndTime"`
  DisplayStartTime time.Time `yaml:"DisplayStartTime"`
  DisplayEndTime time.Time `yaml:"DisplayEndTime"`
  TermTitle string `yaml:"TermTitle"`
}

type SeasonGradeRewardDatas struct {
  Id int64 `yaml:"Id"`
  SeasonGradeRewardsId int `yaml:"SeasonGradeRewardsId"`
  RewardType int `yaml:"RewardType"`
  RewardItemId int `yaml:"RewardItemId"`
  RewardNum int `yaml:"RewardNum"`
  IsEmphasize int `yaml:"IsEmphasize"`
}

type SeasonGradeRewards struct {
  Id int `yaml:"Id"`
  SeasonGradeId int `yaml:"SeasonGradeId"`
  MinTargetNum int `yaml:"MinTargetNum"`
  MaxTargetNum int `yaml:"MaxTargetNum"`
}

type Seasons struct {
  Id int `yaml:"Id"`
  Name string `yaml:"Name"`
  StartTime time.Time `yaml:"StartTime"`
  EndTime time.Time `yaml:"EndTime"`
}

type SectionSkillEffectDetails struct {
  Id int64 `yaml:"Id"`
  SkillEffectDetailType string `yaml:"SkillEffectDetailType"`
  TargetMood int `yaml:"TargetMood"`
  EffectValue int `yaml:"EffectValue"`
}

type SectionSkillEffects struct {
  Id int64 `yaml:"Id"`
  ActionType int `yaml:"ActionType"`
  OrderId int `yaml:"OrderId"`
}

type SectionSkills struct {
  Id int64 `yaml:"Id"`
  Description string `yaml:"Description"`
  ApperanceType int `yaml:"ApperanceType"`
  SkillIcon int `yaml:"SkillIcon"`
  SectionSkillsEffectId string `yaml:"SectionSkillsEffectId"`
}

type SelectTicketExchangeRate struct {
  Id int64 `yaml:"Id"`
  SelectTicketSeriesId int `yaml:"SelectTicketSeriesId"`
  ExchangeItemType int `yaml:"ExchangeItemType"`
  ExchangeItemId int `yaml:"ExchangeItemId"`
  ExchangeItemQuantity int `yaml:"ExchangeItemQuantity"`
}

type SelectTicketSeries struct {
  Id int `yaml:"Id"`
  ExchangeTicketName string `yaml:"ExchangeTicketName"`
  Description string `yaml:"Description"`
  ExchangeTicketId int `yaml:"ExchangeTicketId"`
  OrderId int `yaml:"OrderId"`
  PickUpCardSeriesId_1 int `yaml:"PickUpCardSeriesId_1"`
  PickUpCardSeriesId_2 int `yaml:"PickUpCardSeriesId_2"`
  PickUpCardSeriesId_3 int `yaml:"PickUpCardSeriesId_3"`
  PickUpCardSeriesId_4 int `yaml:"PickUpCardSeriesId_4"`
  PickUpCardSeriesId_5 int `yaml:"PickUpCardSeriesId_5"`
  PickUpCardSeriesId_6 int `yaml:"PickUpCardSeriesId_6"`
  BgType int `yaml:"BgType"`
  StartTime time.Time `yaml:"StartTime"`
  EndTime time.Time `yaml:"EndTime"`
  IsVisibleEndTime int `yaml:"IsVisibleEndTime"`
}

type ShopItems struct {
  Id int `yaml:"Id"`
  ShopId int `yaml:"ShopId"`
  Name string `yaml:"Name"`
  ItemType int `yaml:"ItemType"`
  ItemId int `yaml:"ItemId"`
  ItemQuantity int `yaml:"ItemQuantity"`
  Price int `yaml:"Price"`
  IsPaidSIsCaOnly int `yaml:"IsPaidSIsCaOnly"`
  OrderId int `yaml:"OrderId"`
  Description string `yaml:"Description"`
  RewardTextId int `yaml:"RewardTextId"`
  StartTime time.Time `yaml:"StartTime"`
  EndTime time.Time `yaml:"EndTime"`
}

type Shops struct {
  Id int `yaml:"Id"`
  ShopType int `yaml:"ShopType"`
  Name string `yaml:"Name"`
  OrderId int `yaml:"OrderId"`
}

type SimulationGraphLimit struct {
  Id int `yaml:"Id"`
  NumberOfMember int `yaml:"NumberOfMember"`
  UpperLimitSmile int `yaml:"UpperLimitSmile"`
  UpperLimitPure int `yaml:"UpperLimitPure"`
  UpperLimitCool int `yaml:"UpperLimitCool"`
  UpperLimitMental int `yaml:"UpperLimitMental"`
  UpperLimitBP int `yaml:"UpperLimitBP"`
}

type StageSkillConditionDetails struct {
  Id int64 `yaml:"Id"`
  StageSkillConditionId int64 `yaml:"StageSkillConditionId"`
  SkillConditionDetailType string `yaml:"SkillConditionDetailType"`
  ConditionValue int64 `yaml:"ConditionValue"`
}

type StageSkillConditions struct {
  Id int64 `yaml:"Id"`
  ConditionType int `yaml:"ConditionType"`
}

type StageSkillEffectDetails struct {
  Id int64 `yaml:"Id"`
  StageSkillEffectId int64 `yaml:"StageSkillEffectId"`
  SkillEffectDetailType string `yaml:"SkillEffectDetailType"`
  EffectValue int64 `yaml:"EffectValue"`
}

type StageSkillEffects struct {
  Id int64 `yaml:"Id"`
  ActionType int `yaml:"ActionType"`
}

type StageSkillSets struct {
  Id int64 `yaml:"Id"`
  StageSkillConditionId int64 `yaml:"StageSkillConditionId"`
  StageSkillEffectId int64 `yaml:"StageSkillEffectId"`
}

type Stamps struct {
  Id int `yaml:"Id"`
  Name string `yaml:"Name"`
  StampNo int `yaml:"StampNo"`
  StampType int `yaml:"StampType"`
  CharactersId int `yaml:"CharactersId"`
  StartTime time.Time `yaml:"StartTime"`
  EndTime time.Time `yaml:"EndTime"`
}

type StandardQuestAreas struct {
  Id int `yaml:"Id"`
  Name string `yaml:"Name"`
  Number int `yaml:"Number"`
  Description string `yaml:"Description"`
  OrderId int `yaml:"OrderId"`
  ImageType int `yaml:"ImageType"`
  BgImageId int `yaml:"BgImageId"`
  SoundId int `yaml:"SoundId"`
  GenerationsId int `yaml:"GenerationsId"`
  StartTime time.Time `yaml:"StartTime"`
  EndTime time.Time `yaml:"EndTime"`
}

type StandardQuestStages struct {
  Id int `yaml:"Id"`
  StandardQuestAreasId int `yaml:"StandardQuestAreasId"`
  Name string `yaml:"Name"`
  Description string `yaml:"Description"`
  Hint string `yaml:"Hint"`
  StageType int `yaml:"StageType"`
  UseType int `yaml:"UseType"`
  UseItem int `yaml:"UseItem"`
  UseNum int `yaml:"UseNum"`
  LiveStagesId int `yaml:"LiveStagesId"`
  QuestMusicsType int `yaml:"QuestMusicsType"`
  QuestMusicsDetail int `yaml:"QuestMusicsDetail"`
  DeckRestrictedType int `yaml:"DeckRestrictedType"`
  DeckRestrictedDetail int `yaml:"DeckRestrictedDetail"`
  QuestLevel int `yaml:"QuestLevel"`
  FirstClearRewardSeriesId int64 `yaml:"FirstClearRewardSeriesId"`
  CompleteRewardSeriesId int64 `yaml:"CompleteRewardSeriesId"`
  DropRewardSeriesId int64 `yaml:"DropRewardSeriesId"`
  RandomDropRewardSeriesId int64 `yaml:"RandomDropRewardSeriesId"`
  Score1 int64 `yaml:"Score1"`
  Score2 int64 `yaml:"Score2"`
  Score3 int64 `yaml:"Score3"`
  GainStylePoint int `yaml:"GainStylePoint"`
  GainMusicExp int `yaml:"GainMusicExp"`
  ItemSourceIconId int `yaml:"ItemSourceIconId"`
}

type Stickers struct {
  Id int `yaml:"Id"`
  Name string `yaml:"Name"`
  Text string `yaml:"Text"`
  CategoryType int `yaml:"CategoryType"`
  CategoryName int `yaml:"CategoryName"`
  SeasonId int `yaml:"SeasonId"`
  CharactersId int `yaml:"CharactersId"`
  Priority int `yaml:"Priority"`
  IsVariant int `yaml:"IsVariant"`
  RequirementType int `yaml:"RequirementType"`
  RequirementDetail string `yaml:"RequirementDetail"`
  RequirementValue int `yaml:"RequirementValue"`
  RequirementText string `yaml:"RequirementText"`
  EditRequirementType int `yaml:"EditRequirementType"`
  EditRequirementDetail string `yaml:"EditRequirementDetail"`
  EditRequirementValue int `yaml:"EditRequirementValue"`
  EditRequirementText string `yaml:"EditRequirementText"`
  IsVisibleOnlyPossess int `yaml:"IsVisibleOnlyPossess"`
  StartTime time.Time `yaml:"StartTime"`
  EndTime time.Time `yaml:"EndTime"`
  AvailableStartTime time.Time `yaml:"AvailableStartTime"`
  AvailableEndTime time.Time `yaml:"AvailableEndTime"`
  VariantStartTime time.Time `yaml:"VariantStartTime"`
  VariantEndTime time.Time `yaml:"VariantEndTime"`
}

type StyleMovies struct {
  Id int `yaml:"Id"`
  CardSeriesId int `yaml:"CardSeriesId"`
  MovieType int `yaml:"MovieType"`
  Name string `yaml:"Name"`
  ReleaseConditionText string `yaml:"ReleaseConditionText"`
}

type StyleVoices struct {
  Id int `yaml:"Id"`
  CardSeriesId int `yaml:"CardSeriesId"`
  Name string `yaml:"Name"`
  Priority int `yaml:"Priority"`
  VoiceName string `yaml:"VoiceName"`
  ReleaseConditionText string `yaml:"ReleaseConditionText"`
}

type SubCharacters struct {
  Id int `yaml:"Id"`
  Label string `yaml:"Label"`
}

type TabList struct {
  Id int `yaml:"Id"`
  TabListName string `yaml:"TabListName"`
}

type TextsPlaceHolder struct {
  Id int `yaml:"Id"`
  Description string `yaml:"Description"`
}

type TicketOnlyGachas struct {
  Id int `yaml:"Id"`
  GachaTicketId int `yaml:"GachaTicketId"`
  TicketNum int `yaml:"TicketNum"`
  ConsectiveGachaTimes int `yaml:"ConsectiveGachaTimes"`
  GachaButtonType int `yaml:"GachaButtonType"`
  ButtonPopId int `yaml:"ButtonPopId"`
  MiniBannerPopId int `yaml:"MiniBannerPopId"`
  MiniBannerDispCondition int `yaml:"MiniBannerDispCondition"`
}

type TutorialDeckCards struct {
  Id int `yaml:"Id"`
  TutorialDeckDatasId int `yaml:"TutorialDeckDatasId"`
  SlotNo int `yaml:"SlotNo"`
  CardDatasId int `yaml:"CardDatasId"`
  StyleLevel int `yaml:"StyleLevel"`
  LimitBreakTimes int `yaml:"LimitBreakTimes"`
  SpecialAppealLevel int `yaml:"SpecialAppealLevel"`
  SkillLevel int `yaml:"SkillLevel"`
}

type TutorialDeckDatas struct {
  Id int `yaml:"Id"`
  DeckName string `yaml:"DeckName"`
  DeckNo int `yaml:"DeckNo"`
  GenerationsId int `yaml:"GenerationsId"`
}

type TutorialQuestAreas struct {
  Id int `yaml:"Id"`
  Name string `yaml:"Name"`
  Number int `yaml:"Number"`
  Description string `yaml:"Description"`
  OrderId int `yaml:"OrderId"`
  ImageType int `yaml:"ImageType"`
  BgImageId int `yaml:"BgImageId"`
  Next_TutorialQuestAreasId int `yaml:"Next_TutorialQuestAreasId"`
  GenerationsId int `yaml:"GenerationsId"`
}

type TutorialQuestStages struct {
  Id int `yaml:"Id"`
  TutorialQuestAreasId int `yaml:"TutorialQuestAreasId"`
  Name string `yaml:"Name"`
  Description string `yaml:"Description"`
  Hint string `yaml:"Hint"`
  MapNumber int `yaml:"MapNumber"`
  StageType int `yaml:"StageType"`
  UseType int `yaml:"UseType"`
  Use_ItemsId int `yaml:"Use_ItemsId"`
  UseNum int `yaml:"UseNum"`
  LiveStagesId int `yaml:"LiveStagesId"`
  QuestMusicsId int `yaml:"QuestMusicsId"`
  LoveCorrectionValue int `yaml:"LoveCorrectionValue"`
  BonusVoltage int `yaml:"BonusVoltage"`
  BonusMental int `yaml:"BonusMental"`
  BonusHeart int `yaml:"BonusHeart"`
  BonusLove int `yaml:"BonusLove"`
  FirstClearRewardSeriesId int64 `yaml:"FirstClearRewardSeriesId"`
  CompleteRewardSeriesId int64 `yaml:"CompleteRewardSeriesId"`
  RandomDropRewardSeriesId int64 `yaml:"RandomDropRewardSeriesId"`
  Score1 int64 `yaml:"Score1"`
  Score2 int64 `yaml:"Score2"`
  Score3 int64 `yaml:"Score3"`
  StylePoint int `yaml:"StylePoint"`
  GainMusicExp int `yaml:"GainMusicExp"`
}

type TutorialRewardDatas struct {
  Id int `yaml:"Id"`
  TutorialsId int `yaml:"TutorialsId"`
  RewardType int `yaml:"RewardType"`
  RewardItemId int `yaml:"RewardItemId"`
  RewardNum int `yaml:"RewardNum"`
  LifeTimeDay int `yaml:"LifeTimeDay"`
  RewardTextId int `yaml:"RewardTextId"`
}

type TutorialSchoolIdolStageMovies struct {
  Id int `yaml:"Id"`
  Title string `yaml:"Title"`
  OrderId int `yaml:"OrderId"`
}

type Tutorials struct {
  Id int `yaml:"Id"`
  TutorialType int `yaml:"TutorialType"`
  Step int `yaml:"Step"`
  Description string `yaml:"Description"`
}

type UnitCharacters struct {
  Id int `yaml:"Id"`
  UnitsId int `yaml:"UnitsId"`
  CharactersId int `yaml:"CharactersId"`
  OrderId int `yaml:"OrderId"`
}

type Units struct {
  Id int `yaml:"Id"`
  UnitName string `yaml:"UnitName"`
  OrderId int `yaml:"OrderId"`
  StartTime time.Time `yaml:"StartTime"`
  EndTime time.Time `yaml:"EndTime"`
}

type LiveCharacters struct {
  Id int `yaml:"Id"`
  Label string `yaml:"Label"`
  SkeletonName string `yaml:"SkeletonName"`
  ItemsIds string `yaml:"ItemsIds"`
  PosesIds string `yaml:"PosesIds"`
}

type LiveEventsEvol struct {
  Id int `yaml:"Id"`
  Label string `yaml:"Label"`
  ContentsType int `yaml:"ContentsType"`
  LocationsId int `yaml:"LocationsId"`
  CharactersIds string `yaml:"CharactersIds"`
  CostumesIds string `yaml:"CostumesIds"`
  TimelinesIds string `yaml:"TimelinesIds"`
}

type LiveItems struct {
  Id int `yaml:"Id"`
  Label string `yaml:"Label"`
  BindBoneId int `yaml:"BindBoneId"`
  PosesId int `yaml:"PosesId"`
}

type LiveLocations struct {
  Id int `yaml:"Id"`
  Label string `yaml:"Label"`
  PropsIds string `yaml:"PropsIds"`
}

type LiveMovies struct {
  Id int `yaml:"Id"`
  Label string `yaml:"Label"`
}

type LiveMusic struct {
  Id int `yaml:"Id"`
  Label string `yaml:"Label"`
  MusicId int `yaml:"MusicId"`
  HaveMusic int `yaml:"HaveMusic"`
  HaveMotion int `yaml:"HaveMotion"`
  CharactersCount int `yaml:"CharactersCount"`
  CharactersIds string `yaml:"CharactersIds"`
}

type LivePoses struct {
  Id int `yaml:"Id"`
  Label string `yaml:"Label"`
  HandSide int `yaml:"HandSide"`
}

type LiveProps struct {
  Id int `yaml:"Id"`
  Label string `yaml:"Label"`
}

type LiveTimelinesEvol struct {
  Id int64 `yaml:"Id"`
  Label string `yaml:"Label"`
  MusicId int `yaml:"MusicId"`
  LocationsId int `yaml:"LocationsId"`
  FreeId int `yaml:"FreeId"`
  NextId int64 `yaml:"NextId"`
  MovieIds string `yaml:"MovieIds"`
}

