package rop_test

import (
	"reflect"
	"testing"

	"github.com/bradleyshawkins/rop"
)

func TestPreparePancakes(t *testing.T) {
	tests := []struct {
		Name          string
		PancakeString string
		ExpectedStack rop.Stack
	}{
		{
			Name:          "+",
			PancakeString: "+",
			ExpectedStack: rop.Stack{
				rop.Pancake(true),
			},
		},
		{
			Name:          "--",
			PancakeString: "--",
			ExpectedStack: rop.Stack{
				rop.Pancake(false),
				rop.Pancake(false),
			},
		},
		{
			Name:          "-+-",
			PancakeString: "-+-",
			ExpectedStack: rop.Stack{
				rop.Pancake(false),
				rop.Pancake(true),
				rop.Pancake(false),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			kitchen := rop.Kitchen{}

			stack, err := kitchen.PreparePancakes(test.PancakeString)
			if err != nil {
				t.Errorf("Unexpected error creating pancake stack. Error: %v", err)
			}

			if !reflect.DeepEqual(stack, test.ExpectedStack) {
				t.Errorf("Unexpected pancake stack created. Expected: %v, Got: %v", test.ExpectedStack, stack)
			}
		})
	}
}

func TestPreparePancakesError(t *testing.T) {
	tests := []struct {
		Name          string
		PancakeString string
	}{
		{Name: "Invalid character", PancakeString: "-#+"},
		{Name: "Empty string", PancakeString: ""},
		{Name: "Too Long", PancakeString: "----------++++++++++----------++++++++++----------++++++++++----------++++++++++----------++++++++++-"},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			kitchen := rop.Kitchen{}

			_, err := kitchen.PreparePancakes(test.PancakeString)
			if err == nil {
				t.Errorf("Expected an error but did not get one.")
			}
		})
	}
}
