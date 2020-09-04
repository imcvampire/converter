// +build converter

package testdata

import "github.com/imcvampire/converter"

type A struct {
	String string
}

type B struct {
	String string
}

func ConvertAToB(a *A) *B {
	converter.Build(new(*A), new(*B))

	return &B{}
}
