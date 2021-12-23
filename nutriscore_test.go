package nutriscore

import (
	"encoding/json"
	"fmt"
	"testing"
)

func pp(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

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
			"water",
			args{
				NutritionalData{
					IsWater: true,
				},
				Water,
			},
			NutritionalScore{0, 0, 0, Water},
			"A",
		},
		{
			// https://world.openfoodfacts.org/product/4550002874414/pumpkin-chips-muji
			"pumpkin chips",
			args{
				NutritionalData{
					Energy:              EnergyKJ(2132),
					Sugars:              SugarGram(30),
					SaturatedFattyAcids: SaturatedFattyAcidsGram(9.6),
					Sodium:              SodiumMilligram(160),
					Fruits:              FruitsPercent(88.9),
					Fibre:               FibreGram(7.8),
					Protein:             ProteinGram(3.5),
				},
				Food,
			},
			NutritionalScore{10, 12, 22, Food},
			"C",
		},
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
			NutritionalScore{2, 6, 8, Food},
			"B",
		},
		{
			"worked example 1 from 'Nutrient Profiling Technical Guidance January 2011'",
			args{
				NutritionalData{
					Energy:              EnergyKJ(459),
					Sugars:              SugarGram(13.4),
					SaturatedFattyAcids: SaturatedFattyAcidsGram(1.8),
					Sodium:              SodiumMilligram(0.1),
					Fruits:              FruitsPercent(8),
					Fibre:               FibreGram(0.6),
					Protein:             ProteinGram(6.5),
				},
				Food,
			},
			NutritionalScore{0, 4, 4, Food},
			"B",
		},
		{
			"worked example 2 from 'Nutrient Profiling Technical Guidance January 2011'",
			args{
				NutritionalData{
					Energy:              EnergyKJ(741),
					Sugars:              SugarGram(18.7),
					SaturatedFattyAcids: SaturatedFattyAcidsGram(6.1),
					Sodium:              SodiumMilligram(60),
					Fruits:              FruitsPercent(0),
					Fibre:               FibreGram(0),
					Protein:             ProteinGram(3.6),
				},
				Food,
			},
			NutritionalScore{12, 2, 12, Food},
			"D",
		},
		{
			"worked example 5 from 'Nutrient Profiling Technical Guidance January 2011'",
			args{
				NutritionalData{
					Energy:              EnergyKJ(1504),
					Sugars:              SugarGram(35.7),
					SaturatedFattyAcids: SaturatedFattyAcidsGram(1.4),
					Sodium:              SodiumMilligram(0),
					Fruits:              FruitsPercent(46),
					Fibre:               FibreGram(4.8),
					Protein:             ProteinGram(4.3),
				},
				Food,
			},
			NutritionalScore{6, 8, 12, Food},
			"C",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetNutritionalScore(tt.args.n, tt.args.st)
			if got != tt.want {
				t.Errorf("GetNutritionalScore() = \n%+v\n, want \n%+v", pp(got), pp(tt.want))
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
		st   ScoreType
		want int
	}{
		{SugarGram(46), Food, 10},
		{SugarGram(45), Food, 9},
		{SugarGram(40), Food, 8},
		{SugarGram(36), Food, 7},
		{SugarGram(31), Food, 6},
		{SugarGram(27), Food, 5},
		{SugarGram(22.5), Food, 4},
		{SugarGram(18), Food, 3},
		{SugarGram(13.5), Food, 2},
		{SugarGram(9), Food, 1},
		{SugarGram(4.5), Food, 0},
		{SugarGram(1), Food, 0},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.e), func(t *testing.T) {
			if got := tt.e.GetPoints(tt.st); got != tt.want {
				t.Errorf("SugarGram.GetPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnergyKJ_GetPoints(t *testing.T) {
	tests := []struct {
		e    EnergyKJ
		st   ScoreType
		want int
	}{
		{EnergyKJ(280), Beverage, 10},
		{EnergyKJ(270), Beverage, 9},
		{EnergyKJ(240), Beverage, 8},
		{EnergyKJ(210), Beverage, 7},
		{EnergyKJ(0), Beverage, 0},
		{EnergyKJ(-1), Beverage, 0},
		{EnergyKJ(5), Beverage, 1},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.e), func(t *testing.T) {
			if got := tt.e.GetPoints(tt.st); got != tt.want {
				t.Errorf("EnergyKJ.GetPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
