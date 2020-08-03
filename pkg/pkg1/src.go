package pkg1

import "crypto/sha256"

func Test(i int) error {
	if i == 0 {
		_ = sha256.New()
		panic("")
	} else {
		return nil
	}
}
