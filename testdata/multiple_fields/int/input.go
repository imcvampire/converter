// +build converter

package testdata

import "github.com/imcvampire/converter"

func ConvertAToB(a *A) *B {
	converter.Build(new(*A), new(*B))

	return &B{}
}
