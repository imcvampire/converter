package converter

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConverter(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		testdataDir string
	}{
		{
			name:        "Successfully gen with simple struct",
			testdataDir: "simple",
		},
	}

	for _, test := range tests {
		test := test

		input := readGolden(t, test.testdataDir, testFileTypeInput)
		expected := readGolden(t, test.testdataDir, testFileTypeExpected)

		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, expected, Convert(input))
		})
	}
}

type testFileType int

const (
	testFileTypeInput testFileType = iota
	testFileTypeExpected
)

func (t testFileType) string() string {
	switch t {
	case testFileTypeInput:
		return "input"
	case testFileTypeExpected:
		return "expected"
	default:
		panic("wrong type")
	}
}

func genTestdataFilePath(dir string, fileType testFileType) string {
	return fmt.Sprintf("./testdata/%s/%s.go", dir, fileType.string())
}

func readGolden(t *testing.T, dir string, fileType testFileType) string {
	filePath := genTestdataFilePath(dir, fileType)
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		t.Fatalf("Error loading test data: %s", err)
	}
	return string(data)
}
