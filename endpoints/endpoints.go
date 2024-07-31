package endpoints

import (
	"fmt"
)

const (
	base = "https://api.artifactsmmo.com"
)

// Status return the status of the game server.
//
// https://api.artifactsmmo.com/docs/#/operations/get_status__get
func Status() string {
	return fmt.Sprintf("%s/", base)
}

// ActionMove moves a character on the map using the map's X and Y position.
//
// https://api.artifactsmmo.com/docs/#/operations/action_move_my__name__action_move_post
func ActionMove(characterName string) string {
	return fmt.Sprintf("%s/my/%s/action/move", base, characterName)
}

// ActionEquip equip an item on your character.
//
// https://api.artifactsmmo.com/docs/#/operations/action_equip_item_my__name__action_equip_post
func ActionEquip(characterName string) string {
	return fmt.Sprintf("%s/my/%s/action/equip", base, characterName)
}

// ActionUnequip unequip an item on your character.
//
// https://api.artifactsmmo.com/docs/#/operations/action_unequip_item_my__name__action_unequip_post
func ActionUnequip(characterName string) string {
	return fmt.Sprintf("%s/my/%s/action/unequip", base, characterName)
}

// ActionFight start a fight against a monster on the character's map.
//
// https://api.artifactsmmo.com/docs/#/operations/action_fight_my__name__action_fight_post
func ActionFight(characterName string) string {
	return fmt.Sprintf("%s/my/%s/action/fight", base, characterName)
}

// ActionGathering harvest a resource on the character's map.
//
// https://api.artifactsmmo.com/docs/#/operations/action_gathering_my__name__action_gathering_post
func ActionGathering(characterName string) string {
	return fmt.Sprintf("%s/my/%s/action/gathering", base, characterName)
}

// ActionCrafting crafting an item. The character must be on a map with a workshop
//
// https://api.artifactsmmo.com/docs/#/operations/action_crafting_my__name__action_crafting_post
func ActionCrafting(characterName string) string {
	return fmt.Sprintf("%s/my/%s/action/crafting", base, characterName)
}

// ActionBankDeposit deposit an item in a bank on the character's map.
//
// https://api.artifactsmmo.com/docs/#/operations/action_deposit_bank_my__name__action_bank_deposit_post
func ActionBankDeposit(characterName string) string {
	return fmt.Sprintf("%s/my/%s/action/bank/deposit", base, characterName)
}

// ActionBankDepositGold deposit golds in a bank on the character's map.
//
// https://api.artifactsmmo.com/docs/#/operations/action_deposit_bank_gold_my__name__action_bank_deposit_gold_post
func ActionBankDepositGold(characterName string) string {
	return fmt.Sprintf("%s/my/%s/action/bank/deposit/gold", base, characterName)
}

// ActionRecycling recyling an item. The character must be on a map with a workshop (only for equipments and weapons).
//
// https://api.artifactsmmo.com/docs/#/operations/action_recycling_my__name__action_recycling_post
func ActionRecycling(characterName string) string {
	return fmt.Sprintf("%s/my/%s/action/recycling", base, characterName)
}

// ActionBankWithdraw take an item from your bank and put it in the character's inventory.
//
// https://api.artifactsmmo.com/docs/#/operations/action_withdraw_bank_my__name__action_bank_withdraw_post
func ActionBankWithdraw(characterName string) string {
	return fmt.Sprintf("%s/my/%s/action/bank/withdraw", base, characterName)
}

// ActionBankWithdrawGold withdraw gold from your bank.
//
// https://api.artifactsmmo.com/docs/#/operations/action_withdraw_bank_gold_my__name__action_bank_withdraw_gold_post
func ActionBankWithdrawGold(characterName string) string {
	return fmt.Sprintf("%s/my/%s/action/bank/withdraw/gold", base, characterName)
}

// ActionGeBuy buy an item at the Grand Exchange on the character's map.
//
// https://api.artifactsmmo.com/docs/#/operations/action_ge_buy_item_my__name__action_ge_buy_post
func ActionGeBuy(characterName string) string {
	return fmt.Sprintf("%s/my/%s/action/ge/buy", base, characterName)
}

// ActionGeSell sell an item at the Grand Exchange on the character's map
//
// https://api.artifactsmmo.com/docs/#/operations/action_ge_sell_item_my__name__action_ge_sell_post
func ActionGeSell(characterName string) string {
	return fmt.Sprintf("%s/my/%s/action/ge/sell", base, characterName)
}

// ActionTaskNew accepting a new task.
//
// https://api.artifactsmmo.com/docs/#/operations/action_accept_new_task_my__name__action_task_new_post
func ActionTaskNew(characterName string) string {
	return fmt.Sprintf("%s/my/%s/action/task/new", base, characterName)
}

// ActionTaskComplete complete a task.
//
// https://api.artifactsmmo.com/docs/#/operations/action_complete_task_my__name__action_task_complete_post
func ActionTaskComplete(characterName string) string {
	return fmt.Sprintf("%s/my/%s/action/task/complete", base, characterName)
}

// ActionTaskExchange exchange 3 tasks coins for a random reward. Rewards are exclusive resources for crafting items.
//
// https://api.artifactsmmo.com/docs/#/operations/action_task_exchange_my__name__action_task_exchange_post
func ActionTaskExchange(characterName string) string {
	return fmt.Sprintf("%s/my/%s/action/task/exchange", base, characterName)
}

