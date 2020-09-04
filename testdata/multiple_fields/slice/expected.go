package testdata

func ConvertAToB(a *A) (b *B) {
	b = &B{}

	b.Ints = make([]int, len(a.Ints))
	for i, v := range b.Ints {
		b.Ints[i] = v
	}

	return
}
