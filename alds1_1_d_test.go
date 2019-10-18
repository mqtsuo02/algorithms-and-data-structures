package answer

import "testing"

func TestMaxValue(t *testing.T) {

	type QA struct {
		Input  []int
		Output int
	}

	qas := []QA{
		QA{Input: []int{6, 5, 3, 1, 3, 4, 3}, Output: 3},
		QA{Input: []int{5, 4, 3, 2}, Output: -1},
	}

	for _, qa := range qas {
		res := maxValue(qa.Input)
		if res != qa.Output {
			t.Fatal("test failed", qa, res)
		}
	}
}
