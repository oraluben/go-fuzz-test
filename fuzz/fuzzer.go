package fuzz

import "fuzz-test/pkg/pkg1"

func Fuzz(input []byte) int {
	if len(input) > 0 {
		if pkg1.Test(int(input[0])) == nil {
			return 1
		}
	}
	return 0
}
