package testdata

func ConvertAToB(a *A) (b *B) {
	b = &B{}

	b.Int1 = a.Int1
	b.Int2 = a.Int2

	return
}
