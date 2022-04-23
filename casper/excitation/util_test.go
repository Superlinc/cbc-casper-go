package excitation

import (
	"testing"
)

func TestRateOfWeightMoreThanThreshold(t *testing.T) {
	type args struct {
		weights   []float64
		threshold float64
	}
	tests := []struct {
		name       string
		args       args
		wantLength int
		wantRate   float64
	}{
		{"1", args{
			weights:   []float64{1, 1, 3, 5},
			threshold: 2.0 / 3,
		}, 2, 0.5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLength, gotRate := RateOfWeightMoreThanThreshold(tt.args.weights, tt.args.threshold)
			if gotLength != tt.wantLength {
				t.Errorf("RateOfWeightMoreThanThreshold() gotLength = %v, want %v", gotLength, tt.wantLength)
			}
			if gotRate != tt.wantRate {
				t.Errorf("RateOfWeightMoreThanThreshold() gotRate = %v, want %v", gotRate, tt.wantRate)
			}
		})
	}
}

func TestSum(t *testing.T) {
	type args struct {
		arr []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"1", args{arr: []float64{1, 2, 5, 0.5}}, 8.5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.args.arr); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
