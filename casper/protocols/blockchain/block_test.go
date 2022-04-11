package blockchain

import (
	"cbc-casper-go/casper"
	"cbc-casper-go/casper/simulation"
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

func TestBlock_isInBlockChain(t *testing.T) {
	str := simulation.GenerateBlockchainJsonString([]uint64{10, 11}, "", []interface{}{nil, nil})
	p, err := NewBlockchainProtocol(str, 1)
	if p == nil || err != nil {
		t.Errorf("error")
		return
	}
	p.Execute("M-0-A SJ-1-A M-1-B SJ-0-B M-0-C SJ-1-C M-1-D SJ-0-D")
	prev := p.Msgs["A"]
	for _, b := range []string{"B", "C", "D"} {
		block := p.Msgs[b]
		if !prev.(*Block).isInBlockChain(block) {
			t.Errorf("error")
		}
		if block.(*Block).isInBlockChain(prev) {
			t.Errorf("error")
		}
		prev = block
	}
}
