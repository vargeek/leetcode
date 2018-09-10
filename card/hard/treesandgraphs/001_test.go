package treesandgraphs

import (
	"testing"
)

func Test_ladderLength(t *testing.T) {
	bw1 := "hit"
	ew1 := "cog"
	wl1 := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	want1 := 5

	bw2 := "hit"
	ew2 := "cog"
	wl2 := []string{"hot", "dot", "dog", "lot", "log"}
	want2 := 0

	bw3 := "a"
	ew3 := "c"
	wl3 := []string{"a", "b", "c"}
	want3 := 2

	bw4 := "hit"
	ew4 := "cog"
	wl4 := []string{"hot", "cog", "dot", "dog", "hit", "lot", "log"}
	want4 := 5

	bw5 := "hit"
	ew5 := "cog"
	wl5 := []string{"hot", "hit", "cog", "dot", "dog"}
	want5 := 5

	type args struct {
		beginWord string
		endWord   string
		wordList  []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{bw1, ew1, wl1}, want1},
		{"2", args{bw2, ew2, wl2}, want2},
		{"3", args{bw3, ew3, wl3}, want3},
		{"4", args{bw4, ew4, wl4}, want4},
		{"5", args{bw5, ew5, wl5}, want5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ladderLength(tt.args.beginWord, tt.args.endWord, tt.args.wordList); got != tt.want {
				t.Errorf("ladderLength() = %v, want %v", got, tt.want)
			}
		})
	}

}
