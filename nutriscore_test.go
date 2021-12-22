package nutriscore

import (
	"fmt"
	"testing"
)

func TestGetNutritionalScore(t *testing.T) {
	type args struct {
		n  NutritionalData
		st ScoreType
	}
	tests := []struct {
		name      string
		args      args
		want      NutritionalScore
		wantScore string
	}{
		{
			"calculates nutritional score",
			args{
				NutritionalData{
					Energy:              EnergyFromKcal(0),
					Sugars:              SugarGram(10),
					SaturatedFattyAcids: SaturatedFattyAcidsGram(2),
					Sodium:              SodiumMilligram(500),
					Fruits:              FruitsPercent(60),
					Fibre:               FibreGram(4),
					Protein:             ProteinGram(2),
				},
				Food,
			},
			NutritionalScore{2, Food},
			"B",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetNutritionalScore(tt.args.n, Food)
			if got != tt.want {
				t.Errorf("GetNutritionalScore() = %v, want %v", got, tt.want)
			}
			if gotScore := got.GetNutriScore(); gotScore != tt.wantScore {
				t.Errorf("GetNutriScore() = %v, want %v", gotScore, tt.wantScore)
			}
		})
	}
}

func TestSugarGram_GetPoints(t *testing.T) {
	tests := []struct {
		e    SugarGram
		want int
	}{
		{SugarGram(46), 10},
		{SugarGram(45), 9},
		{SugarGram(40), 8},
		{SugarGram(36), 7},
		{SugarGram(31), 6},
		{SugarGram(27), 5},
		{SugarGram(22.5), 4},
		{SugarGram(18), 3},
		{SugarGram(13.5), 2},
		{SugarGram(9), 1},
		{SugarGram(4.5), 0},
		{SugarGram(1), 0},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.e), func(t *testing.T) {
			if got := tt.e.GetPoints(); got != tt.want {
				t.Errorf("SugarGram.GetPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
