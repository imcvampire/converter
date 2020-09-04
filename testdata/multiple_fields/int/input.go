// +build converter

package testdata

type A struct {
	Int1, Int2 int
}

type B struct {
	Int1, Int2 int
}

func ConvertAToB(a *A) *B {
	converter.Build(new(*A), new(*B))

	return &B{}
}
