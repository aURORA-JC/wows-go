package test

import (
	"fmt"
	"github.com/aURORA-JC/wows-go/pkg/api"
	"github.com/aURORA-JC/wows-go/pkg/client"
	"github.com/aURORA-JC/wows-go/pkg/constants"
	"testing"
)

func TestEncyclopediaService_GetInfo(t *testing.T) {
	c := client.NewAsiaClient(getKey())
	s := api.NewEncyclopediaService(c)

	info, err := s.GetInfo()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	fmt.Printf("%v", info)
}

func TestEncyclopediaService_GetShips(t *testing.T) {
	c := client.NewAsiaClient(getKey())
	s := api.NewEncyclopediaService(c)

	type args struct {
		ids        []int64
		nations    []string
		types      []string
		pagination bool
		limit      int
		page       int
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test1: Default",
			args: args{},
		},
		{
			name: "Test2: Nation",
			args: args{
				nations: []string{"usa", "uk"},
			},
		},
		{
			name: "Test3: ShipId",
			args: args{
				ids: []int64{4181669680, 4276041712},
			},
		},
		{
			name: "Test4: ShipType",
			args: args{
				types: []string{constants.WowsShipTypeBB, constants.WowsShipTypeDD},
			},
		},
		{
			name: "Test5: Pagination",
			args: args{
				pagination: true,
				limit:      2,
				page:       10,
			},
		},
		{
			name: "Test6: More",
			args: args{
				nations:    []string{"usa", "uk"},
				types:      []string{constants.WowsShipTypeBB, constants.WowsShipTypeDD},
				pagination: true,
				limit:      2,
				page:       10,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ships, err := s.GetShips(tt.args.ids, tt.args.nations, tt.args.types, tt.args.pagination, tt.args.limit, tt.args.page)
			if err != nil {
				t.Error(err)
				t.FailNow()
			}

			fmt.Printf("%v\n", ships)
		})
	}

}

func TestEncyclopediaService_GetAchievements(t *testing.T) {
	c := client.NewAsiaClient(getKey())
	s := api.NewEncyclopediaService(c)

	info, err := s.GetAchievements()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	fmt.Printf("%v", info)
}

func TestEncyclopediaService_GetShipProfile(t *testing.T) {
	c := client.NewAsiaClient(getKey())
	s := api.NewEncyclopediaService(c)

	type args struct {
		shipId     int64
		modulesIds []int64
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test1: Default",
			args: args{
				shipId: 4181669680,
			},
		},
		{
			name: "Test2: ModulesIds",
			args: args{
				shipId:     4181669680,
				modulesIds: []int64{0, 0, 3557666608, 0, 0, 0, 3558616880},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			profile, err := s.GetShipProfile(tt.args.shipId, tt.args.modulesIds)
			if err != nil {
				t.Error(err)
				t.FailNow()
			}

			fmt.Printf("%v\n", profile)
		})
	}
}

func TestEncyclopediaService_GetModules(t *testing.T) {
	c := client.NewAsiaClient(getKey())
	s := api.NewEncyclopediaService(c)

	type args struct {
		modulesId   []int64
		modulesType string
		pagination  bool
		limit       int
		page        int
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test1: Default",
			args: args{},
		},
		{
			name: "Test2: ModuleId",
			args: args{
				modulesId: []int64{3268161232, 3343658992},
			},
		},
		{
			name: "Test3: ModuleType",
			args: args{
				modulesType: constants.WowsModuleTypeMainBattery,
			},
		},
		{
			name: "Test4: Pagination",
			args: args{
				pagination: true,
				limit:      10,
				page:       2,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			modules, err := s.GetModules(tt.args.modulesId, tt.args.modulesType,
				tt.args.pagination, tt.args.limit, tt.args.page)
			if err != nil {
				t.Error(err)
				t.FailNow()
			}

			fmt.Printf("%v\n", modules)
		})
	}
}

func TestEncyclopediaService_GetAccountLevels(t *testing.T) {
	c := client.NewAsiaClient(getKey())
	s := api.NewEncyclopediaService(c)

	levels, err := s.GetAccountLevels()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	fmt.Printf("%v\n", levels)
}

