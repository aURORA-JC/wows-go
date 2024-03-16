package api

import (
	"github.com/aURORA-JC/wows-go/pkg/client"
	"github.com/aURORA-JC/wows-go/pkg/models"
	"github.com/aURORA-JC/wows-go/pkg/utils"
	"net/url"
	"strconv"
)

const EncyclopediaPrefix = "/encyclopedia"

type EncyclopediaService struct {
	Client *client.WowsClient
}

func NewEncyclopediaService(c *client.WowsClient) *EncyclopediaService {
	return &EncyclopediaService{Client: c}
}

func (s *EncyclopediaService) GetInfo() (*models.EncyclopediaInfoResponse, error) {
	target := EncyclopediaPrefix + "/info/"

	info := &models.EncyclopediaInfoResponse{}
	err := s.Client.Do(target, &url.Values{}, info)
	if err != nil {
		return nil, err
	}

	return info, nil
}

func (s *EncyclopediaService) GetShips(shipsIds []int64, shipsNations, shipsTypes []string, pagination bool, limit, page int) (*models.EncyclopediaShipsResponse, error) {
	target := EncyclopediaPrefix + "/ships/"

	formData := &url.Values{}
	if shipsIds != nil && len(shipsIds) != 0 {
		formData.Add("ship_id", utils.ConvertArray2Str(shipsIds))
	}

	if shipsNations != nil && len(shipsNations) != 0 {
		formData.Add("nation", utils.ConvertArray2Str(shipsNations))
	}

	if shipsTypes != nil && len(shipsTypes) != 0 {
		formData.Add("type", utils.ConvertArray2Str(shipsTypes))
	}

	if pagination {
		utils.SetPagination(formData, limit, page)
	}

	ships := &models.EncyclopediaShipsResponse{}
	err := s.Client.Do(target, formData, ships)
	if err != nil {
		return nil, err
	}

	return ships, nil
}

func (s *EncyclopediaService) GetAchievements() (*models.EncyclopediaAchievementsResponse, error) {
	target := EncyclopediaPrefix + "/achievements/"

	achievements := &models.EncyclopediaAchievementsResponse{}
	err := s.Client.Do(target, &url.Values{}, achievements)
	if err != nil {
		return nil, err
	}

	return achievements, nil
}

// GetShipProfile parameters of ships in all existing configurations.
//
// The modulesIds array should contain 10 elements, arranged in the following sequence: artillery_id, dive_bomber_id,
// engine_id, fighter_id, fire_control_id, flight_control_id, hull_id, submarine_sonar_id, torpedo_bomber_id,
// torpedoes_id. For instance, if only the artillery and hull modules need to be specified, the array could be
// [0, 0, 3557666608, 0, 0, 0, 3558616880]. The length of the input array is determined by the last module that is set.
func (s *EncyclopediaService) GetShipProfile(shipId int64, modulesIds []int64) (*models.EncyclopediaShipProfileResponse, error) {
	target := EncyclopediaPrefix + "/shipprofile/"

	formData := &url.Values{}
	if shipId != 0 {
		formData.Add("ship_id", strconv.FormatInt(shipId, 10))
	}

	if modulesIds != nil && len(modulesIds) != 0 {
		tmp := make([]int64, 10)
		copy(tmp, modulesIds)
		formData.Add("artillery_id", strconv.FormatInt(tmp[0], 10))
		formData.Add("dive_bomber_id", strconv.FormatInt(tmp[1], 10))
		formData.Add("engine_id", strconv.FormatInt(tmp[2], 10))
		formData.Add("fighter_id", strconv.FormatInt(tmp[3], 10))
		formData.Add("fire_control_id", strconv.FormatInt(tmp[4], 10))
		formData.Add("flight_control_id", strconv.FormatInt(tmp[5], 10))
		formData.Add("hull_id", strconv.FormatInt(tmp[6], 10))
		formData.Add("submarine_sonar_id", strconv.FormatInt(tmp[7], 10))
		formData.Add("torpedo_bomber_id", strconv.FormatInt(tmp[8], 10))
		formData.Add("torpedoes_id", strconv.FormatInt(tmp[9], 10))
	}

	profile := &models.EncyclopediaShipProfileResponse{}
	err := s.Client.Do(target, formData, profile)
	if err != nil {
		return nil, err
	}

	return profile, nil
}

