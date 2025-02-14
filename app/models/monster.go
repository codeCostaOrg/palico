package models

type Element string
type Species string
type Size string

const (
	AMPHIBIAN      Species = "AMPHIBIAN"
	BIRD_WYVERN    Species = "BIRD_WYVERN"
	BRUTE_WYVERN   Species = "BRUTE_WYVERN"
	ELDER_DRAGON   Species = "ELDER_DRAGON"
	FANGED_BEAST   Species = "FANGED_BEAST"
	FANGED_WYVERN  Species = "FANGED_WYVERN"
	FISH           Species = "FISH"
	FLYING_WYVERN  Species = "FLYING_WYVERN"
	HERBIVORE      Species = "HERBIVORE"
	LEVIATHAN      Species = "LEVIATHAN"
	LYNIAN         Species = "LYNIAN"
	NEOPTERON      Species = "NEOPTERON"
	PISCINE_WYVERN Species = "PISCINE_WYVERN"
	TEMNOCERAN     Species = "TEMNOCERAN"
	UNKNOWN        Species = "UNKNOWN"
	WINGDRAKE      Species = "WINGDRAKE"

	FIRE      Element = "FIRE"
	WATER     Element = "WATER"
	ICE       Element = "ICE"
	THUNDER   Element = "THUNDER"
	DRAGON    Element = "DRAGON"
	BLAST     Element = "BLAST"
	POISON    Element = "POISON"
	SLEEP     Element = "SLEEP"
	PARALYSIS Element = "PARALYSIS"
	STUN      Element = "STUN"
	NONE      Element = "NONE"

	SMALL Size = "SMALL"
	LARGE Size = "LARGE"
)

type MonsterResponse struct {
	Name             string  `json:"Name"`
	Description      string  `json:"Description"`
	Species          Species `json:"Species"`
	Size             Size    `json:"Size"`
	TailSeverable    bool    `json:"TailSeverable"`
	ShockTrappable   bool    `json:"ShockTrappable"`
	PitfallTrappable bool    `json:"PitfallTrappable"`
	AvatarURL        string  `json:"AvatarURL"`
	ImgURL           string  `json:"ImgURL"`

	Elements    []Element `json:"Elements"`
	Weaknesses  []Element `json:"Weaknesses"`
	Resistances []Element `json:"Resistances"`

	Error string `json:"Error"`
}
