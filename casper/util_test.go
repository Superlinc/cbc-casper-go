package casper

import "testing"

func TestMaxUint(t *testing.T) {
	tmp := MaxUint(1, 2)
	if tmp != 2 {
		t.Errorf("wrong answer")
	}
}

func TestGetRandomStr(t *testing.T) {
	str := GetRandomStr(10)
	if len(str) != 10 {
		t.Errorf("wrong length")
	}
	for _, c := range str {
		if c < 'a' || c > 'z' {
			t.Errorf("wrong content")
		}
	}
}
