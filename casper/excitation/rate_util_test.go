package excitation

import (
	"reflect"
	"testing"
)

func TestGetRateSlice(t *testing.T) {
	type args struct {
		validators []Validator
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{
			"1",
			args{[]Validator{&NormalValidator{1}, &NormalValidator{1}, &NormalValidator{3}, &NormalValidator{5}}},
			[]float64{0.1, 0.2, 0.5, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetRateSlice(tt.args.validators); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRateSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetRateSliceWithLimit(t *testing.T) {
	type args struct {
		validators []Validator
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{
			"1",
			args{[]Validator{&NormalValidator{1}, &NormalValidator{1}, &NormalValidator{3}, &NormalValidator{5}}},
			[]float64{0.25, 0.5, 0.75, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetRateSliceWithLimit(tt.args.validators); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRateSliceWithLimit() = %v, want %v", got, tt.want)
			}
		})
	}
}
