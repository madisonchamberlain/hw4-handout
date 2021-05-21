package smash

import (
	"hash/fnv"
	"strings"
	"testing"
)

func eq(got, want map[uint32]uint) bool {
	if len(got) != len(want) {
		return false
	}
	for k, v := range want {
		if w, ok := got[k]; !ok || v != w {
			return false
		}
	}
	return true
}

func returnHash(s word) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func returnZero(word) uint32 {
	return 0
}

func returnLen(s word) uint32 {
	return uint32(len(s))
}

func TestSmash(t *testing.T) {
	tests := []struct {
		input   string
		smasher func(word) uint32
		want    map[uint32]uint
	}{
		{"foo", returnZero, map[uint32]uint{0: 1}},
		{"foo bar", returnZero, map[uint32]uint{0: 2}},
		{"foo \n  bar", returnZero, map[uint32]uint{0: 2}},
		{"foo ", returnLen, map[uint32]uint{3: 1}},
		{"foo f \n o \n o ", returnLen, map[uint32]uint{3: 1, 1: 3}},
		{"a ab ab abc abc abc abcd abcd abcd abcd", returnLen, map[uint32]uint{1: 1, 2: 2, 3: 3, 4: 4}},
	}
	for _, test := range tests {
		if !eq(test.want, Smash(strings.NewReader(test.input), test.smasher)) {
			t.Errorf("Smash failed on %s", test.input)
		}
	}
}

func BenchmarkSmash(b *testing.B) {
	r := strings.NewReader(" jhjhfgj jfhgjhfg jhfjghfjghfjhg fjhgfjhgfjghjf gjfghjfhgfj fjghjfghfjghfj shjfhjfhgfjhgfjghf jhfgjhfg\n jhjhfgj jfhgjhfg jhfjghfjghfjhg fjhgfjhgfjghjf gjfghjfhgfj fjghjfghfjghfj shjfhjfhgfjhgfjghf jhfgjhfg\n  jhjhfgj jfhgjhfg jhfjghfjghfjhg fjhgfjhgfjghjf gjfghjfhgfj fjghjfghfjghfj shjfhjfhgfjhgfjghf jhfgjhfg\n  jhjhfgj jfhgjhfg jhfjghfjghfjhg fjhgfjhgfjghjf gjfghjfhgfj fjghjfghfjghfj shjfhjfhgfjhgfjghf jhfgjhfg\n  jhjhfgj jfhgjhfg jhfjghfjghfjhg fjhgfjhgfjghjf gjfghjfhgfj fjghjfghfjghfj shjfhjfhgfjhgfjghf jhfgjhfg\n jhjhfgj jfhgjhfg jhfjghfjghfjhg fjhgfjhgfjghjf gjfghjfhgfj fjghjfghfjghfj shjfhjfhgfjhgfjghf jhfgjhfg  jhjhfgj jfhgjhfg jhfjghfjghfjhg fjhgfjhgfjghjf gjfghjfhgfj fjghjfghfjghfj shjfhjfhgfjhgfjghf jhfgjhfg  jhjhfgj jfhgjhfg jhfjghfjghfjhg fjhgfjhgfjghjf gjfghjfhgfj fjghjfghfjghfj shjfhjfhgfjhgfjghf jhfgjhfg \n  jhjhfgj jfhgjhfg jhfjghfjghfjhg fjhgfjhgfjghjf gjfghjfhgfj fjghjfghfjghfj shjfhjfhgfjhgfjghf jhfgjhfg  jhjhfgj jfhgjhfg jhfjghfjghfjhg fjhgfjhgfjghjf gjfghjfhgfj fjghjfghfjghfj shjfhjfhgfjhgfjghf jhfgjhfg  jhjhfgj jfhgjhfg jhfjghfjghfjhg fjhgfjhgfjghjf gjfghjfhgfj fjghjfghfjghfj shjfhjfhgfjhgfjghf jhfgjhfg\n  jhjhfgj jfhgjhfg jhfjghfjghfjhg fjhgfjhgfjghjf gjfghjfhgfj fjghjfghfjghfj shjfhjfhgfjhgfjghf jhfgjhfg \n  jhjhfgj jfhgjhfg jhfjghfjghfjhg fjhgfjhgfjghjf gjfghjfhgfj fjghjfghfjghfj shjfhjfhgfjhgfjghf jhfgjhfg")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Smash(r, returnHash)
	}

}
