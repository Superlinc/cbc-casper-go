package concurrent

import (
	"github.com/emirpasic/gods/sets/hashset"
	"testing"
)

func Test_isValidEstimate(t *testing.T) {
	type args struct {
		estimate interface{}
	}
	estimate0 := map[string]interface{}{
		"blocks":  []*Block{nil},
		"inputs":  hashset.New(),
		"outputs": hashset.New(),
	}
	estimate1 := map[string]interface{}{
		"blocks":  []*Block{nil},
		"inputs":  hashset.New(1, 2, 3),
		"outputs": hashset.New(4, 5, 6),
	}
	estimate2 := map[string]interface{}{
		"blocks": []*Block{nil},
		"inputs": hashset.New(),
	}
	estimate3 := map[string]interface{}{
		"blocks":  []*Block{nil},
		"outputs": hashset.New(),
	}
	estimate4 := map[string]interface{}{
		"blocks": []*Block{nil},
	}
	estimate5 := map[string]interface{}{
		"blocks":  hashset.New(),
		"inputs":  hashset.New(),
		"outputs": hashset.New(),
	}
	estimate6 := map[string]interface{}{
		"inputs":  hashset.New(),
		"outputs": hashset.New(),
	}
	estimate7 := 0
	estimate8 := true
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"0", args{estimate0}, true},
		{"1", args{estimate1}, true},
		{"2", args{estimate2}, false},
		{"3", args{estimate3}, false},
		{"4", args{estimate4}, false},
		{"5", args{estimate5}, false},
		{"6", args{estimate6}, false},
		{"7", args{estimate7}, false},
		{"8", args{estimate8}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidEstimate(tt.args.estimate); got != tt.want {
				t.Errorf("isValidEstimate() = %v, want %v", got, tt.want)
			}
		})
	}
}
