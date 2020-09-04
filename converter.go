package converter

type Converter interface {
	Convert(input string) (output string)
	Build(src, dist interface{})
}

func Convert(input string) (output string) {
	// TODO: implement

	return ""
}

func Build(src, dist interface{}) {
	// TODO: implement
}
