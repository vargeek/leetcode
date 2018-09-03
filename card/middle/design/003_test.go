package design

import (
	"fmt"
	"testing"
)

func TestRandomizedSet(t *testing.T) {
	randomSet := Constructor()
	want1 := true
	if got := randomSet.Insert(1); got != want1 {
		t.Errorf("randomSet.Insert(1) = %v, want %v", got, want1)
	}

	want2 := false
	if got := randomSet.Remove(2); got != want2 {
		t.Errorf("randomSet.Remove(2) = %v, want %v", got, want2)
	}

	want3 := true
	if got := randomSet.Insert(2); got != want3 {
		t.Errorf("randomSet.Insert(2) = %v, want %v", got, want3)
	}

	fmt.Println(randomSet.GetRandom())
	fmt.Println(randomSet.GetRandom())
	fmt.Println(randomSet.GetRandom())
	fmt.Println(randomSet.GetRandom())
	fmt.Println(randomSet.GetRandom())
	fmt.Println(randomSet.GetRandom())
	fmt.Println(randomSet.GetRandom())
	fmt.Println(randomSet.GetRandom())
	fmt.Println(randomSet.GetRandom())
	fmt.Println(randomSet.GetRandom())
	fmt.Println(randomSet.GetRandom())

	want4 := true
	if got := randomSet.Remove(1); got != want4 {
		t.Errorf("randomSet.Remove(1) = %v, want %v", got, want4)
	}

	want5 := false
	if got := randomSet.Insert(2); got != want5 {
		t.Errorf("randomSet.Remove(1) = %v, want %v", got, want5)
	}

	want6 := 2
	if got := randomSet.GetRandom(); got != want6 {
		t.Errorf("randomSet.Remove(1) = %v, want %v", got, want6)
	}

}
