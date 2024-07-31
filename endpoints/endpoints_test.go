package endpoints_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/google/uuid"
	"github.com/oustrix/gortifacts/endpoints"
	"github.com/stretchr/testify/suite"
)

type endpointsSuite struct {
	suite.Suite

	name string
}

func (s *endpointsSuite) SetupTest() {
	s.name = uuid.NewString()
}

func (s *endpointsSuite) TestStatus() {
	expectedURL := "https://api.artifactsmmo.com/"

	actualURL := endpoints.Status()
	s.Equal(expectedURL, actualURL)
}

func (s *endpointsSuite) TestActionMove() {
	expectedURL := fmt.Sprintf("https://api.artifactsmmo.com/my/%s/action/move", s.name)

	actualURL := endpoints.ActionMove(s.name)
	s.Equal(expectedURL, actualURL)
}

func (s *endpointsSuite) TestActionEquip() {
	expectedURL := fmt.Sprintf("https://api.artifactsmmo.com/my/%s/action/equip", s.name)

	actualURL := endpoints.ActionEquip(s.name)
	s.Equal(expectedURL, actualURL)
}

func (s *endpointsSuite) TestActionUnequip() {
	expectedURL := fmt.Sprintf("https://api.artifactsmmo.com/my/%s/action/unequip", s.name)

	actualURL := endpoints.ActionUnequip(s.name)
	s.Equal(expectedURL, actualURL)
}

func (s *endpointsSuite) TestActionFight() {
	expectedURL := fmt.Sprintf("https://api.artifactsmmo.com/my/%s/action/fight", s.name)

	actualURL := endpoints.ActionFight(s.name)
	s.Equal(expectedURL, actualURL)
}

func (s *endpointsSuite) TestActionGathering() {
	expectedURL := fmt.Sprintf("https://api.artifactsmmo.com/my/%s/action/gathering", s.name)

	actualURL := endpoints.ActionGathering(s.name)
	s.Equal(expectedURL, actualURL)
}

func (s *endpointsSuite) TestActionCrafting() {
	expectedURL := fmt.Sprintf("https://api.artifactsmmo.com/my/%s/action/crafting", s.name)

	actualURL := endpoints.ActionCrafting(s.name)
	s.Equal(expectedURL, actualURL)
}

func (s *endpointsSuite) TestActionBankDeposit() {
	expectedURL := fmt.Sprintf("https://api.artifactsmmo.com/my/%s/action/bank/deposit", s.name)

	actualURL := endpoints.ActionBankDeposit(s.name)
	s.Equal(expectedURL, actualURL)
}

func (s *endpointsSuite) TestActionBankDepositGold() {
	expectedURL := fmt.Sprintf("https://api.artifactsmmo.com/my/%s/action/bank/deposit/gold", s.name)

	actualURL := endpoints.ActionBankDepositGold(s.name)
	s.Equal(expectedURL, actualURL)
}

func (s *endpointsSuite) TestActionRecycling() {
	expectedURL := fmt.Sprintf("https://api.artifactsmmo.com/my/%s/action/recycling", s.name)

	actualURL := endpoints.ActionRecycling(s.name)
	s.Equal(expectedURL, actualURL)
}

func (s *endpointsSuite) TestActionBankWithdraw() {
	expectedURL := fmt.Sprintf("https://api.artifactsmmo.com/my/%s/action/bank/withdraw", s.name)

	actualURL := endpoints.ActionBankWithdraw(s.name)
	s.Equal(expectedURL, actualURL)
}

func (s *endpointsSuite) TestActionBankWithdrawGold() {
	expectedURL := fmt.Sprintf("https://api.artifactsmmo.com/my/%s/action/bank/withdraw/gold", s.name)

	actualURL := endpoints.ActionBankWithdrawGold(s.name)
	s.Equal(expectedURL, actualURL)
}

func (s *endpointsSuite) TestActionGeBuy() {
	expectedURL := fmt.Sprintf("https://api.artifactsmmo.com/my/%s/action/ge/buy", s.name)

	actualURL := endpoints.ActionGeBuy(s.name)
	s.Equal(expectedURL, actualURL)
}

func (s *endpointsSuite) TestActionGeSell() {
	expectedURL := fmt.Sprintf("https://api.artifactsmmo.com/my/%s/action/ge/sell", s.name)

	actualURL := endpoints.ActionGeSell(s.name)
	s.Equal(expectedURL, actualURL)
}

func (s *endpointsSuite) TestActionTaskNew() {
	expectedURL := fmt.Sprintf("https://api.artifactsmmo.com/my/%s/action/task/new", s.name)

	actualURL := endpoints.ActionTaskNew(s.name)
	s.Equal(expectedURL, actualURL)
}

func (s *endpointsSuite) TestActionTaskComplete() {
	expectedURL := fmt.Sprintf("https://api.artifactsmmo.com/my/%s/action/task/complete", s.name)

	actualURL := endpoints.ActionTaskComplete(s.name)
	s.Equal(expectedURL, actualURL)
}

func (s *endpointsSuite) TestActionTaskExchange() {
	expectedURL := fmt.Sprintf("https://api.artifactsmmo.com/my/%s/action/task/exchange", s.name)

	actualURL := endpoints.ActionTaskExchange(s.name)
	s.Equal(expectedURL, actualURL)
}

func (s *endpointsSuite) TestActionDelete() {
	expectedURL := fmt.Sprintf("https://api.artifactsmmo.com/my/%s/action/delete", s.name)

	actualURL := endpoints.ActionDelete(s.name)
	s.Equal(expectedURL, actualURL)
}

