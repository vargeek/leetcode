package design

import (
	"fmt"
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	nums := []int{-6, 10, 184}
	s := Constructor(nums)
	if !reflect.DeepEqual(s.nums, nums) {
		t.Errorf("s.nums should equal to nums")
	}
	reseting := s.Reset()
	if !reflect.DeepEqual(reseting, nums) {
		t.Errorf("reseting should equal to nums")
	}
	shuffling := s.Shuffle()
	fmt.Println(shuffling)

	s.Reset()
	shuffling = s.Shuffle()
	fmt.Println(shuffling)

	s.Reset()
	shuffling = s.Shuffle()
	fmt.Println(shuffling)

}
