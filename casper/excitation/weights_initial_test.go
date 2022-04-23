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
		wantWeights []float64
	}{
		{"1", args{length: 5}, []float64{1, 1, 1, 1, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotWeights := EqualWeights(tt.args.length); !reflect.DeepEqual(gotWeights, tt.wantWeights) {
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
			gotWeights := RandomWeights(tt.args.length)
			for _, weight := range gotWeights {
				if 0 > weight || weight >= 10 {
					t.Errorf("weight: %f out of limit", weight)
				}
			}
		})
	}
}
