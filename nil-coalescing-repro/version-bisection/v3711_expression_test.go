package template

import (
	"bytes"
	"regexp"
	"testing"
)

// Drop into util/template/ of a v3.7.11 checkout and run:
//   GOFLAGS=-mod=mod go test ./util/template/ -run Test_v3711_NilCoalescing_StrictPath -v
// strictRegex matches "item" so item.* is treated as strictly-required (as withItems does).
// Present key resolves; every missing-key form is REJECTED with "item.optionalKey is missing".
func Test_v3711_NilCoalescing_StrictPath(t *testing.T) {
	strict := regexp.MustCompile("item")
	cases := []struct {
		name, expr string
		env        map[string]any
	}{
		{"present", `item.optionalKey ?? 'fallback'`, map[string]any{"item": map[string]any{"name": "a", "optionalKey": "value"}}},
		{"missing-coalesce", `item.optionalKey ?? 'fallback'`, map[string]any{"item": map[string]any{"name": "b"}}},
		{"missing-optchain", `item?.optionalKey ?? 'fallback'`, map[string]any{"item": map[string]any{"name": "b"}}},
		{"missing-bracket", `item['optionalKey'] ?? 'fallback'`, map[string]any{"item": map[string]any{"name": "b"}}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var buf bytes.Buffer
			_, err := expressionReplaceStrict(&buf, c.expr, c.env, strict)
			if err != nil {
				t.Logf("REJECTED on v3.7.11: %q -> %v", c.expr, err)
			} else {
				t.Logf("resolved on v3.7.11: %q -> %q", c.expr, buf.String())
			}
		})
	}
}
