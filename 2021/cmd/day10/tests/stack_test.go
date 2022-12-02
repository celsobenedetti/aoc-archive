package common_test

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/celso-patiri/aoc/cmd/day10/common"
)

type Stack = common.Stack

func TestQueue(t *testing.T) {
	stack := common.MakeNewStack()

	for i := 0; i < 1000; i++ {
		str := strconv.FormatInt(int64(i), 10)
		stack.Push(str)

		if !reflect.DeepEqual(str, stack.Top()) {
			t.Errorf("Stack error: expected %s, got %s", str, stack.Top())
		}

		for j := 0; j < stack.Size()-1; j++ {
			top := stack.Top()
			deq := stack.Pop()

			if !reflect.DeepEqual(top, deq) {
				t.Errorf("Stack error: expected %s, got %s", top, deq)
			}
		}

		for j := 0; j < 10; j++ {
			str = strconv.FormatInt(int64(j), 10)
			stack.Push(str)

			if !reflect.DeepEqual(str, stack.Top()) {
				t.Errorf("Stack error: expected %s, got %s", str, stack.Top())
			}
		}

	}

}
