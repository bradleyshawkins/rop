package rop_test

import (
	"testing"

	"github.com/bradleyshawkins/rop"
)

func TestServePancakes(t *testing.T) {
	tests := []struct {
		Name              string
		Stack             rop.Stack
		ExpectedFlipCount int
	}{
		{
			Name: "-",
			Stack: rop.Stack{
				rop.Pancake(false),
			},
			ExpectedFlipCount: 1,
		},
		{
			Name: "-+",
			Stack: rop.Stack{
				rop.Pancake(false),
				rop.Pancake(true),
			},
			ExpectedFlipCount: 1,
		},
		{
			Name: "+-",
			Stack: rop.Stack{
				rop.Pancake(true),
				rop.Pancake(false),
			},
			ExpectedFlipCount: 2,
		},
		{
			Name: "+-",
			Stack: rop.Stack{
				rop.Pancake(true),
				rop.Pancake(true),
				rop.Pancake(true),
			},
			ExpectedFlipCount: 0,
		},
		{
			Name: "--+-",
			Stack: rop.Stack{
				rop.Pancake(false),
				rop.Pancake(false),
				rop.Pancake(true),
				rop.Pancake(false),
			},
			ExpectedFlipCount: 3,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			waiter := rop.Waiter{}
			actualFlipCount, err := waiter.ServePancakes(test.Stack)
			if err != nil {
				t.Errorf("Unexpected error while serving pancakes. Error: %v", err)
			}

			if actualFlipCount != test.ExpectedFlipCount {
				t.Errorf("Unexpected flip count. Expected: %v, Got: %v", test.ExpectedFlipCount, actualFlipCount)
			}
		})
	}
}
