// +build converter

package testdata

type A struct {
	String1, String2 string
}

type B struct {
	String1, String2 string
}

func ConvertAToB(a *A) *B {
	converter.Build(new(*A), new(*B))

	return &B{}
}
