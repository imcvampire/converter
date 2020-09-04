// +build converter

package testdata

type A struct {
	Ints []int
}

type B struct {
	Ints []int
}

func ConvertAToB(a *A) *B {
	converter.Build(new(*A), new(*B))

	return &B{}
}
