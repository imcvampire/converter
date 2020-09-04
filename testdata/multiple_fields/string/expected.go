package testdata

func ConvertAToB(a *A) (b *B) {
	b = &B{}

	b.String1 = a.String1
	b.String2 = a.String2

	return
}
