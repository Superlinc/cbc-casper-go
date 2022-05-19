package excitation

import "testing"

func TestRateOfWeightMoreThanThreshold(t *testing.T) {
	type args struct {
		validators []Validator
	}
	tests := []struct {
		name     string
		args     args
		wantRate float64
	}{
		{"1", args{
			validators: []Validator{&NormalValidator{1}, &NormalValidator{1}, &NormalValidator{3}, &NormalValidator{5}},
		}, 0.25},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRate := RateOfWeightMoreThanThreshold(tt.args.validators)
			if gotRate != tt.wantRate {
				t.Errorf("RateOfWeightMoreThanThreshold() gotRate = %v, want %v", gotRate, tt.wantRate)
			}
		})
	}
}

func TestGini(t *testing.T) {
	type args struct {
		validators []Validator
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "1",
			args: args{
				validators: []Validator{&NormalValidator{1}, &NormalValidator{1}, &NormalValidator{1}, &NormalValidator{1}},
			},
			want: 0.0,
		},
		{
			name: "2",
			args: args{
				validators: []Validator{&NormalValidator{1}, &NormalValidator{0}, &NormalValidator{0}, &NormalValidator{0}},
			},
			want: 0.75,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Gini(tt.args.validators); got != tt.want {
				t.Errorf("Gini() = %v, want %v", got, tt.want)
			}
		})
	}
}
