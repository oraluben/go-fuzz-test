package fuzz

//go:noescape
func test() {}

func Fuzz(input []byte) int {
	test()
	return 0
}
