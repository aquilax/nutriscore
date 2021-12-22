// Package nutriscore provides utilities for calculating nutritional score and
// Nutri-Score
// More about-score: https://en.wikipedia.org/wiki/Nutri-Score
package nutriscore

// ScoreType is the type of the scored product
type ScoreType int

const (
	// Food is used when calculating nutritional score for general food items
	Food ScoreType = iota
	// Beverage is used when calculating nutritional score for beverages
	Beverage
)

var energyLevels = []float64{3350, 3015, 2680, 2345, 2010, 1675, 1340, 1005, 670, 335}
var sugarsLevels = []float64{45, 40, 36, 31, 27, 22.5, 18, 13.5, 9, 4.5}
var saturatedFattyAcidsLevels = []float64{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
var sodiumLevels = []float64{900, 810, 720, 630, 540, 450, 360, 270, 180, 90}
var fibreLevels = []float64{4.7, 3.7, 2.8, 1.9, 0.9}
var proteinLevels = []float64{8, 6.4, 4.8, 3.2, 1.6}

// NutritionalScore contains the numeric nutritional score value and type of product
type NutritionalScore struct {
	Value     int
	ScoreType ScoreType
}

// EnergyKJ represents the energy density in kJ/100g
type EnergyKJ float64

// SugarGram represents amount of sugars in grams/100g
type SugarGram float64

// SaturatedFattyAcidsGram represents amount of saturated fatty acids in grams/100g
type SaturatedFattyAcidsGram float64

// SodiumMilligram represents amount of sodium in mg/100g
type SodiumMilligram float64

// FruitsPercent represents fruits, vegetables, pulses, nuts, and rapeseed, walnut and olive oils
// as percentage of the total
type FruitsPercent float64

// FibreGram represents amount of fibre in grams/100g
type FibreGram float64

// ProteinGram represents amount of protein in grams/100g
type ProteinGram float64

// EnergyFromKcal converts energy density from kcal to EnergyKJ
func EnergyFromKcal(kcal float64) EnergyKJ {
	return EnergyKJ(kcal * 4.184)
}

// SodiumFromSalt converts salt mg/100g content to sodium content
func SodiumFromSalt(saltMg float64) SodiumMilligram {
	return SodiumMilligram(saltMg / 2.5)
}

// GetPoints returns the nutritional score
func (e EnergyKJ) GetPoints() int {
	return getPointsFromRange(float64(e), energyLevels)
}

// GetPoints returns the nutritional score
func (s SugarGram) GetPoints() int {
	return getPointsFromRange(float64(s), sugarsLevels)
}

// GetPoints returns the nutritional score
func (sfa SaturatedFattyAcidsGram) GetPoints() int {
	return getPointsFromRange(float64(sfa), saturatedFattyAcidsLevels)
}

// GetPoints returns the nutritional score
func (s SodiumMilligram) GetPoints() int {
	return getPointsFromRange(float64(s), sodiumLevels)
}

// GetPoints returns the nutritional score
func (f FruitsPercent) GetPoints() int {
	if f > 80 {
		return 5
	} else if f > 60 {
		return 2
	} else if f > 40 {
		return 1
	}
	return 0
}

// GetPoints returns the nutritional score
func (f FibreGram) GetPoints() int {
	return getPointsFromRange(float64(f), fibreLevels)
}

// GetPoints returns the nutritional score
func (p ProteinGram) GetPoints() int {
	return getPointsFromRange(float64(p), proteinLevels)
}

// NutritionalData represents the source nutritional data used for the calculation
type NutritionalData struct {
	Energy              EnergyKJ
	Sugars              SugarGram
	SaturatedFattyAcids SaturatedFattyAcidsGram
	Sodium              SodiumMilligram
	Fruits              FruitsPercent
	Fibre               FibreGram
	Protein             ProteinGram
}

// GetNutritionalScore calculates the nutritional score for nutritional data n of type st
func GetNutritionalScore(n NutritionalData, st ScoreType) NutritionalScore {
	if st != Food {
		panic("not implemented")
	}
	fruitPoints := n.Fruits.GetPoints()
	fibrePoints := n.Fibre.GetPoints()
	negative := n.Energy.GetPoints() + n.Sugars.GetPoints() + n.SaturatedFattyAcids.GetPoints() + n.Sodium.GetPoints()
	positive := fruitPoints + fibrePoints + n.Protein.GetPoints()

	if negative >= 11 && fruitPoints < 5 {
		return NutritionalScore{negative - fibrePoints - fruitPoints, st}
	}
	return NutritionalScore{negative - positive, st}
}

// GetNutriScore returns the Nutri-Score rating
func (ns NutritionalScore) GetNutriScore() string {
	scoreToLetter := []string{"A", "B", "C", "D", "E"}
	if ns.ScoreType == Food {
		return scoreToLetter[getPointsFromRange(float64(ns.Value), []float64{18, 10, 2, -1})]
	}
	return scoreToLetter[getPointsFromRange(float64(ns.Value), []float64{9, 5, 1, -2})]
}

func getPointsFromRange(v float64, steps []float64) int {
	lenSteps := len(steps)
	for i, l := range steps {
		if v > l {
			return lenSteps - i
		}
	}
	return 0
}
