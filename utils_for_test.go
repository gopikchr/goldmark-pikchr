package pikchr

import (
	"bytes"
	"strings"
)

// unlines returns the string formed by joining the provided strings after
// appending a newline to each.
func unlines(lines ...string) string {
	return strings.Join(lines, "\n") + "\n"
}

// trimLeadingSpace trims the leading space from each line in the given text.
func trimLeadingSpace(text string) string {
	var buf bytes.Buffer

	for i, line := range strings.Split(text, "\n") {
		if i > 0 {
			buf.WriteRune('\n')
		}

		buf.WriteString(strings.TrimLeft(line, " \t"))
	}
	return buf.String()
}
