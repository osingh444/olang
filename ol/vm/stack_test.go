package vm

import (
	"testing"
)

func TestStaccc(t *testing.T) {

	s := CreateStack()

	s.Push(1)

	if s.Pop() != 1 {
		t.Error("case 1 failed")
	}

	if s.Pop() != 0 {
		t.Error("null case failed")
	}

	nums := []int{4, 3, 2, 1}

	for _, num := range nums {
		s.Push(num)
	}

	index := 3

	for index > 0 {
		if s.Pop() != nums[index] {
			t.Error("case 2 failed")
		}
		index -= 1
	}
}
