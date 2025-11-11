package pokeapi

// list location areas
type RespLocations struct {
	Count int `json:"count"`
	Next *string `json:"next"`
	Previous *string `json:"previous"`
	Results []struct {
		Name string `json:"name"`
		Url string `json:"url"`
	} `json:"results"`
}

// explore location area 
type Resource struct {
    Name string `json:"name"`
    URL  string `json:"url"`
}

type LocalizedName struct {
    Language Resource `json:"language"`
    Name     string   `json:"name"`
}

type EncounterMethodRate struct {
    EncounterMethod Resource                   `json:"encounter_method"`
    VersionDetails  []EncounterMethodVerDetail `json:"version_details"`
}

type EncounterMethodVerDetail struct {
    Rate    int      `json:"rate"`
    Version Resource `json:"version"`
}

type EncounterDetail struct {
    Chance          int        `json:"chance"`
    ConditionValues []Resource `json:"condition_values"`
    MaxLevel        int        `json:"max_level"`
    Method          Resource   `json:"method"`
    MinLevel        int        `json:"min_level"`
}

type PokemonVersionDetail struct {
    EncounterDetails []EncounterDetail `json:"encounter_details"`
    MaxChance        int               `json:"max_chance"`
    Version          Resource          `json:"version"`
}

type PokemonEncounter struct {
    Pokemon        Resource               `json:"pokemon"`
    VersionDetails []PokemonVersionDetail `json:"version_details"`
}

// Top-level response for explore location (location-area).
type ExploreLocationResponse struct {
    ID                   int                     `json:"id"`
    Name                 string                  `json:"name"`
    GameIndex            int                     `json:"game_index"`
    Location             Resource                `json:"location"`
    Names                []LocalizedName         `json:"names"`
    EncounterMethodRates []EncounterMethodRate  `json:"encounter_method_rates"`
    PokemonEncounters    []PokemonEncounter      `json:"pokemon_encounters"`
}