package gortifacts

import (
	"time"
)

// Cooldown represents cooldown for next action.
type Cooldown struct {
	TotalSeconds     int            `json:"total_seconds"`     // The total seconds of the cooldown.
	RemainingSeconds int            `json:"remaining_seconds"` // The remaining seconds of the cooldown.
	Expiration       time.Time      `json:"expiration"`        // The expiration of the cooldown.
	Reason           CooldownReason `json:"reason"`            // The reason of the cooldown.
}

// CooldownReason is the reason of the cooldown
type CooldownReason string

const (
	CooldownreasonMovement     CooldownReason = "movement"
	CooldownReasonFight        CooldownReason = "fight"
	CooldownReasonCrafting     CooldownReason = "crafting"
	CooldownReasonGathering    CooldownReason = "gathering"
	CooldownReasonBuyGe        CooldownReason = "buy_ge"
	CooldownReasonSellGe       CooldownReason = "sell_ge"
	CooldownReasonDeleteItem   CooldownReason = "delete_item"
	CooldownreasonBankDeposite CooldownReason = "deposit_bank"
	CooldownReasonBankWithdraw CooldownReason = "withdraw_bank"
	CooldownReasonEquip        CooldownReason = "equip"
	CooldownReasonUnquip       CooldownReason = "unequip"
	CooldownReasonTask         CooldownReason = "task"
	CooldownReasonRecycling    CooldownReason = "recycling"
)

