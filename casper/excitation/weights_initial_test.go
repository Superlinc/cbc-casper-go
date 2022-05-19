package excitation

import (
	"reflect"
	"testing"
)

func TestEqualWeights(t *testing.T) {
	type args struct {
		length int
	}
	tests := []struct {
		name        string
		args        args
		wantWeights []Validator
	}{
		{"1", args{length: 3}, []Validator{&NormalValidator{weight: 1}, &NormalValidator{weight: 1}, &NormalValidator{weight: 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotWeights := EqualWeights(tt.args.length, false); !reflect.DeepEqual(gotWeights, tt.wantWeights) {
				t.Errorf("EqualWeights() = %v, want %v", gotWeights, tt.wantWeights)
			}
		})
	}
}

func TestRandomWeights(t *testing.T) {
	type args struct {
		length int
	}
	tests := []struct {
		name string
		args args
	}{
		{"1", args{length: 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotWeights := RandomWeights(tt.args.length, false)
			for _, validator := range gotWeights {
				if 0 >= validator.Weight() || validator.Weight() > 10 {
					t.Errorf("weight: %d out of limit", validator.Weight())
				}
			}
		})
	}
}