func (s *EncyclopediaService) GetModules(moduleIds []int64, moduleType string, pagination bool, limit, page int) (*models.EncyclopediaModulesResponse, error) {
	target := EncyclopediaPrefix + "/modules/"

	formData := &url.Values{}
	if moduleIds != nil && len(moduleIds) != 0 {
		formData.Add("commander_id", utils.ConvertArray2Str(moduleIds))
	}

	if moduleType != "" {
		formData.Add("type", moduleType)
	}

	if pagination {
		utils.SetPagination(formData, limit, page)
	}

	modules := &models.EncyclopediaModulesResponse{}
	err := s.Client.Do(target, formData, modules)
	if err != nil {
		return nil, err
	}

	return modules, nil
}

func (s *EncyclopediaService) GetAccountLevels() (*models.EncyclopediaAccountLevelsResponse, error) {
	target := EncyclopediaPrefix + "/accountlevels/"

	accountLevels := &models.EncyclopediaAccountLevelsResponse{}
	err := s.Client.Do(target, &url.Values{}, accountLevels)
	if err != nil {
		return nil, err
	}

	return accountLevels, nil
}

func (s *EncyclopediaService) GetCrews(crewsIds []int64) (*models.EncyclopediaCrewsResponse, error) {
	target := EncyclopediaPrefix + "/crews/"

	formData := &url.Values{}
	if crewsIds != nil && len(crewsIds) != 0 {
		formData.Add("commander_id", utils.ConvertArray2Str(crewsIds))
	}

	crews := &models.EncyclopediaCrewsResponse{}
	err := s.Client.Do(target, formData, crews)
	if err != nil {
		return nil, err
	}

	return crews, nil
}

func (s *EncyclopediaService) GetCrewSkills(skillsIds []int64) (*models.EncyclopediaCrewSkillsResponse, error) {
	target := EncyclopediaPrefix + "/crewskills/"

	formData := &url.Values{}
	if skillsIds != nil && len(skillsIds) != 0 {
		formData.Add("skill_id", utils.ConvertArray2Str(skillsIds))
	}

	skills := &models.EncyclopediaCrewSkillsResponse{}
	err := s.Client.Do(target, formData, skills)
	if err != nil {
		return nil, err
	}

	return skills, nil
}

func (s *EncyclopediaService) GetCrewRanks(nation string) (*models.EncyclopediaCrewRanksResponse, error) {
	target := EncyclopediaPrefix + "/crewranks/"

	formData := &url.Values{}
	if nation != "" {
		formData.Add("nation", nation)
	}

	ranks := &models.EncyclopediaCrewRanksResponse{}
	err := s.Client.Do(target, formData, ranks)
	if err != nil {
		return nil, err
	}

	return ranks, nil
}

func (s *EncyclopediaService) GetBattleTypes() (*models.EncyclopediaBattleTypesResponse, error) {
	target := EncyclopediaPrefix + "/battletypes/"

	types := &models.EncyclopediaBattleTypesResponse{}
	err := s.Client.Do(target, &url.Values{}, types)
	if err != nil {
		return nil, err
	}

	return types, nil
}

func (s *EncyclopediaService) GetConsumables(consumablesIds []int64, consumablesType string, pagination bool, limit, page int) (*models.EncyclopediaConsumablesResponse, error) {
	target := EncyclopediaPrefix + "/consumables/"

	formData := &url.Values{}
	if consumablesIds != nil && len(consumablesIds) != 0 {
		formData.Add("consumable_id", utils.ConvertArray2Str(consumablesIds))
	}

	if consumablesType != "" {
		formData.Add("type", consumablesType)
	}

	if pagination {
		utils.SetPagination(formData, limit, page)
	}

	consumables := &models.EncyclopediaConsumablesResponse{}
	err := s.Client.Do(target, formData, consumables)
	if err != nil {
		return nil, err
	}

	return consumables, nil
}

func (s *EncyclopediaService) GetCollections() (*models.EncyclopediaCollectionsResponse, error) {
	target := EncyclopediaPrefix + "/collections/"

	collections := &models.EncyclopediaCollectionsResponse{}
	err := s.Client.Do(target, &url.Values{}, collections)
	if err != nil {
		return nil, err
	}

	return collections, nil
}

func (s *EncyclopediaService) GetCollectionCards() (*models.EncyclopediaCollectionCardsResponse, error) {
	target := EncyclopediaPrefix + "/collectioncards/"

	cards := &models.EncyclopediaCollectionCardsResponse{}
	err := s.Client.Do(target, &url.Values{}, cards)
	if err != nil {
		return nil, err
	}

	return cards, nil
}

func (s *EncyclopediaService) GetBattleArenas() (*models.EncyclopediaBattleArenasResponse, error) {
	target := EncyclopediaPrefix + "/battlearenas/"

	arenas := &models.EncyclopediaBattleArenasResponse{}
	err := s.Client.Do(target, &url.Values{}, arenas)
	if err != nil {
		return nil, err
	}

	return arenas, nil
}