func (s *endpointsSuite) TestLogs() {
	expectedURL := fmt.Sprintf("https://api.artifactsmmo.com/my/%s/logs", s.name)

	actualURL := endpoints.Logs(s.name)
	s.Equal(expectedURL, actualURL)
}

func (s *endpointsSuite) TestLogsAllCharacters() {
	expectedURL := fmt.Sprintf("https://api.artifactsmmo.com/my/logs")

	actualURL := endpoints.LogsAllCharacters()
	s.Equal(expectedURL, actualURL)
}

func (s *endpointsSuite) TestMyCharacters() {
	expectedURL := fmt.Sprintf("https://api.artifactsmmo.com/my/characters")

	actualURL := endpoints.MyCharacters()
	s.Equal(expectedURL, actualURL)
}

func (s *endpointsSuite) TestBankItems() {
	expectedURL := fmt.Sprintf("https://api.artifactsmmo.com/my/bank/items")

	actualURL := endpoints.BankItems()
	s.Equal(expectedURL, actualURL)
}

func (s *endpointsSuite) TestBankGolds() {
	expectedURL := fmt.Sprintf("https://api.artifactsmmo.com/my/bank/golds")

	actualURL := endpoints.BankGolds()
	s.Equal(expectedURL, actualURL)
}

func (s *endpointsSuite) TestChangePassword() {
	expectedURL := fmt.Sprintf("https://api.artifactsmmo.com/my/change_password")

	actualURL := endpoints.ChangePassword()
	s.Equal(expectedURL, actualURL)
}

func (s *endpointsSuite) TestCreateCharacter() {
	expectedURL := fmt.Sprintf("https://api.artifactsmmo.com/characters/create")

	actualURL := endpoints.CreateCharacter()
	s.Equal(expectedURL, actualURL)
}

func (s *endpointsSuite) TestCharacters() {
	expectedURL := fmt.Sprintf("https://api.artifactsmmo.com/characters")

	actualURL := endpoints.Characters()
	s.Equal(expectedURL, actualURL)
}

func (s *endpointsSuite) TestCharacter() {
	expectedURL := fmt.Sprintf("https://api.artifactsmmo.com/characters/%s", s.name)

	actualURL := endpoints.Character(s.name)
	s.Equal(expectedURL, actualURL)
}

func (s *endpointsSuite) TestMaps() {
	expectedURL := fmt.Sprintf("https://api.artifactsmmo.com/maps")

	actualURL := endpoints.Maps()
	s.Equal(expectedURL, actualURL)
}

func (s *endpointsSuite) TestMap() {
	x, y := rand.Int(), rand.Int()
	expectedURL := fmt.Sprintf("https://api.artifactsmmo.com/maps/%d/%d", x, y)

	actualURL := endpoints.Map(x, y)
	s.Equal(expectedURL, actualURL)
}

func (s *endpointsSuite) TestItems() {
	expectedURL := fmt.Sprintf("https://api.artifactsmmo.com/items")

	actualURL := endpoints.Items()
	s.Equal(expectedURL, actualURL)
}

func (s *endpointsSuite) TestItem() {
	expectedURL := fmt.Sprintf("https://api.artifactsmmo.com/items/%s", s.name)

	actualURL := endpoints.Item(s.name)
	s.Equal(expectedURL, actualURL)
}

func (s *endpointsSuite) TestMonsters() {
	expectedURL := fmt.Sprintf("https://api.artifactsmmo.com/monsters")

	actualURL := endpoints.Monsters()
	s.Equal(expectedURL, actualURL)
}

func (s *endpointsSuite) TestMonster() {
	expectedURL := fmt.Sprintf("https://api.artifactsmmo.com/monsters/%s", s.name)

	actualURL := endpoints.Monster(s.name)
	s.Equal(expectedURL, actualURL)
}

func (s *endpointsSuite) TestResources() {
	expectedURL := fmt.Sprintf("https://api.artifactsmmo.com/resources")

	actualURL := endpoints.Resources()
	s.Equal(expectedURL, actualURL)
}

func (s *endpointsSuite) TestResource() {
	expectedURL := fmt.Sprintf("https://api.artifactsmmo.com/resources/%s", s.name)

	actualURL := endpoints.Resource(s.name)
	s.Equal(expectedURL, actualURL)
}

func (s *endpointsSuite) TestEvents() {
	expectedURL := fmt.Sprintf("https://api.artifactsmmo.com/events")

	actualURL := endpoints.Events()
	s.Equal(expectedURL, actualURL)
}

func (s *endpointsSuite) TestGeItems() {
	expectedURL := fmt.Sprintf("https://api.artifactsmmo.com/ge")

	actualURL := endpoints.GeItems()
	s.Equal(expectedURL, actualURL)
}

func (s *endpointsSuite) TestGeItem() {
	expectedURL := fmt.Sprintf("https://api.artifactsmmo.com/ge/%s", s.name)

	actualURL := endpoints.GeItem(s.name)
	s.Equal(expectedURL, actualURL)
}

func (s *endpointsSuite) CreateAccount() {
	expectedURL := fmt.Sprintf("https://api.artifactsmmo.com/accounts/create")

	actualURL := endpoints.CreateAccount()
	s.Equal(expectedURL, actualURL)
}

func (s *endpointsSuite) TestGenerateToken() {
	expectedURL := fmt.Sprintf("https://api.artifactsmmo.com/token")

	actualURL := endpoints.GenerateToken()
	s.Equal(expectedURL, actualURL)
}

func TestEndpoints(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(endpointsSuite))
}
