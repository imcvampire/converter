package testdata

func ConvertAToB(a *A) (b *B) {
	b = &B{}

	b.String = a.String

	return
}
