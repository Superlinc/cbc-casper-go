package blockchain

import (
	"cbc-casper-go/casper"
	"testing"
)

func TestBlock_ConflictWith(t *testing.T) {
	type fields struct {
		Message *Block
	}
	type args struct {
		message casper.Messager
	}
	b1 := NewBlock(nil, nil, nil, 0, 0)
	b2 := NewBlock(b1, nil, nil, 0, 0)
	b3 := NewBlock(nil, nil, nil, 0, 0)
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{"one", fields{Message: b1}, args{message: b2}, false, false},
		{"two", fields{Message: b1}, args{message: b3}, true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := tt.fields.Message
			got, err := b.ConflictWith(tt.args.message)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConflictWith() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ConflictWith() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValidEstimate(t *testing.T) {
	type args struct {
		estimate interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{nil}, true},
		{"", args{0}, false},
		{"", args{true}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidEstimate(tt.args.estimate); got != tt.want {
				t.Errorf("isValidEstimate() = %v, want %v", got, tt.want)
			}
		})
	}
}