// ActionDelete delete an item from your character's inventory.
//
// https://api.artifactsmmo.com/docs/#/operations/action_delete_item_my__name__action_delete_post
func ActionDelete(characterName string) string {
	return fmt.Sprintf("%s/my/%s/action/delete", base, characterName)
}

// Logs history of your character's actions.
//
// https://api.artifactsmmo.com/docs/#/operations/get_character_logs_my__name__logs_get
func Logs(characterName string) string {
	return fmt.Sprintf("%s/my/%s/logs", base, characterName)
}

// LogsAllCharacters history of the last actions of all your characters.
//
// https://api.artifactsmmo.com/docs/#/operations/get_all_characters_logs_my_logs_get
func LogsAllCharacters() string {
	return fmt.Sprintf("%s/my/logs", base)
}

// MyCharacters list of your characters.
//
// https://api.artifactsmmo.com/docs/#/operations/get_my_characters_my_characters_get
func MyCharacters() string {
	return fmt.Sprintf("%s/my/characters", base)
}

// BankItems fetch all items in your bank.
//
// https://api.artifactsmmo.com/docs/#/operations/get_bank_items_my_bank_items_get
func BankItems() string {
	return fmt.Sprintf("%s/my/bank/items", base)
}

// BankGolds fetch golds in your bank.
//
// https://api.artifactsmmo.com/docs/#/operations/get_bank_golds_my_bank_gold_get
func BankGolds() string {
	return fmt.Sprintf("%s/my/bank/golds", base)
}

// ChangePassword change your account password. Changing the password reset the account token.
//
// https://api.artifactsmmo.com/docs/#/operations/change_password_my_change_password_post
func ChangePassword() string {
	return fmt.Sprintf("%s/my/change_password", base)
}

// CreateCharacter create new character on your account. You can create up to 5 characters.
//
// https://api.artifactsmmo.com/docs/#/operations/create_character_characters_create_post
func CreateCharacter() string {
	return fmt.Sprintf("%s/characters/create", base)
}

// Characters fetch characters details.
//
// https://api.artifactsmmo.com/docs/#/operations/get_all_characters_characters__get
func Characters() string {
	return fmt.Sprintf("%s/characters", base)
}

// Character retrieve the details of a character.
//
// https://api.artifactsmmo.com/docs/#/operations/get_character_characters__name__get
func Character(name string) string {
	return fmt.Sprintf("%s/characters/%s", base, name)
}

// Maps fetch maps details.
//
// https://api.artifactsmmo.com/docs/#/operations/get_all_maps_maps__get
func Maps() string {
	return fmt.Sprintf("%s/maps", base)
}

// Map retrieve the details of a map.
//
// https://api.artifactsmmo.com/docs/#/operations/get_map_maps__x___y__get
func Map(x, y int) string {
	return fmt.Sprintf("%s/maps/%d/%d", base, x, y)
}

// Items fetch items details.
//
// https://api.artifactsmmo.com/docs/#/operations/get_all_items_items__get
func Items() string {
	return fmt.Sprintf("%s/items", base)
}

// Item retrieve the details of a item.
//
// https://api.artifactsmmo.com/docs/#/operations/get_item_items__code__get
func Item(code string) string {
	return fmt.Sprintf("%s/items/%s", base, code)
}

// Monsters fetch monsters details.
//
// https://api.artifactsmmo.com/docs/#/operations/get_all_monsters_monsters__get
func Monsters() string {
	return fmt.Sprintf("%s/monsters", base)
}

// Monster retrieve the details of a monster.
//
// https://api.artifactsmmo.com/docs/#/operations/get_monster_monsters__code__get
func Monster(code string) string {
	return fmt.Sprintf("%s/monsters/%s", base, code)
}

// Resources fetch resources details.
//
// https://api.artifactsmmo.com/docs/#/operations/get_all_resources_resources__get
func Resources() string {
	return fmt.Sprintf("%s/resources", base)
}

// Resource retrieve the details of a resource.
//
// https://api.artifactsmmo.com/docs/#/operations/get_resources_resources__code__get#Path-Parameters
func Resource(code string) string {
	return fmt.Sprintf("%s/resources/%s", base, code)
}

// Events fetch events details.
//
// https://api.artifactsmmo.com/docs/#/operations/get_all_events_events__get
func Events() string {
	return fmt.Sprintf("%s/events", base)
}

// GeItems fetch Grand Exchange items details.
//
// https://api.artifactsmmo.com/docs/#/operations/get_all_ge_items_ge__get
func GeItems() string {
	return fmt.Sprintf("%s/ge", base)
}

// GeItem retrieve the details of a Grand Exchange item.
//
// https://api.artifactsmmo.com/docs/#/operations/get_ge_item_ge__code__get
func GeItem(code string) string {
	return fmt.Sprintf("%s/ge/%s", base, code)
}

// CreateAccount create an account.
//
// https://api.artifactsmmo.com/docs/#/operations/create_account_accounts_create_post
func CreateAccount() string {
	return fmt.Sprintf("%s/accounts/create", base)
}

// GenerateToken use your account as HTTPBasic Auth to generate your token to use the API. You can also generate your token directly on the website.
//
// https://api.artifactsmmo.com/docs/#/operations/generate_token_token__post
func GenerateToken() string {
	return fmt.Sprintf("%s/token", base)
}
