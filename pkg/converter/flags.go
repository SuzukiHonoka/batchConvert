package converter

import (
	"encoding/json"
	"strings"
)

type Flags map[string]string

// String converts the flags to args string
func (f Flags) String() string {
	var sb strings.Builder
	index := 1
	for k, v := range f {
		sb.WriteRune('-')
		sb.WriteString(k)
		// bypass boolean flag
		if v != "" {
			sb.WriteRune(' ')
			sb.WriteString(v)
		}
		// don't add whitespace suffix
		if index < len(f) {
			sb.WriteRune(' ')
		}
		index++
	}
	return sb.String()
}

// Json converts the flags to json string
func (f Flags) Json() string {
	b, err := json.Marshal(f)
	if err != nil {
		return ""
	}
	return string(b)
}
