package k1

import (
	"fmt"
	"strings"
)

// JoinStringers joins given stringers into a single string.
func JoinStringers[T fmt.Stringer](vals []T, sep string) string {
	if len(vals) == 0 {
		return ""
	}
	parts := make([]string, len(vals))
	for i, v := range vals {
		parts[i] = v.String()
	}
	return strings.Join(parts, sep)
}