func TestEncyclopediaService_GetCrews(t *testing.T) {
	c := client.NewAsiaClient(getKey())
	s := api.NewEncyclopediaService(c)

	type args struct {
		crewsIds []int64
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test1: Default",
			args: args{},
		},
		{
			name: "Test2: CrewsIds",
			args: args{
				crewsIds: []int64{3557992400, 3810698960},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			crews, err := s.GetCrews(tt.args.crewsIds)
			if err != nil {
				t.Error(err)
				t.FailNow()
			}

			fmt.Printf("%v\n", crews)
		})
	}
}

func TestEncyclopediaService_GetCrewSkills(t *testing.T) {
	c := client.NewAsiaClient(getKey())
	s := api.NewEncyclopediaService(c)

	type args struct {
		skillIds []int64
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test1: Default",
			args: args{},
		},
		{
			name: "Test2: SkillsIds",
			args: args{
				skillIds: []int64{1, 9},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			skills, err := s.GetCrewSkills(tt.args.skillIds)
			if err != nil {
				t.Error(err)
				t.FailNow()
			}

			fmt.Printf("%v\n", skills)
		})
	}
}

func TestEncyclopediaService_GetCrewRanks(t *testing.T) {
	c := client.NewAsiaClient(getKey())
	s := api.NewEncyclopediaService(c)

	type args struct {
		nation string
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test1: Default",
			args: args{},
		},
		{
			name: "Test2: Nation",
			args: args{
				nation: constants.WowsNationPanAmerica,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ranks, err := s.GetCrewRanks(tt.args.nation)
			if err != nil {
				t.Error(err)
				t.FailNow()
			}

			fmt.Printf("%v\n", ranks)
		})
	}
}

func TestEncyclopediaService_GetBattleTypes(t *testing.T) {
	c := client.NewAsiaClient(getKey())
	s := api.NewEncyclopediaService(c)

	types, err := s.GetBattleTypes()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	fmt.Printf("%v\n", types)
}

func TestEncyclopediaService_GetConsumables(t *testing.T) {
	c := client.NewAsiaClient(getKey())
	s := api.NewEncyclopediaService(c)

	type args struct {
		consumablesIds  []int64
		consumablesType string
		pagination      bool
		limit           int
		page            int
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test1: Default",
			args: args{},
		},
		{
			name: "Test2: ConsumablesIds",
			args: args{
				consumablesIds: []int64{3645304016, 4235382704},
			},
		},
		{
			name: "Test3: ConsumablesType",
			args: args{
				consumablesType: constants.WowsConsumableTypePermanentCamouflages,
			},
		},
		{
			name: "Test4: Pagination",
			args: args{
				pagination: true,
				limit:      10,
				page:       2,
			},
		},
		{
			name: "Test5: More",
			args: args{
				consumablesType: constants.WowsConsumableTypePermanentCamouflages,
				pagination:      true,
				limit:           10,
				page:            2,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			consumables, err := s.GetConsumables(tt.args.consumablesIds, tt.args.consumablesType, tt.args.pagination, tt.args.limit, tt.args.page)
			if err != nil {
				t.Error(err)
				t.FailNow()
			}

			fmt.Printf("%v\n", consumables)
		})
	}
}

func TestEncyclopediaService_GetCollections(t *testing.T) {
	c := client.NewAsiaClient(getKey())
	s := api.NewEncyclopediaService(c)

	collections, err := s.GetCollections()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	fmt.Printf("%v\n", collections)
}

func TestEncyclopediaService_GetCollectionCards(t *testing.T) {
	c := client.NewAsiaClient(getKey())
	s := api.NewEncyclopediaService(c)

	cards, err := s.GetCollectionCards()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	fmt.Printf("%v\n", cards)
}

func TestEncyclopediaService_GetBattleArenas(t *testing.T) {
	c := client.NewAsiaClient(getKey())
	s := api.NewEncyclopediaService(c)

	arenas, err := s.GetBattleArenas()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	fmt.Printf("%v\n", arenas)
}
