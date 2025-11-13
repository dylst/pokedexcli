package pokeapi

// Top-level Pokemon response (most commonly used fields).
type Pokemon struct {
    ID             int                  `json:"id"`
    Name           string               `json:"name"`
    BaseExperience int                  `json:"base_experience"`
    Height         int                  `json:"height"`
    Weight         int                  `json:"weight"`
    Abilities      []PokemonAbility     `json:"abilities"`
    Forms          []Resource           `json:"forms"`
    GameIndices    []GameIndex          `json:"game_indices"`
    HeldItems      []HeldItem           `json:"held_items"`
    LocationAreaEncounters string        `json:"location_area_encounters"`
    Moves          []PokemonMove        `json:"moves"`
    Species        Resource             `json:"species"`
    Stats          []PokemonStat        `json:"stats"`
    Types          []PokemonTypeSlot    `json:"types"`
}

// Abilities
type PokemonAbility struct {
    IsHidden bool     `json:"is_hidden"`
    Slot     int      `json:"slot"`
    Ability  Resource `json:"ability"`
}

// Game index info
type GameIndex struct {
    GameIndex int      `json:"game_index"`
    Version   Resource `json:"version"`
}

// Held items
type HeldItem struct {
    Item           Resource                 `json:"item"`
    VersionDetails []HeldItemVersionDetail  `json:"version_details"`
}

type HeldItemVersionDetail struct {
    Rarity  int      `json:"rarity"`
    Version Resource `json:"version"`
}

// Moves and version-group specific info
type PokemonMove struct {
    Move                    Resource                   `json:"move"`
    VersionGroupDetails     []MoveVersionGroupDetail   `json:"version_group_details"`
}

type MoveVersionGroupDetail struct {
    MoveLearnMethod Resource `json:"move_learn_method"`
    VersionGroup    Resource `json:"version_group"`
    LevelLearnedAt  int      `json:"level_learned_at"`
}

// Stats
type PokemonStat struct {
    BaseStat int      `json:"base_stat"`
    Effort   int      `json:"effort"`
    Stat     Resource `json:"stat"`
}

// Types
type PokemonTypeSlot struct {
    Slot int      `json:"slot"`
    Type Resource `json:"type"`
}

