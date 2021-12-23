# nutriscore [![Go Reference](https://pkg.go.dev/badge/github.com/aquilax/nutriscore.svg)](https://pkg.go.dev/github.com/aquilax/nutriscore)

Go library for calculating the Nutri-Score

Based on https://www.santepubliquefrance.fr/content/download/150263/file/2021_07_21_QR_scientifique_et_technique_V41_EN.pdf

## Usage

```go
package nutriscore_test

import (
	"fmt"

	ns "github.com/aquilax/nutriscore"
)

func ExampleGetNutritionalScore() {
	ns := ns.GetNutritionalScore(ns.NutritionalData{
		Energy:              ns.EnergyFromKcal(0),
		Sugars:              ns.SugarGram(10),
		SaturatedFattyAcids: ns.SaturatedFattyAcidsGram(2),
		Sodium:              ns.SodiumMilligram(500),
		Fruits:              ns.FruitsPercent(60),
		Fibre:               ns.FibreGram(4),
		Protein:             ns.ProteinGram(2),
	}, ns.Food)
	fmt.Printf("Nutritional score: %d\n", ns.Value)
	fmt.Printf("NutriScore: %s\n", ns.GetNutriScore())
	// Output:
	// Nutritional score: 2
	// NutriScore: B
}
```