// Character represents character.
type Character struct {
	Name                    string        `json:"name"`                      // Name of the character.
	Skin                    CharacterSkin `json:"skin"`                      // Character skin code.
	Level                   int           `json:"level"`                     // Combat level.
	XP                      int           `json:"xp"`                        // The current xp level of the combat level.
	MaxXP                   int           `json:"max_xp"`                    // XP required to level up the character.
	TotalXP                 int           `json:"total_xp"`                  // Total XP of your character.
	Gold                    int           `json:"gold"`                      // The numbers of golds on this character.
	Speed                   int           `json:"speed"`                     // *Not available, on the roadmap. Character movement speed.
	MiningLevel             int           `json:"mining_level"`              // Mining level.
	MiningXP                int           `json:"mining-xp"`                 // The current xp level of the Mining skill.
	MiningMaxXP             int           `json:"mining_max_xp"`             // Mining XP required to level up the skill.
	WoodcuttingLevel        int           `json:"woodcutting_level"`         // Woodcutting level.
	WoodcuttingXP           int           `json:"woodcutting_xp"`            // The current xp level of the Woodcutting skill.
	WoodcuttingMaxXP        int           `json:"woodcutting_max_xp"`        // Woodcutting XP required to level up the skill.
	FishingLevel            int           `json:"fishing_level"`             // Fishing level.
	FishingXP               int           `json:"fishing_xp"`                // The current xp level of the Fishing skill.
	FishingMaxXP            int           `json:"fishing_max_xp"`            // Fishing XP required to level up the skill.
	WeaponcraftingLevel     int           `json:"weaponcrafting_level"`      // Weaponcrafting level.
	WeaponcraftingXP        int           `json:"weaponcrafting_xp"`         // The current xp level of the Weaponcrafting skill.
	WeaponcraftingMaxXP     int           `json:"weaponcrafting_max_xp"`     // Weaponcrafting XP required to level up the skill.
	GearcraftingLevel       int           `json:"gearcrafting_level"`        // Gearcrafting level.
	GearcraftingXP          int           `json:"gearcrafting_xp"`           // The current xp level of the Gearcrafting skill.
	GearcraftingMaxXP       int           `json:"gearcrafting_max_xp"`       // Gearcrafting XP required to level up the skill.
	JewelrycraftingLevel    int           `json:"jewelrycrafting_level"`     // Jewelrycrafting level.
	JewelrycraftingXP       int           `json:"jewelrycrafting_xp"`        // The current xp level of the Jewelrycrafting skill
	JewelrycraftingMaxXP    int           `json:"jewelrycrafting_max_xp"`    // Jewelrycrafting XP required to level up the skill.
	CookingLevel            int           `json:"cooking_level"`             // The current xp level of the Cooking skill.
	CookingXP               int           `json:"cooking_xp"`                // Cooking XP.
	CookingMaxXP            int           `json:"cooking_max_xp"`            // Cooking XP required to level up the skill.
	HP                      int           `json:"hp"`                        // Character HP.
	Haste                   int           `json:"haste"`                     // *Character Haste. Increase speed attack (reduce fight cooldown)
	CriticalStrike          int           `json:"critical_strike"`           // *Not available, on the roadmap. Character Critical Strike. Critical strikes increase the attack's damage.
	Stamina                 int           `json:"stamina"`                   // *Not available, on the roadmap. Regenerates life at the start of each turn.
	AttackFire              int           `json:"attack_fire"`               // Fire attack.
	AttackEarth             int           `json:"attack_earth"`              // Earth attack.
	AttackWater             int           `json:"attack_water"`              // Water attack.
	AttackAir               int           `json:"attack_air"`                // Air attack.
	DmgFire                 int           `json:"dmg_fire"`                  // % Fire damage.
	DmgEarth                int           `json:"dmg_earth"`                 // % Earth damage.
	DmgWater                int           `json:"dmg_water"`                 // % Water damage.
	DmgAir                  int           `json:"dmg_air"`                   // % Air damage.
	ResFire                 int           `json:"res_fire"`                  // % Fire resistance.
	ResEarth                int           `json:"res_earth"`                 // % Earth resistance.
	ResWater                int           `json:"res_water"`                 // % Water resistance.
	ResAir                  int           `json:"res_air"`                   // % Air resistance.
	X                       int           `json:"x"`                         // Character x coordinate.
	Y                       int           `json:"y"`                         // Character y coordinate.
	Cooldown                int           `json:"cooldown"`                  // Cooldown in seconds.
	CooldownExpiration      time.Time     `json:"cooldown_expiration"`       // Datetime Cooldown expiration.
	WeaponSlot              string        `json:"weapon_slot"`               // Weapon slot.
	ShieldSlot              string        `json:"shield_slot"`               // Shield slot.
	HelmetSlot              string        `json:"helmet_slot"`               // Helmet slot.
	BodyArmorSlot           string        `json:"body_armor_slot"`           // Body armor slot.
	LegArmorSlot            string        `json:"leg_armor_slot"`            // Leg armor slot.
	BootsSlot               string        `json:"boots_slot"`                // Boots slot.
	Ring1Slot               string        `json:"ring1_slot"`                // Ring 1 slot.
	Ring2Slot               string        `json:"ring2_slot"`                // Ring 2 slot.
	AmuletSlot              string        `json:"amulet_slot"`               // Amulet slot.
	Artifact1Slot           string        `json:"artifact1_slot"`            // Artifact 1 slot.
	Artifact2Slot           string        `json:"artifact2_slot"`            // Artifact 2 slot.
	Artifact3Slot           string        `json:"artifact3_slot"`            // Artifact 3 slot.
	Consumable1Slot         string        `json:"consumable1_slot"`          // Consumable 1 slot.
	Consumable1SlotQuantity int           `json:"consumable1_slot_quantity"` // Consumable 1 slot quantity.
	Consumable2Slot         string        `json:"consumable2_slot"`          // Consumable 2 slot.
	Consumable2SlotQuantity int           `json:"consumable2_slot_quantity"` // Consumable 2 slot quantity.
	Task                    string        `json:"task"`                      // Task in progress.
	TaskType                string        `json:"task_type"`                 // Task type.
	TaskProgress            int           `json:"task_progress"`             // Task progression.
	TaskTotal               int           `json:"task_total"`                // Task total objective.
	InventoryMaxItems       int           `json:"inventory_max_items"`       // Inventory max items.
	Inventory               []struct {
		Slot     int    `json:"slot"`     // Inventory slot identifier.
		Code     string `json:"code"`     // Item code.
		Quantity int    `json:"quantity"` // Quantity in the slot.
	} `json:"inventory"` // List of inventory slots.
}

// CharacterSkin is character's skin code.
type CharacterSkin string

const (
	CharacterSkinMen1   CharacterSkin = "men1"
	CharacterSkinMen2   CharacterSkin = "men2"
	CharacterSkinMen3   CharacterSkin = "men3"
	CharacterSkinWomen1 CharacterSkin = "women1"
	CharacterSkinWomen2 CharacterSkin = "women2"
	CharacterSkinWomen3 CharacterSkin = "women3"
)
