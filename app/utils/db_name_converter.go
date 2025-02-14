package utils

import (
	"strings"

	"codecosta.com/palico/app/models"
)

func ConvertElementsToString(dbElements []models.Element) string {
	var result []string

	for _, dbElement := range dbElements {
		switch dbElement {
		case models.FIRE:
			result = append(result, models.FIRE_ELEMENT_ICON+" "+string(dbElement))
		case models.WATER:
			result = append(result, models.WATER_ELEMENT_ICON+" "+string(dbElement))
		case models.ICE:
			result = append(result, models.ICE_ELEMENT_ICON+" "+string(dbElement))
		case models.THUNDER:
			result = append(result, models.THUNDER_ELEMENT_ICON+" "+string(dbElement))
		case models.DRAGON:
			result = append(result, models.DRAGON_ELEMENT_ICON+" "+string(dbElement))
		case models.BLAST:
			result = append(result, models.BLAST_ELEMENT_ICON+" "+string(dbElement))
		case models.POISON:
			result = append(result, models.POISON_ELEMENT_ICON+" "+string(dbElement))
		case models.SLEEP:
			result = append(result, models.SLEEP_ELEMENT_ICON+" "+string(dbElement))
		case models.PARALYSIS:
			result = append(result, models.PARALYSIS_ELEMENT_ICON+" "+string(dbElement))
		case models.STUN:
			result = append(result, models.STUN_ELEMENT_ICON+" "+string(dbElement))
		case models.NONE:
			result = append(result, string(dbElement))
		default:
			LogSystemError("ConvertElementToIcon", string(dbElement)+" Element not found")
		}
	}

	if len(result) == 0 {
		return "UNKNOWN"
	}

	return strings.Join(result, "\n")
}

func ConvertTrapsToString(shockTrappable bool, pitfallTrappable bool) string {
	var result []string

	if pitfallTrappable {
		result = append(result, string(models.PITFALL_TRAP_ICON)+" Pitfall Trap")
	}
	if shockTrappable {
		result = append(result, string(models.SHOCK_TRAP_ICON)+" Shock Trap")
	}

	if len(result) == 0 {
		return "None"
	}

	return strings.Join(result, "\n")
}

func ConvertBoolToString(b bool) string {
	if b {
		return "Yes"
	}
	return "No"
}
