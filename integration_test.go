package pikchr_test

import (
	"testing"

	pikchr "github.com/gopikchr/goldmark-pikchr"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/testutil"
)

func TestIntegration(t *testing.T) {
	t.Parallel()

	testutil.DoTestCaseFile(
		goldmark.New(goldmark.WithExtensions(&pikchr.Extender{})),
		"testdata/tests.txt",
		t,
	)
}
